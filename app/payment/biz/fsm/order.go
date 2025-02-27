package fsm

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

type OrderStatus string

const (
	CREATED OrderStatus = "CREATED"
	PAYING  OrderStatus = "PAYING"
	FINISH  OrderStatus = "FINISH"
)

type OrderEvent string

const (
	CreateOrder    OrderEvent = "CREATE_ORDER"
	StartPayment   OrderEvent = "START_PAYMENT"
	PaymentSuccess OrderEvent = "PAYMENT_SUCCESS"
	PaymentFailed  OrderEvent = "PAYMENT_FAILED"
)

type OrderFSM struct {
	FSM *fsm.FSM
}

func NewOrder() *OrderFSM {
	o := &OrderFSM{}
	o.FSM = fsm.NewFSM(
		string(CREATED),
		fsm.Events{
			{Name: string(StartPayment), Src: []string{string(CREATED)}, Dst: string(PAYING)},
			{Name: string(PaymentSuccess), Src: []string{string(PAYING)}, Dst: string(FINISH)},
			{Name: string(PaymentFailed), Src: []string{string(PAYING)}, Dst: string(CREATED)},
		},
		fsm.Callbacks{},
	)
	return o
}

// CreateOrder processes the CREATE_ORDER event
func (o *OrderFSM) CreateOrder(ctx context.Context) error {
	// This is typically the initial state, so no transition is needed
	// Could add initialization logic here if needed
	return nil
}

// StartPayment processes the START_PAYMENT event
func (o *OrderFSM) StartPayment(ctx context.Context) error {
	err := o.FSM.Event(ctx, string(StartPayment))
	if err != nil {
		return fmt.Errorf("failed to start payment: %w", err)
	}
	return nil
}

// PaymentSuccess processes the PAYMENT_SUCCESS event
func (o *OrderFSM) PaymentSuccess(ctx context.Context) error {
	err := o.FSM.Event(ctx, string(PaymentSuccess))
	if err != nil {
		return fmt.Errorf("failed to process successful payment: %w", err)
	}
	return nil
}

// PaymentFailed processes the PAYMENT_FAILED event
func (o *OrderFSM) PaymentFailed(ctx context.Context) error {
	err := o.FSM.Event(ctx, string(PaymentFailed))
	if err != nil {
		return fmt.Errorf("failed to process payment failure: %w", err)
	}
	return nil
}
