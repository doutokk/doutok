package payment

import (
	"context"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	
)

type RPCClient interface {
	KitexClient() paymentservice.Client
	Service() string
	Charge(ctx context.Context, Req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error)
	StartPayment(ctx context.Context, Req *payment.StartPaymentReq, callOptions ...callopt.Option) (r *payment.StartPaymentResp, err error)
	CallBack(ctx context.Context, Req *payment.AlipayCallbackNotification, callOptions ...callopt.Option) (r *payment.AlipayCallbackNotificationResp, err error)
	GetOrderPayemntStatus(ctx context.Context, Req *payment.GetOrderPayemntStatusReq, callOptions ...callopt.Option) (r *payment.GetOrderPayemntStatusResp, err error)
	Cancel(ctx context.Context, Req *payment.CancelPaymentReq, callOptions ...callopt.Option) (r *payment.CancelPaymentResp, err error)
	DirectPayment(ctx context.Context, Req *payment.DirectPaymentReq, callOptions ...callopt.Option) (r *payment.DirectPaymentResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := paymentservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient paymentservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() paymentservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Charge(ctx context.Context, Req *payment.ChargeReq, callOptions ...callopt.Option) (r *payment.ChargeResp, err error) {
	return c.kitexClient.Charge(ctx, Req, callOptions...)
}

func (c *clientImpl) StartPayment(ctx context.Context, Req *payment.StartPaymentReq, callOptions ...callopt.Option) (r *payment.StartPaymentResp, err error) {
	return c.kitexClient.StartPayment(ctx, Req, callOptions...)
}

func (c *clientImpl) CallBack(ctx context.Context, Req *payment.AlipayCallbackNotification, callOptions ...callopt.Option) (r *payment.AlipayCallbackNotificationResp, err error) {
	return c.kitexClient.CallBack(ctx, Req, callOptions...)
}

func (c *clientImpl) GetOrderPayemntStatus(ctx context.Context, Req *payment.GetOrderPayemntStatusReq, callOptions ...callopt.Option) (r *payment.GetOrderPayemntStatusResp, err error) {
	return c.kitexClient.GetOrderPayemntStatus(ctx, Req, callOptions...)
}

func (c *clientImpl) Cancel(ctx context.Context, Req *payment.CancelPaymentReq, callOptions ...callopt.Option) (r *payment.CancelPaymentResp, err error) {
	return c.kitexClient.Cancel(ctx, Req, callOptions...)
}

func (c *clientImpl) DirectPayment(ctx context.Context, Req *payment.DirectPaymentReq, callOptions ...callopt.Option) (r *payment.DirectPaymentResp, err error) {
	return c.kitexClient.DirectPayment(ctx, Req, callOptions...)
}
