package main

import (
	"context"
	"github.com/doutokk/doutok/app/payment/biz/service"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}

// StartPayment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) StartPayment(ctx context.Context, req *payment.StartPaymentReq) (resp *payment.StartPaymentResp, err error) {
	resp, err = service.NewStartPaymentService(ctx).Run(req)

	return resp, err
}

// CallBack implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) CallBack(ctx context.Context, req *payment.AlipayCallbackNotification) (resp *payment.AlipayCallbackNotificationResp, err error) {
	resp, err = service.NewCallBackService(ctx).Run(req)

	return resp, err
}

// GetOrderPayemntStatus implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) GetOrderPayemntStatus(ctx context.Context, req *payment.GetOrderPayemntStatusReq) (resp *payment.GetOrderPayemntStatusResp, err error) {
	resp, err = service.NewGetOrderPayemntStatusService(ctx).Run(req)

	return resp, err
}
