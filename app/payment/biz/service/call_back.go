package service

import (
	"context"
	"fmt"
	"github.com/doutokk/doutok/app/payment/biz/fsm"
	"github.com/doutokk/doutok/app/payment/biz/pay"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
)

type CallBackService struct {
	ctx context.Context
}

// NewCallBackService new CallBackService
func NewCallBackService(ctx context.Context) *CallBackService {
	return &CallBackService{ctx: ctx}
}

// Run create note info
func (s *CallBackService) Run(req *payment.AlipayCallbackNotification) (resp *payment.AlipayCallbackNotificationResp, err error) {
	// Finish your business logic.

	fmt.Printf("CallBackService is called with req: %+v\n", req)

	pay.VerifyNotifyCallback(s.ctx, pay.AlipayCallbackNotification{
		GmtCreate:      req.GmtCreate,
		Charset:        req.Charset,
		GmtPayment:     req.GmtPayment,
		NotifyTime:     req.NotifyTime,
		Subject:        req.Subject,
		Sign:           req.Sign,
		BuyerId:        req.BuyerId,
		InvoiceAmount:  req.InvoiceAmount,
		Version:        req.Version,
		NotifyId:       req.NotifyId,
		FundBillList:   req.FundBillList,
		NotifyType:     req.NotifyType,
		OutTradeNo:     req.OutTradeNo,
		TotalAmount:    req.TotalAmount,
		TradeStatus:    req.TradeStatus,
		TradeNo:        req.TradeNo,
		AuthAppId:      req.AuthAppId,
		ReceiptAmount:  req.ReceiptAmount,
		PointAmount:    req.PointAmount,
		BuyerPayAmount: req.BuyerPayAmount,
		AppId:          req.AppId,
		SignType:       req.SignType,
		SellerId:       req.SellerId,
	})

	orderFSM, err := fsm.RestoreFromDB(req.OutTradeNo)
	err = orderFSM.PaymentSuccess(s.ctx)
	if err != nil {
		return nil, err
	}

	return &payment.AlipayCallbackNotificationResp{
		Success: "true",
	}, nil
}
