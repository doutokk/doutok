package service

import (
	"context"
	"errors"
	"fmt"

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

	// Check if the user ID in request matches with context
	if req.UserId != uint32(userId) {
		klog.Warn("User ID mismatch in cancel order request")
		return &payment.CancelOrderResp{
			Success: false,
			Message: "User ID mismatch",
		}, errors.New("unauthorized")
	}

	fmt.Printf("CancelOrderService is called with req: %+v\n", req)

	// Restore order FSM from database
	var orderFSM *fsm.PayOrderFSM
	orderFSM, err = fsm.RestoreFromDB(req.OrderId)
	if err != nil {
		klog.Errorf("Failed to restore order FSM: %v", err)
		return &payment.CancelOrderResp{
			Success: false,
			Message: "Order not found",
		}, err
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
