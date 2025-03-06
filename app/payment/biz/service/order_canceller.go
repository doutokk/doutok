package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/payment/biz/fsm"
)

// OrderCancellerService provides methods to cancel orders
type OrderCancellerService struct{}

// NewOrderCancellerService creates a new OrderCancellerService
func NewOrderCancellerService() *OrderCancellerService {
	return &OrderCancellerService{}
}

// CancelOrder cancels an order by ID
func (s *OrderCancellerService) CancelOrder(ctx context.Context, orderID string) error {
	klog.Infof("Cancelling order with ID: %s", orderID)

	// Restore order FSM from database
	orderFSM, err := fsm.RestoreFromDB(orderID)
	if err != nil {
		klog.Errorf("Failed to restore order FSM: %v", err)
		return err
	}

	// Check current status - only cancel if not in FINISH state
	if orderFSM.GetStatus() != fsm.FINISH {
		// Cancel the order
		if err := orderFSM.CancelPayment(ctx); err != nil {
			klog.Errorf("Failed to cancel order: %v", err)
			return err
		}
		klog.Infof("Successfully cancelled order %s", orderID)
	} else {
		klog.Infof("Order %s is already in FINISH state, skipping cancellation", orderID)
	}

	return nil
}
