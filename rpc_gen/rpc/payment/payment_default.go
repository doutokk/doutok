package payment

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
)

func Charge(ctx context.Context, req *payment.ChargeReq, callOptions ...callopt.Option) (resp *payment.ChargeResp, err error) {
	resp, err = defaultClient.Charge(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Charge call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func StartPayment(ctx context.Context, req *payment.StartPaymentReq, callOptions ...callopt.Option) (resp *payment.StartPaymentResp, err error) {
	resp, err = defaultClient.StartPayment(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "StartPayment call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CallBack(ctx context.Context, req *payment.AlipayCallbackNotification, callOptions ...callopt.Option) (resp *payment.AlipayCallbackNotificationResp, err error) {
	resp, err = defaultClient.CallBack(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CallBack call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetOrderPayemntStatus(ctx context.Context, req *payment.GetOrderPayemntStatusReq, callOptions ...callopt.Option) (resp *payment.GetOrderPayemntStatusResp, err error) {
	resp, err = defaultClient.GetOrderPayemntStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetOrderPayemntStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Cancel(ctx context.Context, req *payment.CancelPaymentReq, callOptions ...callopt.Option) (resp *payment.CancelPaymentResp, err error) {
	resp, err = defaultClient.Cancel(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Cancel call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
