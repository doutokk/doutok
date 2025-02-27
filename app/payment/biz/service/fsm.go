package service

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

func main() {
	order := NewOrder()
	fmt.Println("Current Status:", order.FSM.Current())

	err := order.FSM.Event(context.Background(), string(StartPayment))
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Current Status:", order.FSM.Current())

	err = order.FSM.Event(context.Background(), string(PaymentSuccess))
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Current Status:", order.FSM.Current())
}
