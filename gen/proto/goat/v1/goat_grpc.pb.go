// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package goatv1

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

// GoatServiceClient is the client API for GoatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoatServiceClient interface {
	WhatIsDoing(ctx context.Context, in *WhatIsDoingRequest, opts ...grpc.CallOption) (*WhatIsDoingResponse, error)
}

type goatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoatServiceClient(cc grpc.ClientConnInterface) GoatServiceClient {
	return &goatServiceClient{cc}
}

func (c *goatServiceClient) WhatIsDoing(ctx context.Context, in *WhatIsDoingRequest, opts ...grpc.CallOption) (*WhatIsDoingResponse, error) {
	out := new(WhatIsDoingResponse)
	err := c.cc.Invoke(ctx, "/goat.v1.GoatService/WhatIsDoing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoatServiceServer is the server API for GoatService service.
// All implementations should embed UnimplementedGoatServiceServer
// for forward compatibility
type GoatServiceServer interface {
	WhatIsDoing(context.Context, *WhatIsDoingRequest) (*WhatIsDoingResponse, error)
}

// UnimplementedGoatServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGoatServiceServer struct {
}

func (UnimplementedGoatServiceServer) WhatIsDoing(context.Context, *WhatIsDoingRequest) (*WhatIsDoingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WhatIsDoing not implemented")
}

// UnsafeGoatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoatServiceServer will
// result in compilation errors.
type UnsafeGoatServiceServer interface {
	mustEmbedUnimplementedGoatServiceServer()
}

func RegisterGoatServiceServer(s grpc.ServiceRegistrar, srv GoatServiceServer) {
	s.RegisterService(&GoatService_ServiceDesc, srv)
}

func _GoatService_WhatIsDoing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhatIsDoingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoatServiceServer).WhatIsDoing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goat.v1.GoatService/WhatIsDoing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoatServiceServer).WhatIsDoing(ctx, req.(*WhatIsDoingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoatService_ServiceDesc is the grpc.ServiceDesc for GoatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goat.v1.GoatService",
	HandlerType: (*GoatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WhatIsDoing",
			Handler:    _GoatService_WhatIsDoing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/goat/v1/goat.proto",
}
