// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: rocket/v1/rocket.proto

package rocket

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

// RocketServiceClient is the client API for RocketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RocketServiceClient interface {
	GetRocket(ctx context.Context, in *GetRocketRequest, opts ...grpc.CallOption) (*GetRocketResponse, error)
	AddRocket(ctx context.Context, in *AddRocketRequest, opts ...grpc.CallOption) (*AddRocketResponse, error)
	DeleteRocket(ctx context.Context, in *DeleteRocketRequest, opts ...grpc.CallOption) (*DeleteRocketResponse, error)
}

type rocketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRocketServiceClient(cc grpc.ClientConnInterface) RocketServiceClient {
	return &rocketServiceClient{cc}
}

func (c *rocketServiceClient) GetRocket(ctx context.Context, in *GetRocketRequest, opts ...grpc.CallOption) (*GetRocketResponse, error) {
	out := new(GetRocketResponse)
	err := c.cc.Invoke(ctx, "/rocket.RocketService/GetRocket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rocketServiceClient) AddRocket(ctx context.Context, in *AddRocketRequest, opts ...grpc.CallOption) (*AddRocketResponse, error) {
	out := new(AddRocketResponse)
	err := c.cc.Invoke(ctx, "/rocket.RocketService/AddRocket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rocketServiceClient) DeleteRocket(ctx context.Context, in *DeleteRocketRequest, opts ...grpc.CallOption) (*DeleteRocketResponse, error) {
	out := new(DeleteRocketResponse)
	err := c.cc.Invoke(ctx, "/rocket.RocketService/DeleteRocket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RocketServiceServer is the server API for RocketService service.
// All implementations must embed UnimplementedRocketServiceServer
// for forward compatibility
type RocketServiceServer interface {
	GetRocket(context.Context, *GetRocketRequest) (*GetRocketResponse, error)
	AddRocket(context.Context, *AddRocketRequest) (*AddRocketResponse, error)
	DeleteRocket(context.Context, *DeleteRocketRequest) (*DeleteRocketResponse, error)
}

// UnimplementedRocketServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRocketServiceServer struct {
}

func (UnimplementedRocketServiceServer) GetRocket(context.Context, *GetRocketRequest) (*GetRocketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRocket not implemented")
}
func (UnimplementedRocketServiceServer) AddRocket(context.Context, *AddRocketRequest) (*AddRocketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRocket not implemented")
}
func (UnimplementedRocketServiceServer) DeleteRocket(context.Context, *DeleteRocketRequest) (*DeleteRocketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRocket not implemented")
}
func (UnimplementedRocketServiceServer) mustEmbedUnimplementedRocketServiceServer() {}

// UnsafeRocketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RocketServiceServer will
// result in compilation errors.
type UnsafeRocketServiceServer interface {
	mustEmbedUnimplementedRocketServiceServer()
}

func RegisterRocketServiceServer(s grpc.ServiceRegistrar, srv RocketServiceServer) {
	s.RegisterService(&RocketService_ServiceDesc, srv)
}

func _RocketService_GetRocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRocketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RocketServiceServer).GetRocket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocket.RocketService/GetRocket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RocketServiceServer).GetRocket(ctx, req.(*GetRocketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RocketService_AddRocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRocketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RocketServiceServer).AddRocket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocket.RocketService/AddRocket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RocketServiceServer).AddRocket(ctx, req.(*AddRocketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RocketService_DeleteRocket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRocketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RocketServiceServer).DeleteRocket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rocket.RocketService/DeleteRocket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RocketServiceServer).DeleteRocket(ctx, req.(*DeleteRocketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RocketService_ServiceDesc is the grpc.ServiceDesc for RocketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RocketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rocket.RocketService",
	HandlerType: (*RocketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRocket",
			Handler:    _RocketService_GetRocket_Handler,
		},
		{
			MethodName: "AddRocket",
			Handler:    _RocketService_AddRocket_Handler,
		},
		{
			MethodName: "DeleteRocket",
			Handler:    _RocketService_DeleteRocket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rocket/v1/rocket.proto",
}
