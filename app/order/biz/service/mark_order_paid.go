package service

import (
	"context"
	"github.com/doutokk/doutok/app/order/biz/dal/query"
	order "github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
)

type MarkOrderPaidService struct {
	ctx context.Context
}

// NewMarkOrderPaidService new MarkOrderPaidService
func NewMarkOrderPaidService(ctx context.Context) *MarkOrderPaidService {
	return &MarkOrderPaidService{ctx: ctx}
}

// Run create note info
func (s *MarkOrderPaidService) Run(req *order.MarkOrderPaidReq) (resp *order.MarkOrderPaidResp, err error) {
	// Finish your business logic.
	o := query.Q.Order
	_, err = query.Q.Order.Where(o.OrderID.Eq(req.OrderId)).Update(o.PaidStatus, true)
	if err != nil {
		return nil, err
	}

	return
}
