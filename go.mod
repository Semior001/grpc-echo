module github.com/Semior001/grpc-echo

go 1.23.4

replace github.com/Semior001/grpc-echo/echopb => ./echopb

require (
	github.com/Semior001/grpc-echo/echopb v0.0.0-00010101000000-000000000000
	github.com/jessevdk/go-flags v1.6.1
	golang.org/x/sync v0.8.0
	google.golang.org/grpc v1.68.1
	google.golang.org/protobuf v1.35.2
)

require (
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
)
