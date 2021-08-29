// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: server_side.proto

//包名

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

// 定义发送请求信息
type ServerSideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 定义发送的参数
	// 参数类型 参数名 标识号(不可重复)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ServerSideRequest) Reset() {
	*x = ServerSideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_side_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerSideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerSideRequest) ProtoMessage() {}

func (x *ServerSideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_side_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerSideRequest.ProtoReflect.Descriptor instead.
func (*ServerSideRequest) Descriptor() ([]byte, []int) {
	return file_server_side_proto_rawDescGZIP(), []int{0}
}

func (x *ServerSideRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 定义响应信息
type ServerSideResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ServerSideResp) Reset() {
	*x = ServerSideResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_side_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerSideResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerSideResp) ProtoMessage() {}

func (x *ServerSideResp) ProtoReflect() protoreflect.Message {
	mi := &file_server_side_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerSideResp.ProtoReflect.Descriptor instead.
func (*ServerSideResp) Descriptor() ([]byte, []int) {
	return file_server_side_proto_rawDescGZIP(), []int{1}
}

func (x *ServerSideResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_server_side_proto protoreflect.FileDescriptor

var file_server_side_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x27, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x2a, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4e, 0x0a, 0x0a,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x0f, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_side_proto_rawDescOnce sync.Once
	file_server_side_proto_rawDescData = file_server_side_proto_rawDesc
)

func file_server_side_proto_rawDescGZIP() []byte {
	file_server_side_proto_rawDescOnce.Do(func() {
		file_server_side_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_side_proto_rawDescData)
	})
	return file_server_side_proto_rawDescData
}

var file_server_side_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_side_proto_goTypes = []interface{}{
	(*ServerSideRequest)(nil), // 0: pb.ServerSideRequest
	(*ServerSideResp)(nil),    // 1: pb.ServerSideResp
}
var file_server_side_proto_depIdxs = []int32{
	0, // 0: pb.ServerSide.ServerSideHello:input_type -> pb.ServerSideRequest
	1, // 1: pb.ServerSide.ServerSideHello:output_type -> pb.ServerSideResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_side_proto_init() }
func file_server_side_proto_init() {
	if File_server_side_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_side_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerSideRequest); i {
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
		file_server_side_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerSideResp); i {
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
			RawDescriptor: file_server_side_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_side_proto_goTypes,
		DependencyIndexes: file_server_side_proto_depIdxs,
		MessageInfos:      file_server_side_proto_msgTypes,
	}.Build()
	File_server_side_proto = out.File
	file_server_side_proto_rawDesc = nil
	file_server_side_proto_goTypes = nil
	file_server_side_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerSideClient is the client API for ServerSide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerSideClient interface {
	//一个ServerSideHello的方法
	ServerSideHello(ctx context.Context, in *ServerSideRequest, opts ...grpc.CallOption) (ServerSide_ServerSideHelloClient, error)
}

type serverSideClient struct {
	cc grpc.ClientConnInterface
}

func NewServerSideClient(cc grpc.ClientConnInterface) ServerSideClient {
	return &serverSideClient{cc}
}

func (c *serverSideClient) ServerSideHello(ctx context.Context, in *ServerSideRequest, opts ...grpc.CallOption) (ServerSide_ServerSideHelloClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServerSide_serviceDesc.Streams[0], "/pb.ServerSide/ServerSideHello", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverSideServerSideHelloClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerSide_ServerSideHelloClient interface {
	Recv() (*ServerSideResp, error)
	grpc.ClientStream
}

type serverSideServerSideHelloClient struct {
	grpc.ClientStream
}

func (x *serverSideServerSideHelloClient) Recv() (*ServerSideResp, error) {
	m := new(ServerSideResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerSideServer is the server API for ServerSide service.
type ServerSideServer interface {
	//一个ServerSideHello的方法
	ServerSideHello(*ServerSideRequest, ServerSide_ServerSideHelloServer) error
}

// UnimplementedServerSideServer can be embedded to have forward compatible implementations.
type UnimplementedServerSideServer struct {
}

func (*UnimplementedServerSideServer) ServerSideHello(*ServerSideRequest, ServerSide_ServerSideHelloServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerSideHello not implemented")
}

func RegisterServerSideServer(s *grpc.Server, srv ServerSideServer) {
	s.RegisterService(&_ServerSide_serviceDesc, srv)
}

func _ServerSide_ServerSideHello_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerSideRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerSideServer).ServerSideHello(m, &serverSideServerSideHelloServer{stream})
}

type ServerSide_ServerSideHelloServer interface {
	Send(*ServerSideResp) error
	grpc.ServerStream
}

type serverSideServerSideHelloServer struct {
	grpc.ServerStream
}

func (x *serverSideServerSideHelloServer) Send(m *ServerSideResp) error {
	return x.ServerStream.SendMsg(m)
}

var _ServerSide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ServerSide",
	HandlerType: (*ServerSideServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerSideHello",
			Handler:       _ServerSide_ServerSideHello_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server_side.proto",
}
