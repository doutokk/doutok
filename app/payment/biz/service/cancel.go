package service

import (
	"context"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
)

type CancelService struct {
	ctx context.Context
}

// NewCancelService new CancelService
func NewCancelService(ctx context.Context) *CancelService {
	return &CancelService{ctx: ctx}
}

// Run create note info
func (s *CancelService) Run(req *payment.CancelPaymentReq) (resp *payment.CancelPaymentResp, err error) {
	// Finish your business logic.

	return
}
