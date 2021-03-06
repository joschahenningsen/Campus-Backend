// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// CampusClient is the client API for Campus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CampusClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type campusClient struct {
	cc grpc.ClientConnInterface
}

func NewCampusClient(cc grpc.ClientConnInterface) CampusClient {
	return &campusClient{cc}
}

func (c *campusClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/api.Campus/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampusServer is the server API for Campus service.
// All implementations must embed UnimplementedCampusServer
// for forward compatibility
type CampusServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedCampusServer()
}

// UnimplementedCampusServer must be embedded to have forward compatible implementations.
type UnimplementedCampusServer struct {
}

func (UnimplementedCampusServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedCampusServer) mustEmbedUnimplementedCampusServer() {}

// UnsafeCampusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CampusServer will
// result in compilation errors.
type UnsafeCampusServer interface {
	mustEmbedUnimplementedCampusServer()
}

func RegisterCampusServer(s grpc.ServiceRegistrar, srv CampusServer) {
	s.RegisterService(&Campus_ServiceDesc, srv)
}

func _Campus_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampusServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Campus/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampusServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Campus_ServiceDesc is the grpc.ServiceDesc for Campus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Campus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Campus",
	HandlerType: (*CampusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Campus_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "CampusService.proto",
}
