package service

import (
	"context"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
)

type GetOrderPayemntStatusService struct {
	ctx context.Context
}

// NewGetOrderPayemntStatusService new GetOrderPayemntStatusService
func NewGetOrderPayemntStatusService(ctx context.Context) *GetOrderPayemntStatusService {
	return &GetOrderPayemntStatusService{ctx: ctx}
}

// Run create note info
func (s *GetOrderPayemntStatusService) Run(req *payment.GetOrderPayemntStatusReq) (resp *payment.GetOrderPayemntStatusResp, err error) {
	// Finish your business logic.

	return
}
