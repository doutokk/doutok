package service

import (
	"context"
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
func (s *CallBackService) Run(req *payment.CallBackReq) (resp *payment.AlipayCallbackNotificationResp, err error) {
	// Finish your business logic.

	return
}
