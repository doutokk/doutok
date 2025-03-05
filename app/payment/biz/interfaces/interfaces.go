package interfaces

import "context"

// OrderCanceller defines the interface for order cancellation services
type OrderCanceller interface {
	CancelOrder(ctx context.Context, orderID string) error
}

// DelayedMessageSender defines the interface for sending delayed messages
type DelayedMessageSender interface {
	SendDelayedOrderCancellation(ctx context.Context, orderID string) error
}
