// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("peer.proto", fileDescriptor_055ae5a865fc1c9e)
}

var fileDescriptor_055ae5a865fc1c9e = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0x3d, 0x4b, 0x03, 0x41,
	0x10, 0x40, 0x2f, 0xf8, 0x41, 0x32, 0x44, 0x85, 0xc1, 0x6a, 0xcb, 0xc3, 0xc2, 0x2a, 0xc4, 0xb5,
	0xb1, 0x49, 0x13, 0x49, 0x29, 0x2c, 0xe7, 0x2f, 0xd8, 0xdb, 0x1b, 0xce, 0x03, 0x77, 0x67, 0x99,
	0x59, 0x0b, 0xff, 0xbd, 0x70, 0x7b, 0x85, 0x5a, 0xa5, 0x1a, 0xde, 0x9b, 0x37, 0x30, 0x00, 0x99,
	0x48, 0x76, 0x59, 0xb8, 0x30, 0x5e, 0xcd, 0xc3, 0x6c, 0x03, 0xc7, 0xc8, 0xa9, 0x4a, 0x73, 0x13,
	0x49, 0xd5, 0x8f, 0x54, 0xd1, 0x32, 0xac, 0x4f, 0x69, 0x60, 0x51, 0x12, 0x3c, 0xc0, 0x9d, 0x13,
	0x0e, 0xa4, 0xea, 0x84, 0x33, 0xab, 0xff, 0x44, 0xac, 0xd9, 0xee, 0xe8, 0x4b, 0xf8, 0x78, 0xff,
	0xea, 0xe3, 0x54, 0xcc, 0xfd, 0x6f, 0xd7, 0x91, 0x66, 0x4e, 0x4a, 0x6d, 0x83, 0x0f, 0x70, 0xe9,
	0xa6, 0x34, 0xe2, 0x76, 0xd9, 0x9f, 0x62, 0x2e, 0xdf, 0xe6, 0x0f, 0xb5, 0x8d, 0x3d, 0xc0, 0xe6,
	0x95, 0x63, 0xe6, 0x44, 0xa9, 0xe0, 0x1e, 0xd6, 0x1d, 0x8d, 0x93, 0x16, 0x12, 0xbc, 0x5d, 0xc2,
	0xb7, 0xfa, 0x9f, 0xf9, 0xc7, 0x6d, 0xf3, 0xb8, 0xda, 0xaf, 0xec, 0x0b, 0x5c, 0x38, 0xeb, 0xf0,
	0x09, 0x36, 0x47, 0x61, 0x3f, 0x04, 0xaf, 0xe5, 0xbc, 0xcb, 0xfe, 0x7a, 0x96, 0xcf, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x34, 0x3d, 0xa0, 0x84, 0x22, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EndorserClient is the client API for Endorser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndorserClient interface {
	ProcessProposal(ctx context.Context, in *BatchSubmit, opts ...grpc.CallOption) (*BatchResponse, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type endorserClient struct {
	cc grpc.ClientConnInterface
}

func NewEndorserClient(cc grpc.ClientConnInterface) EndorserClient {
	return &endorserClient{cc}
}

func (c *endorserClient) ProcessProposal(ctx context.Context, in *BatchSubmit, opts ...grpc.CallOption) (*BatchResponse, error) {
	out := new(BatchResponse)
	err := c.cc.Invoke(ctx, "/proto.Endorser/ProcessProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *endorserClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.Endorser/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndorserServer is the server API for Endorser service.
type EndorserServer interface {
	ProcessProposal(context.Context, *BatchSubmit) (*BatchResponse, error)
	Ping(context.Context, *Empty) (*Empty, error)
}

// UnimplementedEndorserServer can be embedded to have forward compatible implementations.
type UnimplementedEndorserServer struct {
}

func (*UnimplementedEndorserServer) ProcessProposal(ctx context.Context, req *BatchSubmit) (*BatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessProposal not implemented")
}
func (*UnimplementedEndorserServer) Ping(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}

func RegisterEndorserServer(s *grpc.Server, srv EndorserServer) {
	s.RegisterService(&_Endorser_serviceDesc, srv)
}

func _Endorser_ProcessProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchSubmit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndorserServer).ProcessProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Endorser/ProcessProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndorserServer).ProcessProposal(ctx, req.(*BatchSubmit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Endorser_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndorserServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Endorser/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndorserServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Endorser_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Endorser",
	HandlerType: (*EndorserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessProposal",
			Handler:    _Endorser_ProcessProposal_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Endorser_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "peer.proto",
}

// ComponentClient is the client API for Component service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ComponentClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (Component_RegisterClient, error)
}

type componentClient struct {
	cc grpc.ClientConnInterface
}

func NewComponentClient(cc grpc.ClientConnInterface) ComponentClient {
	return &componentClient{cc}
}

func (c *componentClient) Register(ctx context.Context, opts ...grpc.CallOption) (Component_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Component_serviceDesc.Streams[0], "/proto.Component/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &componentRegisterClient{stream}
	return x, nil
}

type Component_RegisterClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type componentRegisterClient struct {
	grpc.ClientStream
}

func (x *componentRegisterClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *componentRegisterClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ComponentServer is the server API for Component service.
type ComponentServer interface {
	Register(Component_RegisterServer) error
}

// UnimplementedComponentServer can be embedded to have forward compatible implementations.
type UnimplementedComponentServer struct {
}

func (*UnimplementedComponentServer) Register(srv Component_RegisterServer) error {
	return status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterComponentServer(s *grpc.Server, srv ComponentServer) {
	s.RegisterService(&_Component_serviceDesc, srv)
}

func _Component_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ComponentServer).Register(&componentRegisterServer{stream})
}

type Component_RegisterServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type componentRegisterServer struct {
	grpc.ServerStream
}

func (x *componentRegisterServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *componentRegisterServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Component_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Component",
	HandlerType: (*ComponentServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _Component_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "peer.proto",
}

// P2PClient is the client API for P2P service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type P2PClient interface {
	Broadcast(ctx context.Context, opts ...grpc.CallOption) (P2P_BroadcastClient, error)
}

type p2PClient struct {
	cc grpc.ClientConnInterface
}

func NewP2PClient(cc grpc.ClientConnInterface) P2PClient {
	return &p2PClient{cc}
}

func (c *p2PClient) Broadcast(ctx context.Context, opts ...grpc.CallOption) (P2P_BroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &_P2P_serviceDesc.Streams[0], "/proto.P2P/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &p2PBroadcastClient{stream}
	return x, nil
}

type P2P_BroadcastClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type p2PBroadcastClient struct {
	grpc.ClientStream
}

func (x *p2PBroadcastClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *p2PBroadcastClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// P2PServer is the server API for P2P service.
type P2PServer interface {
	Broadcast(P2P_BroadcastServer) error
}

// UnimplementedP2PServer can be embedded to have forward compatible implementations.
type UnimplementedP2PServer struct {
}

func (*UnimplementedP2PServer) Broadcast(srv P2P_BroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}

func RegisterP2PServer(s *grpc.Server, srv P2PServer) {
	s.RegisterService(&_P2P_serviceDesc, srv)
}

func _P2P_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(P2PServer).Broadcast(&p2PBroadcastServer{stream})
}

type P2P_BroadcastServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type p2PBroadcastServer struct {
	grpc.ServerStream
}

func (x *p2PBroadcastServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *p2PBroadcastServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _P2P_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.P2P",
	HandlerType: (*P2PServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Broadcast",
			Handler:       _P2P_Broadcast_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "peer.proto",
}
