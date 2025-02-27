package service

import (
	"context"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
)

type StartPaymentService struct {
	ctx context.Context
}

// NewStartPaymentService new StartPaymentService
func NewStartPaymentService(ctx context.Context) *StartPaymentService {
	return &StartPaymentService{ctx: ctx}
}

// Run create note info
func (s *StartPaymentService) Run(req *payment.StartPaymentReq) (resp *payment.StartPaymentResp, err error) {
	// Finish your business logic.

	return
}
