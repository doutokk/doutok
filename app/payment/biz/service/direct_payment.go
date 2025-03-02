package service

import (
	"context"
	"github.com/doutokk/doutok/app/payment/biz/fsm"
	"github.com/doutokk/doutok/app/payment/biz/pay"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
)

type DirectPaymentService struct {
	ctx context.Context
}

// NewDirectPaymentService new DirectPaymentService
func NewDirectPaymentService(ctx context.Context) *DirectPaymentService {
	return &DirectPaymentService{ctx: ctx}
}

// Run create note info
func (s *DirectPaymentService) Run(req *payment.DirectPaymentReq) (resp *payment.DirectPaymentResp, err error) {
	// Finish your business logic.

	oi := req.OutTradeNo

	os, err := fsm.RestoreFromDB(oi)
	if err != nil {
		return
	}
	check := os.DirectCheck(pay.ReturnCallbackParams{
		Charset:     req.Charset,
		OutTradeNo:  req.OutTradeNo,
		Method:      req.Method,
		TotalAmount: req.TotalAmount,
		Sign:        req.Sign,
		TradeNo:     req.TradeNo,
		AuthAppId:   req.AuthAppId,
		Version:     req.Version,
		AppId:       req.AppId,
		SignType:    req.SignType,
		SellerId:    req.SellerId,
		Timestamp:   req.Timestamp,
	})
	resp = &payment.DirectPaymentResp{
		Success: check,
		Message: "",
		TradeNo: req.TradeNo,
	}
	return
}
