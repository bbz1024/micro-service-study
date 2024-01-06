// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pay

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

// PayServiceClient is the client API for PayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayServiceClient interface {
	Pay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error)
	PayRevert(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error)
}

type payServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayServiceClient(cc grpc.ClientConnInterface) PayServiceClient {
	return &payServiceClient{cc}
}

func (c *payServiceClient) Pay(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error) {
	out := new(PayResponse)
	err := c.cc.Invoke(ctx, "/pay.PayService/Pay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServiceClient) PayRevert(ctx context.Context, in *PayRequest, opts ...grpc.CallOption) (*PayResponse, error) {
	out := new(PayResponse)
	err := c.cc.Invoke(ctx, "/pay.PayService/PayRevert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayServiceServer is the server API for PayService service.
// All implementations must embed UnimplementedPayServiceServer
// for forward compatibility
type PayServiceServer interface {
	Pay(context.Context, *PayRequest) (*PayResponse, error)
	PayRevert(context.Context, *PayRequest) (*PayResponse, error)
	mustEmbedUnimplementedPayServiceServer()
}

// UnimplementedPayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPayServiceServer struct {
}

func (UnimplementedPayServiceServer) Pay(context.Context, *PayRequest) (*PayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}
func (UnimplementedPayServiceServer) PayRevert(context.Context, *PayRequest) (*PayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayRevert not implemented")
}
func (UnimplementedPayServiceServer) mustEmbedUnimplementedPayServiceServer() {}

// UnsafePayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayServiceServer will
// result in compilation errors.
type UnsafePayServiceServer interface {
	mustEmbedUnimplementedPayServiceServer()
}

func RegisterPayServiceServer(s grpc.ServiceRegistrar, srv PayServiceServer) {
	s.RegisterService(&PayService_ServiceDesc, srv)
}

func _PayService_Pay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServiceServer).Pay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay.PayService/Pay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServiceServer).Pay(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayService_PayRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServiceServer).PayRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pay.PayService/PayRevert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServiceServer).PayRevert(ctx, req.(*PayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayService_ServiceDesc is the grpc.ServiceDesc for PayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pay.PayService",
	HandlerType: (*PayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pay",
			Handler:    _PayService_Pay_Handler,
		},
		{
			MethodName: "PayRevert",
			Handler:    _PayService_PayRevert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pay.proto",
}
