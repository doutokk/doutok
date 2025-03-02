package fsm

import (
	"context"
	"fmt"
	"github.com/doutokk/doutok/app/payment/biz/pay"
	//"github.com/doutokk/doutok/app/payment/infra/rpc"
	//"github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
	"time"

	"github.com/doutokk/doutok/app/payment/biz/dal/model"
	"github.com/doutokk/doutok/app/payment/biz/dal/query"

	"github.com/looplab/fsm"
)

type PayOrderStatus string

const (
	CREATED PayOrderStatus = "CREATED"
	PAYING  PayOrderStatus = "PAYING"
	FINISH  PayOrderStatus = "FINISH"
)

type PayOrderEvent string

const (
	CreateOrder    PayOrderEvent = "CREATE_ORDER"
	StartPayment   PayOrderEvent = "START_PAYMENT"
	PaymentSuccess PayOrderEvent = "PAYMENT_SUCCESS"
	PaymentFailed  PayOrderEvent = "PAYMENT_FAILED"
)

type PayOrderFSM struct {
	fsm     *fsm.FSM
	orderId string
	data    CreatePayOrderReq
}

type CreatePayOrderReq struct {
	UserId  uint32
	OrderId string
	Amount  float32
}

// 状态机似乎放错地方了，感觉放在 order 模块内更合适，或者说两边涉及到钱的，都应该有一个自己的状态机

func RestoreFromDB(orderId string) (*PayOrderFSM, error) {
	l := query.Q.PaymentLog
	paymentLog, err := l.Where(l.OrderId.Eq(orderId)).First()
	if err != nil {
		return nil, fmt.Errorf("failed to restore payment log: %w", err)
	}

	o := &PayOrderFSM{}
	o.fsm = fsm.NewFSM(
		paymentLog.Status,
		fsm.Events{
			{Name: string(StartPayment), Src: []string{string(CREATED)}, Dst: string(PAYING)},
			{Name: string(PaymentSuccess), Src: []string{string(PAYING)}, Dst: string(FINISH)},
			{Name: string(PaymentFailed), Src: []string{string(PAYING)}, Dst: string(CREATED)},
		},
		fsm.Callbacks{},
	)
	o.orderId = orderId
	//order, err := rpc.OrderClient.GetOrder(context.Background(), &order.GetOrderReq{Id: orderId})
	//if err != nil {
	//	return nil, fmt.Errorf("failed to get order: %w", err)
	//}
	o.data = CreatePayOrderReq{
		UserId:  uint32(paymentLog.UserId),
		OrderId: orderId,
		Amount:  paymentLog.Amount,
	}
	return o, nil
}

func NewOrder(req CreatePayOrderReq) (*PayOrderFSM, error) {
	o := &PayOrderFSM{}
	o.fsm = fsm.NewFSM(
		string(CREATED),
		fsm.Events{
			{Name: string(StartPayment), Src: []string{string(CREATED)}, Dst: string(PAYING)},
			{Name: string(PaymentSuccess), Src: []string{string(PAYING)}, Dst: string(FINISH)},
			{Name: string(PaymentFailed), Src: []string{string(PAYING)}, Dst: string(CREATED)},
		},
		fsm.Callbacks{},
	)

	l := query.Q.PaymentLog
	err := l.Create(&model.PaymentLog{
		UserId:        req.UserId,
		OrderId:       req.OrderId,
		TransactionId: "",
		Status:        string(CREATED),
		Amount:        req.Amount,
		PayAt:         time.Now(),
	})
	o.orderId = req.OrderId
	o.data = req
	if err != nil {
		return nil, fmt.Errorf("failed to create payment log: %w", err)
	}

	return o, nil
}

// StartPayment processes the START_PAYMENT event
func (o *PayOrderFSM) StartPayment(ctx context.Context) (string, error) {
	err := o.fsm.Event(ctx, string(StartPayment))

	l := query.Q.PaymentLog
	_, err = l.Where(l.OrderId.Eq(o.orderId)).Update(l.Status, string(PAYING))

	if err != nil {
		return "", fmt.Errorf("failed to start payment: %w", err)
	}
	url, err := pay.CreatePayOrder(o.orderId, float64(o.data.Amount))
	if err != nil {
		return "", err
	}

	return url, nil
}

// PaymentSuccess processes the PAYMENT_SUCCESS event
func (o *PayOrderFSM) PaymentSuccess(ctx context.Context) error {
	err := o.fsm.Event(ctx, string(PaymentSuccess))
	l := query.Q.PaymentLog
	_, err = l.Where(l.OrderId.Eq(o.orderId)).Update(l.Status, string(FINISH))

	if err != nil {
		return fmt.Errorf("failed to process successful payment: %w", err)
	}
	return nil
}

// PaymentFailed processes the PAYMENT_FAILED event
func (o *PayOrderFSM) PaymentFailed(ctx context.Context) error {
	err := o.fsm.Event(ctx, string(PaymentFailed))
	l := query.Q.PaymentLog
	pay.CancelOrder(o.orderId)
	_, err = l.Where(l.OrderId.Eq(o.orderId)).Update(l.Status, string(CREATED))
	if err != nil {
		return fmt.Errorf("failed to process payment failure: %w", err)
	}
	return nil
}

func (o *PayOrderFSM) GetStatus() PayOrderStatus {
	return PayOrderStatus(o.fsm.Current())
}
