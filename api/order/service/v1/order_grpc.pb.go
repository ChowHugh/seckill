// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderReply, error)
	CreateSeckillOrder(ctx context.Context, in *CreateSeckillOrderRequest, opts ...grpc.CallOption) (*CreateSeckillOrderReply, error)
	CreateSeckillOrderTry(ctx context.Context, in *CreateSeckillOrderRequest, opts ...grpc.CallOption) (*CreateSeckillOrderReply, error)
	CreateSeckillOrderConfirm(ctx context.Context, in *CreateSeckillOrderRequest, opts ...grpc.CallOption) (*CreateSeckillOrderReply, error)
	CreateSeckillOrderCancel(ctx context.Context, in *CreateSeckillOrderRequest, opts ...grpc.CallOption) (*CreateSeckillOrderReply, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderReply, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderReply, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderReply, error)
	ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderReply, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderReply, error) {
	out := new(CreateOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateSeckillOrder(ctx context.Context, in *CreateSeckillOrderRequest, opts ...grpc.CallOption) (*CreateSeckillOrderReply, error) {
	out := new(CreateSeckillOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/CreateSeckillOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateSeckillOrderTry(ctx context.Context, in *CreateSeckillOrderRequest, opts ...grpc.CallOption) (*CreateSeckillOrderReply, error) {
	out := new(CreateSeckillOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/CreateSeckillOrderTry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateSeckillOrderConfirm(ctx context.Context, in *CreateSeckillOrderRequest, opts ...grpc.CallOption) (*CreateSeckillOrderReply, error) {
	out := new(CreateSeckillOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/CreateSeckillOrderConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateSeckillOrderCancel(ctx context.Context, in *CreateSeckillOrderRequest, opts ...grpc.CallOption) (*CreateSeckillOrderReply, error) {
	out := new(CreateSeckillOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/CreateSeckillOrderCancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderReply, error) {
	out := new(UpdateOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderReply, error) {
	out := new(DeleteOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/DeleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*GetOrderReply, error) {
	out := new(GetOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) ListOrder(ctx context.Context, in *ListOrderRequest, opts ...grpc.CallOption) (*ListOrderReply, error) {
	out := new(ListOrderReply)
	err := c.cc.Invoke(ctx, "/api.order.service.v1.Order/ListOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderReply, error)
	CreateSeckillOrder(context.Context, *CreateSeckillOrderRequest) (*CreateSeckillOrderReply, error)
	CreateSeckillOrderTry(context.Context, *CreateSeckillOrderRequest) (*CreateSeckillOrderReply, error)
	CreateSeckillOrderConfirm(context.Context, *CreateSeckillOrderRequest) (*CreateSeckillOrderReply, error)
	CreateSeckillOrderCancel(context.Context, *CreateSeckillOrderRequest) (*CreateSeckillOrderReply, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderReply, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderReply, error)
	GetOrder(context.Context, *GetOrderRequest) (*GetOrderReply, error)
	ListOrder(context.Context, *ListOrderRequest) (*ListOrderReply, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServer) CreateSeckillOrder(context.Context, *CreateSeckillOrderRequest) (*CreateSeckillOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeckillOrder not implemented")
}
func (UnimplementedOrderServer) CreateSeckillOrderTry(context.Context, *CreateSeckillOrderRequest) (*CreateSeckillOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeckillOrderTry not implemented")
}
func (UnimplementedOrderServer) CreateSeckillOrderConfirm(context.Context, *CreateSeckillOrderRequest) (*CreateSeckillOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeckillOrderConfirm not implemented")
}
func (UnimplementedOrderServer) CreateSeckillOrderCancel(context.Context, *CreateSeckillOrderRequest) (*CreateSeckillOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeckillOrderCancel not implemented")
}
func (UnimplementedOrderServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderServer) DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderServer) GetOrder(context.Context, *GetOrderRequest) (*GetOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOrderServer) ListOrder(context.Context, *ListOrderRequest) (*ListOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOrder not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateSeckillOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSeckillOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateSeckillOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/CreateSeckillOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateSeckillOrder(ctx, req.(*CreateSeckillOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateSeckillOrderTry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSeckillOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateSeckillOrderTry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/CreateSeckillOrderTry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateSeckillOrderTry(ctx, req.(*CreateSeckillOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateSeckillOrderConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSeckillOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateSeckillOrderConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/CreateSeckillOrderConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateSeckillOrderConfirm(ctx, req.(*CreateSeckillOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateSeckillOrderCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSeckillOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateSeckillOrderCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/CreateSeckillOrderCancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateSeckillOrderCancel(ctx, req.(*CreateSeckillOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/DeleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).DeleteOrder(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_ListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).ListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order.service.v1.Order/ListOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).ListOrder(ctx, req.(*ListOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.order.service.v1.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _Order_CreateOrder_Handler,
		},
		{
			MethodName: "CreateSeckillOrder",
			Handler:    _Order_CreateSeckillOrder_Handler,
		},
		{
			MethodName: "CreateSeckillOrderTry",
			Handler:    _Order_CreateSeckillOrderTry_Handler,
		},
		{
			MethodName: "CreateSeckillOrderConfirm",
			Handler:    _Order_CreateSeckillOrderConfirm_Handler,
		},
		{
			MethodName: "CreateSeckillOrderCancel",
			Handler:    _Order_CreateSeckillOrderCancel_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _Order_UpdateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _Order_DeleteOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Order_GetOrder_Handler,
		},
		{
			MethodName: "ListOrder",
			Handler:    _Order_ListOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/order.proto",
}
