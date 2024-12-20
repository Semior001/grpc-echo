// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: echopb/echo.proto

package echopb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping string `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echopb_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echopb_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_echopb_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetPing() string {
	if x != nil {
		return x.Ping
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers            map[string]string      `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body               string                 `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	RemoteAddr         string                 `protobuf:"bytes,3,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	ReceivedAt         *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
	HandlerReachedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=handler_reached_at,json=handlerReachedAt,proto3" json:"handler_reached_at,omitempty"`
	HandlerRespondedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=handler_responded_at,json=handlerRespondedAt,proto3" json:"handler_responded_at,omitempty"`
	SentAt             *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echopb_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_echopb_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_echopb_echo_proto_rawDescGZIP(), []int{1}
}

func (x *EchoResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *EchoResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *EchoResponse) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *EchoResponse) GetReceivedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceivedAt
	}
	return nil
}

func (x *EchoResponse) GetHandlerReachedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.HandlerReachedAt
	}
	return nil
}

func (x *EchoResponse) GetHandlerRespondedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.HandlerRespondedAt
	}
	return nil
}

func (x *EchoResponse) GetSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

var File_echopb_echo_proto protoreflect.FileDescriptor

var file_echopb_echo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0xcc, 0x03, 0x0a, 0x0c, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65,
	0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x3b,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x41, 0x74, 0x12, 0x48, 0x0a, 0x12, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x10, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x61, 0x63,
	0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4c, 0x0a, 0x14, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x12, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0x4c, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x19, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x63,
	0x68, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x65, 0x6d, 0x69, 0x6f, 0x72, 0x30, 0x30, 0x31, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x65, 0x63, 0x68, 0x6f, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x3b, 0x65, 0x63, 0x68, 0x6f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_echopb_echo_proto_rawDescOnce sync.Once
	file_echopb_echo_proto_rawDescData = file_echopb_echo_proto_rawDesc
)

func file_echopb_echo_proto_rawDescGZIP() []byte {
	file_echopb_echo_proto_rawDescOnce.Do(func() {
		file_echopb_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echopb_echo_proto_rawDescData)
	})
	return file_echopb_echo_proto_rawDescData
}

var file_echopb_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_echopb_echo_proto_goTypes = []interface{}{
	(*EchoRequest)(nil),           // 0: grpc_echo.v1.EchoRequest
	(*EchoResponse)(nil),          // 1: grpc_echo.v1.EchoResponse
	nil,                           // 2: grpc_echo.v1.EchoResponse.HeadersEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_echopb_echo_proto_depIdxs = []int32{
	2, // 0: grpc_echo.v1.EchoResponse.headers:type_name -> grpc_echo.v1.EchoResponse.HeadersEntry
	3, // 1: grpc_echo.v1.EchoResponse.received_at:type_name -> google.protobuf.Timestamp
	3, // 2: grpc_echo.v1.EchoResponse.handler_reached_at:type_name -> google.protobuf.Timestamp
	3, // 3: grpc_echo.v1.EchoResponse.handler_responded_at:type_name -> google.protobuf.Timestamp
	3, // 4: grpc_echo.v1.EchoResponse.sent_at:type_name -> google.protobuf.Timestamp
	0, // 5: grpc_echo.v1.EchoService.Echo:input_type -> grpc_echo.v1.EchoRequest
	1, // 6: grpc_echo.v1.EchoService.Echo:output_type -> grpc_echo.v1.EchoResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_echopb_echo_proto_init() }
func file_echopb_echo_proto_init() {
	if File_echopb_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echopb_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_echopb_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_echopb_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echopb_echo_proto_goTypes,
		DependencyIndexes: file_echopb_echo_proto_depIdxs,
		MessageInfos:      file_echopb_echo_proto_msgTypes,
	}.Build()
	File_echopb_echo_proto = out.File
	file_echopb_echo_proto_rawDesc = nil
	file_echopb_echo_proto_goTypes = nil
	file_echopb_echo_proto_depIdxs = nil
}
