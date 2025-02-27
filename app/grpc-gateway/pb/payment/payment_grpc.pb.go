// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: payment.proto

package payment

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PaymentService_Charge_FullMethodName                = "/payment.PaymentService/Charge"
	PaymentService_StartPayment_FullMethodName          = "/payment.PaymentService/StartPayment"
	PaymentService_CallBack_FullMethodName              = "/payment.PaymentService/CallBack"
	PaymentService_GetOrderPayemntStatus_FullMethodName = "/payment.PaymentService/GetOrderPayemntStatus"
)

// PaymentServiceClient is the client API for PaymentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentServiceClient interface {
	Charge(ctx context.Context, in *ChargeReq, opts ...grpc.CallOption) (*ChargeResp, error)
	StartPayment(ctx context.Context, in *StartPaymentReq, opts ...grpc.CallOption) (*StartPaymentResp, error)
	CallBack(ctx context.Context, in *CallBackReq, opts ...grpc.CallOption) (*AlipayCallbackNotificationResp, error)
	GetOrderPayemntStatus(ctx context.Context, in *GetOrderPayemntStatusReq, opts ...grpc.CallOption) (*GetOrderPayemntStatusResp, error)
}

type paymentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentServiceClient(cc grpc.ClientConnInterface) PaymentServiceClient {
	return &paymentServiceClient{cc}
}

func (c *paymentServiceClient) Charge(ctx context.Context, in *ChargeReq, opts ...grpc.CallOption) (*ChargeResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChargeResp)
	err := c.cc.Invoke(ctx, PaymentService_Charge_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) StartPayment(ctx context.Context, in *StartPaymentReq, opts ...grpc.CallOption) (*StartPaymentResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartPaymentResp)
	err := c.cc.Invoke(ctx, PaymentService_StartPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) CallBack(ctx context.Context, in *CallBackReq, opts ...grpc.CallOption) (*AlipayCallbackNotificationResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlipayCallbackNotificationResp)
	err := c.cc.Invoke(ctx, PaymentService_CallBack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentServiceClient) GetOrderPayemntStatus(ctx context.Context, in *GetOrderPayemntStatusReq, opts ...grpc.CallOption) (*GetOrderPayemntStatusResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderPayemntStatusResp)
	err := c.cc.Invoke(ctx, PaymentService_GetOrderPayemntStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentServiceServer is the server API for PaymentService service.
// All implementations must embed UnimplementedPaymentServiceServer
// for forward compatibility.
type PaymentServiceServer interface {
	Charge(context.Context, *ChargeReq) (*ChargeResp, error)
	StartPayment(context.Context, *StartPaymentReq) (*StartPaymentResp, error)
	CallBack(context.Context, *CallBackReq) (*AlipayCallbackNotificationResp, error)
	GetOrderPayemntStatus(context.Context, *GetOrderPayemntStatusReq) (*GetOrderPayemntStatusResp, error)
	mustEmbedUnimplementedPaymentServiceServer()
}

// UnimplementedPaymentServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPaymentServiceServer struct{}

func (UnimplementedPaymentServiceServer) Charge(context.Context, *ChargeReq) (*ChargeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Charge not implemented")
}
func (UnimplementedPaymentServiceServer) StartPayment(context.Context, *StartPaymentReq) (*StartPaymentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPayment not implemented")
}
func (UnimplementedPaymentServiceServer) CallBack(context.Context, *CallBackReq) (*AlipayCallbackNotificationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallBack not implemented")
}
func (UnimplementedPaymentServiceServer) GetOrderPayemntStatus(context.Context, *GetOrderPayemntStatusReq) (*GetOrderPayemntStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderPayemntStatus not implemented")
}
func (UnimplementedPaymentServiceServer) mustEmbedUnimplementedPaymentServiceServer() {}
func (UnimplementedPaymentServiceServer) testEmbeddedByValue()                        {}

// UnsafePaymentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentServiceServer will
// result in compilation errors.
type UnsafePaymentServiceServer interface {
	mustEmbedUnimplementedPaymentServiceServer()
}

func RegisterPaymentServiceServer(s grpc.ServiceRegistrar, srv PaymentServiceServer) {
	// If the following call pancis, it indicates UnimplementedPaymentServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PaymentService_ServiceDesc, srv)
}

func _PaymentService_Charge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChargeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).Charge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_Charge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).Charge(ctx, req.(*ChargeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_StartPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartPaymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).StartPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_StartPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).StartPayment(ctx, req.(*StartPaymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_CallBack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallBackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).CallBack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_CallBack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).CallBack(ctx, req.(*CallBackReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentService_GetOrderPayemntStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderPayemntStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentServiceServer).GetOrderPayemntStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PaymentService_GetOrderPayemntStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentServiceServer).GetOrderPayemntStatus(ctx, req.(*GetOrderPayemntStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PaymentService_ServiceDesc is the grpc.ServiceDesc for PaymentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaymentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "payment.PaymentService",
	HandlerType: (*PaymentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Charge",
			Handler:    _PaymentService_Charge_Handler,
		},
		{
			MethodName: "StartPayment",
			Handler:    _PaymentService_StartPayment_Handler,
		},
		{
			MethodName: "CallBack",
			Handler:    _PaymentService_CallBack_Handler,
		},
		{
			MethodName: "GetOrderPayemntStatus",
			Handler:    _PaymentService_GetOrderPayemntStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "payment.proto",
}
