// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestClient interface {
	UnaryCall(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*TestMessage, error)
	ServerStream(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (Test_ServerStreamClient, error)
	BidiStream(ctx context.Context, opts ...grpc.CallOption) (Test_BidiStreamClient, error)
}

type testClient struct {
	cc grpc.ClientConnInterface
}

func NewTestClient(cc grpc.ClientConnInterface) TestClient {
	return &testClient{cc}
}

func (c *testClient) UnaryCall(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*TestMessage, error) {
	out := new(TestMessage)
	err := c.cc.Invoke(ctx, "/rektorphi.arpcnet.test.Test/UnaryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) ServerStream(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (Test_ServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Test_ServiceDesc.Streams[0], "/rektorphi.arpcnet.test.Test/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Test_ServerStreamClient interface {
	Recv() (*TestMessage, error)
	grpc.ClientStream
}

type testServerStreamClient struct {
	grpc.ClientStream
}

func (x *testServerStreamClient) Recv() (*TestMessage, error) {
	m := new(TestMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testClient) BidiStream(ctx context.Context, opts ...grpc.CallOption) (Test_BidiStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Test_ServiceDesc.Streams[1], "/rektorphi.arpcnet.test.Test/BidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testBidiStreamClient{stream}
	return x, nil
}

type Test_BidiStreamClient interface {
	Send(*TestMessage) error
	Recv() (*TestMessage, error)
	grpc.ClientStream
}

type testBidiStreamClient struct {
	grpc.ClientStream
}

func (x *testBidiStreamClient) Send(m *TestMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testBidiStreamClient) Recv() (*TestMessage, error) {
	m := new(TestMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServer is the server API for Test service.
// All implementations must embed UnimplementedTestServer
// for forward compatibility
type TestServer interface {
	UnaryCall(context.Context, *TestMessage) (*TestMessage, error)
	ServerStream(*TestMessage, Test_ServerStreamServer) error
	BidiStream(Test_BidiStreamServer) error
	mustEmbedUnimplementedTestServer()
}

// UnimplementedTestServer must be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (UnimplementedTestServer) UnaryCall(context.Context, *TestMessage) (*TestMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryCall not implemented")
}
func (UnimplementedTestServer) ServerStream(*TestMessage, Test_ServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStream not implemented")
}
func (UnimplementedTestServer) BidiStream(Test_BidiStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BidiStream not implemented")
}
func (UnimplementedTestServer) mustEmbedUnimplementedTestServer() {}

// UnsafeTestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServer will
// result in compilation errors.
type UnsafeTestServer interface {
	mustEmbedUnimplementedTestServer()
}

func RegisterTestServer(s grpc.ServiceRegistrar, srv TestServer) {
	s.RegisterService(&Test_ServiceDesc, srv)
}

func _Test_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rektorphi.arpcnet.test.Test/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).UnaryCall(ctx, req.(*TestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TestMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServer).ServerStream(m, &testServerStreamServer{stream})
}

type Test_ServerStreamServer interface {
	Send(*TestMessage) error
	grpc.ServerStream
}

type testServerStreamServer struct {
	grpc.ServerStream
}

func (x *testServerStreamServer) Send(m *TestMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Test_BidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).BidiStream(&testBidiStreamServer{stream})
}

type Test_BidiStreamServer interface {
	Send(*TestMessage) error
	Recv() (*TestMessage, error)
	grpc.ServerStream
}

type testBidiStreamServer struct {
	grpc.ServerStream
}

func (x *testBidiStreamServer) Send(m *TestMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testBidiStreamServer) Recv() (*TestMessage, error) {
	m := new(TestMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Test_ServiceDesc is the grpc.ServiceDesc for Test service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Test_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rektorphi.arpcnet.test.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _Test_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStream",
			Handler:       _Test_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidiStream",
			Handler:       _Test_BidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rektorphi/arpcnet/test/test_service.proto",
}
