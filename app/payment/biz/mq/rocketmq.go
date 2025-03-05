package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/payment/conf"
)

const (
	TopicOrderAutoCancel = "topic_order_auto_cancel"
)

// OrderCancelMessage represents the message structure for order cancellation
type OrderCancelMessage struct {
	OrderID string `json:"order_id"`
	UserID  uint32 `json:"user_id"`
}

var (
	rocketProducer rocketmq.Producer
	rocketConsumer rocketmq.PushConsumer
)

// Initialize sets up RocketMQ producer and consumer
func Initialize() error {
	// Initialize producer
	var err error
	rocketProducer, err = rocketmq.NewProducer(
		producer.WithNameServer([]string{conf.GetConf().RocketMQ.NamesrvAddr}),
		producer.WithGroupName(conf.GetConf().RocketMQ.GroupID),
		producer.WithRetry(2),
	)
	if err != nil {
		return fmt.Errorf("failed to create RocketMQ producer: %w", err)
	}

	err = rocketProducer.Start()
	if err != nil {
		return fmt.Errorf("failed to start RocketMQ producer: %w", err)
	}

	klog.Info("RocketMQ producer started successfully")
	return nil
}

// CleanUp properly shuts down RocketMQ clients
func CleanUp() {
	if rocketProducer != nil {
		err := rocketProducer.Shutdown()
		if err != nil {
			klog.Errorf("Error shutting down RocketMQ producer: %v", err)
		}
	}

	if rocketConsumer != nil {
		err := rocketConsumer.Shutdown()
		if err != nil {
			klog.Errorf("Error shutting down RocketMQ consumer: %v", err)
		}
	}
}

// SendOrderCancelDelayedMessage sends a delayed message to auto-cancel an order after the delay period
func SendOrderCancelDelayedMessage(orderID string, userID uint32, delay time.Duration) error {
	msg := &OrderCancelMessage{
		OrderID: orderID,
		UserID:  userID,
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal order cancel message: %w", err)
	}

	message := primitive.NewMessage(TopicOrderAutoCancel, body)

	// Set the delay level (RocketMQ uses predefined delay levels)
	// 1s, 5s, 10s, 30s, 1m, 2m, 3m, 4m, 5m, 6m, 7m, 8m, 9m, 10m, 20m, 30m, 1h, 2h
	// We want 30m, which is level 16
	message.WithDelayTimeLevel(16) // 30 minutes

	_, err = rocketProducer.SendSync(context.Background(), message)
	if err != nil {
		return fmt.Errorf("failed to send delayed message: %w", err)
	}

	klog.Infof("Sent delayed message for order auto-cancellation: OrderID=%s, UserID=%d", orderID, userID)
	return nil
}

// StartOrderCancelConsumer initializes and starts the consumer for order auto-cancellation
func StartOrderCancelConsumer(handleOrderCancel func(context.Context, string, uint32) error) error {
	var err error
	rocketConsumer, err = rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{conf.GetConf().RocketMQ.NamesrvAddr}),
		consumer.WithGroupName(conf.GetConf().RocketMQ.GroupID),
		consumer.WithConsumerModel(consumer.Clustering),
	)
	if err != nil {
		return fmt.Errorf("failed to create RocketMQ consumer: %w", err)
	}

	err = rocketConsumer.Subscribe(TopicOrderAutoCancel, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			var message OrderCancelMessage
			err := json.Unmarshal(msg.Body, &message)
			if err != nil {
				klog.Errorf("Failed to unmarshal order cancel message: %v", err)
				continue
			}

			klog.Infof("Received order auto-cancellation message: OrderID=%s, UserID=%d", message.OrderID, message.UserID)

			err = handleOrderCancel(ctx, message.OrderID, message.UserID)
			if err != nil {
				klog.Errorf("Failed to auto-cancel order %s: %v", message.OrderID, err)
				return consumer.ConsumeRetryLater, nil
			}
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		return fmt.Errorf("failed to subscribe to topic: %w", err)
	}

	err = rocketConsumer.Start()
	if err != nil {
		return fmt.Errorf("failed to start RocketMQ consumer: %w", err)
	}

	klog.Info("RocketMQ consumer started successfully")
	return nil
}
