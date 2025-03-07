// Code generated by Kitex v0.9.1. DO NOT EDIT.

package paymentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Charge(ctx context.Context, Req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error)
	StartPayment(ctx context.Context, Req *payment.StartPaymentReq, callOptions ...callopt.Option) (r *payment.StartPaymentResp, err error)
	CallBack(ctx context.Context, Req *payment.AlipayCallbackNotification, callOptions ...callopt.Option) (r *payment.AlipayCallbackNotificationResp, err error)
	GetOrderPayemntStatus(ctx context.Context, Req *payment.GetOrderPayemntStatusReq, callOptions ...callopt.Option) (r *payment.GetOrderPayemntStatusResp, err error)
	Cancel(ctx context.Context, Req *payment.CancelPaymentReq, callOptions ...callopt.Option) (r *payment.CancelPaymentResp, err error)
	DirectPayment(ctx context.Context, Req *payment.DirectPaymentReq, callOptions ...callopt.Option) (r *payment.DirectPaymentResp, err error)
	CancelOrder(ctx context.Context, Req *payment.CancelOrderReq, callOptions ...callopt.Option) (r *payment.CancelOrderResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kPaymentServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kPaymentServiceClient struct {
	*kClient
}

func (p *kPaymentServiceClient) Charge(ctx context.Context, Req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Charge(ctx, Req)
}

func (p *kPaymentServiceClient) StartPayment(ctx context.Context, Req *payment.StartPaymentReq, callOptions ...callopt.Option) (r *payment.StartPaymentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.StartPayment(ctx, Req)
}

func (p *kPaymentServiceClient) CallBack(ctx context.Context, Req *payment.AlipayCallbackNotification, callOptions ...callopt.Option) (r *payment.AlipayCallbackNotificationResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CallBack(ctx, Req)
}

func (p *kPaymentServiceClient) GetOrderPayemntStatus(ctx context.Context, Req *payment.GetOrderPayemntStatusReq, callOptions ...callopt.Option) (r *payment.GetOrderPayemntStatusResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOrderPayemntStatus(ctx, Req)
}

func (p *kPaymentServiceClient) Cancel(ctx context.Context, Req *payment.CancelPaymentReq, callOptions ...callopt.Option) (r *payment.CancelPaymentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Cancel(ctx, Req)
}

func (p *kPaymentServiceClient) DirectPayment(ctx context.Context, Req *payment.DirectPaymentReq, callOptions ...callopt.Option) (r *payment.DirectPaymentResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DirectPayment(ctx, Req)
}

func (p *kPaymentServiceClient) CancelOrder(ctx context.Context, Req *payment.CancelOrderReq, callOptions ...callopt.Option) (r *payment.CancelOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CancelOrder(ctx, Req)
}
