package service

import (
	"context"
	order "github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
)

type GetOrderService struct {
	ctx context.Context
} // NewGetOrderService new GetOrderService
func NewGetOrderService(ctx context.Context) *GetOrderService {
	return &GetOrderService{ctx: ctx}
}

// Run create note info
func (s *GetOrderService) Run(req *order.GetOrderReq) (resp *order.GetOrderResp, err error) {
	// Finish your business logic.

	return
}
