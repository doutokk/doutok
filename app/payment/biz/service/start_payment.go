package service

import (
	"context"
	"fmt"
	"github.com/doutokk/doutok/app/payment/biz/fsm"
	"github.com/doutokk/doutok/app/payment/infra/rpc"
	"github.com/doutokk/doutok/common/utils"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
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

	fmt.Printf("StartPaymentService is called with req: %+v\n", req)
	userId := utils.GetUserId(s.ctx)
	oi := req.OrderId
	amount := 0
	rpc.OrderClient.GetOrder(s.ctx, &order.GetOrderReq{Id: oi})
	fsm.NewOrder(fsm.CreatePayOrderReq{
		UserId:  uint32(userId),
		OrderId: req.OrderId,
		Amount:  float32(amount),
	})

	return
}
