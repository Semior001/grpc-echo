// Package main is an application entrypoint.
package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/jessevdk/go-flags"
	"github.com/Semior001/grpc-echo/echopb"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc"
	"net"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/health"
	"github.com/Semior001/grpc-echo/pkg/grpcx"
	"github.com/Semior001/grpc-echo/pkg/service"
	"google.golang.org/grpc/credentials"
	"sync"
	"google.golang.org/grpc/grpclog"
	"time"
	"google.golang.org/grpc/keepalive"
)

var opts struct {
	SSL struct {
		Enable bool   `long:"enable"        env:"ENABLE" description:"Enable SSL"`
		Cert   string `long:"cert"          env:"CERT"   description:"path to cert.pem file"`
		Key    string `long:"key"           env:"KEY"    description:"path to key.pem file"`
	} `group:"ssl" namespace:"ssl" env-namespace:"SSL" description:"ssl settings"`

	Keepalive struct {
		MaxConnIdle time.Duration `long:"max-conn-idle" env:"MAX_CONN_IDLE"      default:"3s"   description:"max time a connection can be idle"`
		MaxConnAge  time.Duration `long:"max-conn-age"  env:"MAX_CONN_AGE_GRACE" default:"5s"   description:"max time a connection can exist (jitter +/-10%)"`
		Time        time.Duration `long:"time"          env:"TIME"               default:"1s"   description:"interval between server pings"`
	} `group:"keepalive" namespace:"keepalive" env-namespace:"KEEPALIVE" description:"keepalive settings"`

	StreamTimeout time.Duration `long:"stream-timeout" env:"STREAM_TIMEOUT" default:"5s" description:"stream timeout, 0 means no timeout"`

	Addr  string `short:"a" long:"addr" env:"ADDR" default:":8080" description:"Address to listen on"`
	JSON  bool   `long:"json"           env:"JSON"                 description:"Enable JSON logging"`
	Debug bool   `long:"debug"          env:"DEBUG"                description:"Enable debug mode"`
}

var version = "unknown"

func getVersion() string {
	if bi, ok := debug.ReadBuildInfo(); ok && version == "unknown" {
		return bi.Main.Version
	}
	return version
}

func main() {
	_, _ = fmt.Fprintf(os.Stderr, "grpc-echo %s\n", getVersion())

	if _, err := flags.Parse(&opts); err != nil {
		os.Exit(1)
	}

	setupLog(opts.Debug, opts.JSON)

	ctx, cancel := context.WithCancel(context.Background())
	go func() { // catch signal and invoke graceful termination
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		sig := <-stop
		slog.Warn("caught signal", slog.Any("signal", sig))
		cancel()
	}()

	if err := run(ctx); err != nil {
		slog.Error("failed to start service", slog.Any("error", err))
	}
}

func run(ctx context.Context) error {
	svc := &service.EchoService{}
	healthHandler := health.NewServer()

	var cred credentials.TransportCredentials

	lis, err := net.Listen("tcp", opts.Addr)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", opts.Addr, err)
	}

	if opts.SSL.Enable {
		if opts.SSL.Cert == "" || opts.SSL.Key == "" {
			return fmt.Errorf("cert and key must be provided for static ssl")
		}

		slog.Info("using static ssl",
			slog.String("cert", opts.SSL.Cert),
			slog.String("key", opts.SSL.Key))

		if cred, err = credentials.NewServerTLSFromFile(opts.SSL.Cert, opts.SSL.Key); err != nil {
			return fmt.Errorf("load cert and key: %w", err)
		}
	}

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			svc.AppendTimestampInterceptor,
			grpcx.LogUnaryInterceptor,
		),
		grpc.ChainStreamInterceptor(
			grpcx.LogStreamInterceptor,
			grpcx.TimeoutStreamInterceptor(opts.StreamTimeout),
		),
		grpc.Creds(cred),
		grpc.ConnectionTimeout(5*time.Second),
		grpc.MaxConcurrentStreams(1000),
		grpc.MaxHeaderListSize(1024*4),     // 4KB
		grpc.MaxRecvMsgSize(1024*4),        // 4KB
		grpc.MaxSendMsgSize(1024*4),        // 4KB
		grpc.HeaderTableSize(1024*4),       // 4KB
		grpc.InitialWindowSize(1024*4),     // 4KB
		grpc.InitialConnWindowSize(1024*4), // 4KB
		grpc.WriteBufferSize(1024*4),       // 4KB
		grpc.ReadBufferSize(1024*4),        // 4KB
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: opts.Keepalive.MaxConnIdle,
			MaxConnectionAge:  opts.Keepalive.MaxConnAge,
			Time:              opts.Keepalive.Time,
		}),
	)
	healthpb.RegisterHealthServer(srv, healthHandler)
	echopb.RegisterEchoServiceServer(srv, svc)
	reflection.Register(srv)

	ewg, ctx := errgroup.WithContext(ctx)
	ewg.Go(func() error {
		slog.Info("listening gRPC", slog.String("addr", lis.Addr().String()))
		healthHandler.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
		if err := srv.Serve(lis); err != nil {
			return fmt.Errorf("proxy server: %w", err)
		}
		return nil
	})
	ewg.Go(func() error {
		<-ctx.Done()
		slog.Info("shutting down gRPC")
		healthHandler.SetServingStatus("", healthpb.HealthCheckResponse_NOT_SERVING)
		srv.GracefulStop()
		return nil
	})

	if err := ewg.Wait(); err != nil {
		return err
	}

	return nil
}

var setupLoggerOnce sync.Once

func setupLog(dbg, json bool) {
	defer slog.Info("prepared logger", slog.Bool("debug", dbg), slog.Bool("json", json))

	handlerOpts := &slog.HandlerOptions{Level: slog.LevelInfo}

	if dbg {
		handlerOpts.Level = slog.LevelDebug
	}

	var handler slog.Handler
	if json {
		handler = slog.NewJSONHandler(os.Stderr, handlerOpts)
	} else {
		handler = slog.NewTextHandler(os.Stderr, handlerOpts)
	}

	slog.SetDefault(slog.New(handler))
	setupLoggerOnce.Do(func() {
		grpclog.SetLoggerV2(grpcx.Logger{
			Logger: slog.NewLogLogger(handler, slog.LevelDebug),
		})
	})
}
