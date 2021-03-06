// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: ping.svc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"` // type of request
	Rid  string `protobuf:"bytes,2,opt,name=rid,proto3" json:"rid,omitempty"`   // request id = uuid
	Tid  string `protobuf:"bytes,3,opt,name=tid,proto3" json:"tid,omitempty"`   // transaction id = uuid
	From string `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"` // requestor service name
	To   string `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`     // receiver service name
	Msg  string `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg,omitempty"`   // msg = "PING"
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ping_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_ping_svc_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PingRequest) GetRid() string {
	if x != nil {
		return x.Rid
	}
	return ""
}

func (x *PingRequest) GetTid() string {
	if x != nil {
		return x.Tid
	}
	return ""
}

func (x *PingRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PingRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *PingRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type PongReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"` // type of request
	Rid  string `protobuf:"bytes,2,opt,name=rid,proto3" json:"rid,omitempty"`   // response id = uuid
	Tid  string `protobuf:"bytes,3,opt,name=tid,proto3" json:"tid,omitempty"`   // transaction id = uuid
	From string `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"` // requestor service name
	To   string `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`     // receiver service name
	Msg  string `protobuf:"bytes,6,opt,name=msg,proto3" json:"msg,omitempty"`   // msg = "PONG"
}

func (x *PongReply) Reset() {
	*x = PongReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ping_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongReply) ProtoMessage() {}

func (x *PongReply) ProtoReflect() protoreflect.Message {
	mi := &file_ping_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongReply.ProtoReflect.Descriptor instead.
func (*PongReply) Descriptor() ([]byte, []int) {
	return file_ping_svc_proto_rawDescGZIP(), []int{1}
}

func (x *PongReply) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *PongReply) GetRid() string {
	if x != nil {
		return x.Rid
	}
	return ""
}

func (x *PongReply) GetTid() string {
	if x != nil {
		return x.Tid
	}
	return ""
}

func (x *PongReply) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PongReply) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *PongReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_ping_svc_proto protoreflect.FileDescriptor

var file_ping_svc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x7b, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x79, 0x0a, 0x09, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x72, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x30, 0x0a, 0x04,
	0x50, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ping_svc_proto_rawDescOnce sync.Once
	file_ping_svc_proto_rawDescData = file_ping_svc_proto_rawDesc
)

func file_ping_svc_proto_rawDescGZIP() []byte {
	file_ping_svc_proto_rawDescOnce.Do(func() {
		file_ping_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_ping_svc_proto_rawDescData)
	})
	return file_ping_svc_proto_rawDescData
}

var file_ping_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ping_svc_proto_goTypes = []interface{}{
	(*PingRequest)(nil), // 0: pb.PingRequest
	(*PongReply)(nil),   // 1: pb.PongReply
}
var file_ping_svc_proto_depIdxs = []int32{
	0, // 0: pb.Ping.ping:input_type -> pb.PingRequest
	1, // 1: pb.Ping.ping:output_type -> pb.PongReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ping_svc_proto_init() }
func file_ping_svc_proto_init() {
	if File_ping_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ping_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_ping_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongReply); i {
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
			RawDescriptor: file_ping_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ping_svc_proto_goTypes,
		DependencyIndexes: file_ping_svc_proto_depIdxs,
		MessageInfos:      file_ping_svc_proto_msgTypes,
	}.Build()
	File_ping_svc_proto = out.File
	file_ping_svc_proto_rawDesc = nil
	file_ping_svc_proto_goTypes = nil
	file_ping_svc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PingClient is the client API for Ping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingClient interface {
	// Ping command
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error)
}

type pingClient struct {
	cc grpc.ClientConnInterface
}

func NewPingClient(cc grpc.ClientConnInterface) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PongReply, error) {
	out := new(PongReply)
	err := c.cc.Invoke(ctx, "/pb.Ping/ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServer is the server API for Ping service.
type PingServer interface {
	// Ping command
	Ping(context.Context, *PingRequest) (*PongReply, error)
}

// UnimplementedPingServer can be embedded to have forward compatible implementations.
type UnimplementedPingServer struct {
}

func (*UnimplementedPingServer) Ping(context.Context, *PingRequest) (*PongReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterPingServer(s *grpc.Server, srv PingServer) {
	s.RegisterService(&_Ping_serviceDesc, srv)
}

func _Ping_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Ping/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ping_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Ping",
	HandlerType: (*PingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ping",
			Handler:    _Ping_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ping.svc.proto",
}
