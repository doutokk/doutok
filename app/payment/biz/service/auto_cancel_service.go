package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/payment/biz/dal/query"
	"github.com/doutokk/doutok/app/payment/biz/fsm"
)

// AutoCancelOrderService handles automatic cancellation of orders
type AutoCancelOrderService struct {
}

// NewAutoCancelOrderService creates a new instance of AutoCancelOrderService
func NewAutoCancelOrderService() *AutoCancelOrderService {
	return &AutoCancelOrderService{}
}

// HandleOrderCancel processes the auto-cancellation of an order
func (s *AutoCancelOrderService) HandleOrderCancel(ctx context.Context, orderID string, userID uint32) error {
	klog.Infof("Processing auto-cancellation for order %s", orderID)

	// Check current order status in the database
	l := query.Q.PaymentLog
	paymentLog, err := l.Where(l.OrderId.Eq(orderID)).First()
	if err != nil {
		return fmt.Errorf("failed to find payment log for order %s: %w", orderID, err)
	}

	// If the order is already in FINISH state, do nothing
	if paymentLog.Status == string(fsm.FINISH) {
		klog.Infof("Order %s is already in FINISH state, skipping auto-cancellation", orderID)
		return nil
	}

	// If order is in CANCELLED state, do nothing
	if paymentLog.Status == string(fsm.CANCELLED) {
		klog.Infof("Order %s is already in CANCELLED state, skipping auto-cancellation", orderID)
		return nil
	}

	// Restore the FSM and cancel the order
	orderFSM, err := fsm.RestoreFromDB(orderID)
	if err != nil {
		return fmt.Errorf("failed to restore order FSM for auto-cancellation: %w", err)
	}

	err = orderFSM.CancelPayment(ctx)
	if err != nil {
		return fmt.Errorf("failed to auto-cancel order %s: %w", orderID, err)
	}

	klog.Infof("Successfully auto-cancelled order %s after timeout", orderID)
	return nil
}
