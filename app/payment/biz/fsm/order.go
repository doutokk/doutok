package fsm

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/payment/biz/interfaces"
	"github.com/doutokk/doutok/app/payment/biz/pay"
	"github.com/doutokk/doutok/app/payment/conf"
	"github.com/doutokk/doutok/common/lock"

	"github.com/doutokk/doutok/app/payment/biz/dal/model"
	"github.com/doutokk/doutok/app/payment/biz/dal/query"

	"github.com/looplab/fsm"
)

type PayOrderStatus string

const (
	CREATED   PayOrderStatus = "CREATED"
	PAYING    PayOrderStatus = "PAYING"
	FINISH    PayOrderStatus = "FINISH"
	CANCELLED PayOrderStatus = "CANCELLED" // New state for cancelled orders
)

type PayOrderEvent string

const (
	CreateOrder    PayOrderEvent = "CREATE_ORDER"
	StartPayment   PayOrderEvent = "START_PAYMENT"
	PaymentSuccess PayOrderEvent = "PAYMENT_SUCCESS"
	PaymentFailed  PayOrderEvent = "PAYMENT_FAILED"
	CancelOrder    PayOrderEvent = "CANCEL_ORDER" // New event for cancelling orders
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
			{Name: string(CancelOrder), Src: []string{string(CREATED), string(PAYING)}, Dst: string(CANCELLED)},
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

// Redis address should come from configuration
var redLock = lock.NewRedLock(conf.GetConf().Redis.Address)

// Global variable for the delayed message sender
var delayedMessageSender interfaces.DelayedMessageSender

// SetDelayedMessageSender sets the implementation for sending delayed messages
func SetDelayedMessageSender(sender interfaces.DelayedMessageSender) {
	delayedMessageSender = sender
}

// NewOrder creates a new payment order with distributed locking
func NewOrder(req CreatePayOrderReq) (*PayOrderFSM, error) {
	o := &PayOrderFSM{}
	o.fsm = fsm.NewFSM(
		string(CREATED),
		fsm.Events{
			{Name: string(StartPayment), Src: []string{string(CREATED)}, Dst: string(PAYING)},
			{Name: string(PaymentSuccess), Src: []string{string(PAYING)}, Dst: string(FINISH)},
			{Name: string(PaymentFailed), Src: []string{string(PAYING)}, Dst: string(CREATED)},
			{Name: string(CancelOrder), Src: []string{string(CREATED), string(PAYING)}, Dst: string(CANCELLED)},
		},
		fsm.Callbacks{},
	)

	// Acquire lock for this order
	lockKey := fmt.Sprintf("payment_order_lock:%s", req.OrderId)
	locked, err := redLock.TryLock(lockKey, 5*time.Second, 30*time.Second, 100*time.Millisecond, true)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire lock: %w", err)
	}
	if !locked {
		return nil, fmt.Errorf("could not acquire lock for order %s", req.OrderId)
	}

	defer func() {
		// Release lock when function returns
		redLock.Unlock(lockKey)
	}()

	l := query.Q.PaymentLog
	err = l.Create(&model.PaymentLog{
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

	// Schedule delayed order cancellation if we have a sender configured
	if delayedMessageSender != nil {
		if err := delayedMessageSender.SendDelayedOrderCancellation(context.Background(), req.OrderId); err != nil {
			klog.Warnf("Failed to schedule delayed cancellation for order %s: %v", req.OrderId, err)
			// Continue even if scheduling fails, as this is not critical for order creation
		}
	}

	return o, nil
}

// StartPayment processes the START_PAYMENT event
func (o *PayOrderFSM) StartPayment(ctx context.Context) (string, error) {
	// Acquire lock for this order
	lockKey := fmt.Sprintf("payment_order_lock:%s", o.orderId)
	locked, err := redLock.TryLock(lockKey, 5*time.Second, 30*time.Second, 100*time.Millisecond, true)
	if err != nil {
		return "", fmt.Errorf("failed to acquire lock: %w", err)
	}
	if !locked {
		return "", fmt.Errorf("could not acquire lock for order %s", o.orderId)
	}

	defer func() {
		// Release lock when function returns
		redLock.Unlock(lockKey)
	}()

	err = o.fsm.Event(ctx, string(StartPayment))

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
	// Acquire lock for this order
	lockKey := fmt.Sprintf("payment_order_lock:%s", o.orderId)
	locked, err := redLock.TryLock(lockKey, 5*time.Second, 30*time.Second, 100*time.Millisecond, true)
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %w", err)
	}
	if !locked {
		return fmt.Errorf("could not acquire lock for order %s", o.orderId)
	}

	defer func() {
		// Release lock when function returns
		redLock.Unlock(lockKey)
	}()

	err = o.fsm.Event(ctx, string(PaymentSuccess))
	l := query.Q.PaymentLog
	_, err = l.Where(l.OrderId.Eq(o.orderId)).Update(l.Status, string(FINISH))

	if err != nil {
		return fmt.Errorf("failed to process successful payment: %w", err)
	}
	return nil
}

// PaymentFailed processes the PAYMENT_FAILED event
func (o *PayOrderFSM) PaymentFailed(ctx context.Context) error {
	// Acquire lock for this order
	lockKey := fmt.Sprintf("payment_order_lock:%s", o.orderId)
	locked, err := redLock.TryLock(lockKey, 5*time.Second, 30*time.Second, 100*time.Millisecond, true)
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %w", err)
	}
	if !locked {
		return fmt.Errorf("could not acquire lock for order %s", o.orderId)
	}

	defer func() {
		// Release lock when function returns
		redLock.Unlock(lockKey)
	}()

	err = o.fsm.Event(ctx, string(PaymentFailed))
	l := query.Q.PaymentLog
	pay.CancelOrder(o.orderId)
	_, err = l.Where(l.OrderId.Eq(o.orderId)).Update(l.Status, string(CREATED))
	if err != nil {
		return fmt.Errorf("failed to process payment failure: %w", err)
	}
	return nil
}

// CancelPayment processes the CANCEL_ORDER event
func (o *PayOrderFSM) CancelPayment(ctx context.Context) error {
	previousState := o.fsm.Current()
	err := o.fsm.Event(ctx, string(CancelOrder))
	if err != nil {
		return fmt.Errorf("failed to change state to cancelled: %w", err)
	}

	// If the order was in PAYING state, cancel the payment with the payment provider
	if o.fsm.Current() == string(CANCELLED) && previousState == string(PAYING) {
		pay.CancelOrder(o.orderId)
	}

	// Update database
	l := query.Q.PaymentLog
	_, err = l.Where(l.OrderId.Eq(o.orderId)).Update(l.Status, string(CANCELLED))
	if err != nil {
		return fmt.Errorf("failed to update payment status to cancelled: %w", err)
	}

	return nil
}

func (o *PayOrderFSM) GetStatus() PayOrderStatus {
	return PayOrderStatus(o.fsm.Current())
}

func (o *PayOrderFSM) DirectCheck(params pay.ReturnCallbackParams) bool {
	fmt.Println("DirectCheck")

	// Add locking only for the state-changing operations
	if pay.VerifyReturnCallback(params) {
		// The call to PaymentSuccess already includes locking
		o.PaymentSuccess(context.Background())
		return true
	}

	if _, ok := pay.TradeQuery(context.Background(), params.TradeNo); ok {
		// The call to PaymentSuccess already includes locking
		o.PaymentSuccess(context.Background())
		return true
	}

	return false
}
