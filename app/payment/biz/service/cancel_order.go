package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/doutokk/doutok/app/payment/infra/rpc"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/order"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/payment/biz/fsm"
	"github.com/doutokk/doutok/common/utils"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
)

type CancelOrderService struct {
	ctx context.Context
}

// NewCancelOrderService new CancelOrderService
func NewCancelOrderService(ctx context.Context) *CancelOrderService {
	return &CancelOrderService{ctx: ctx}
}

// Run create note info
func (s *CancelOrderService) Run(req *payment.CancelOrderReq) (resp *payment.CancelOrderResp, err error) {
	// Get user ID from context
	userId := utils.GetUserId(&s.ctx)
	oi := req.OrderId
	r, err := rpc.OrderClient.GetOrder(s.ctx, &order.GetOrderReq{Id: oi})
	if err != nil {
		return
	}
	if r.Order.UserId != uint32(userId) {
		klog.Warn("用户名和订单 id 不匹配")
		return nil, errors.New("")
	}

	fmt.Printf("CancelOrderService is called with req: %+v\n", req)

	// Restore order FSM from database
	var orderFSM *fsm.PayOrderFSM
	orderFSM, err = fsm.RestoreFromDB(req.OrderId)
	if err != nil {

		amount := 0.0
		for _, item := range r.Order.OrderItems {
			amount += float64(item.Cost)
		}
		orderFSM, err = fsm.NewOrder(fsm.CreatePayOrderReq{
			UserId:  userId,
			OrderId: req.OrderId,
			Amount:  float32(amount),
		})

		//klog.Errorf("Failed to restore order FSM: %v", err)
		//return &payment.CancelOrderResp{
		//	Success: false,
		//	Message: "Order not found",
		//}, err
	}

	// Cancel the order
	err = orderFSM.CancelPayment(s.ctx)
	if err != nil {
		klog.Errorf("Failed to cancel order: %v", err)
		return &payment.CancelOrderResp{
			Success: false,
			Message: fmt.Sprintf("Failed to cancel order: %v", err),
		}, err
	}

	// Return success response
	resp = &payment.CancelOrderResp{
		Success:     true,
		Message:     "Order cancelled successfully",
		OrderStatus: string(orderFSM.GetStatus()),
	}

	return resp, nil
}
