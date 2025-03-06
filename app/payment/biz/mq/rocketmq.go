package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/payment/biz/interfaces"
	"github.com/doutokk/doutok/app/payment/conf"
)

// RocketMQClient encapsulates RocketMQ producer and consumer
type RocketMQClient struct {
	producer       rocketmq.Producer
	consumer       rocketmq.PushConsumer
	config         *conf.Config
	once           sync.Once
	orderCanceller interfaces.OrderCanceller
}

// OrderCancelMessage represents the message for order cancellation
type OrderCancelMessage struct {
	OrderID string `json:"order_id"`
}

var (
	mqClient *RocketMQClient
	mux      sync.Mutex
)

// GetMQClient returns a singleton RocketMQ client
func GetMQClient() *RocketMQClient {
	if mqClient == nil {
		mux.Lock()
		defer mux.Unlock()
		if mqClient == nil {
			mqClient = &RocketMQClient{
				config: conf.GetConf(),
			}
		}
	}
	return mqClient
}

// SetOrderCanceller sets the handler for order cancellation
func (c *RocketMQClient) SetOrderCanceller(canceller interfaces.OrderCanceller) {
	c.orderCanceller = canceller
}

// InitProducer initializes the RocketMQ producer
func (c *RocketMQClient) InitProducer() error {
	var err error
	c.once.Do(func() {
		opts := []producer.Option{
			producer.WithNameServer(c.config.RocketMQ.NameServer),
			producer.WithGroupName(c.config.RocketMQ.GroupName),
			producer.WithRetry(2),
		}

		// Add ACL if credentials are provided
		if c.config.RocketMQ.AccessKey != "" && c.config.RocketMQ.SecretKey != "" {
			opts = append(opts, producer.WithCredentials(primitive.Credentials{
				AccessKey: c.config.RocketMQ.AccessKey,
				SecretKey: c.config.RocketMQ.SecretKey,
			}))
		}

		c.producer, err = rocketmq.NewProducer(opts...)
		if err != nil {
			klog.Errorf("Failed to create RocketMQ producer: %v", err)
			return
		}

		if err = c.producer.Start(); err != nil {
			klog.Errorf("Failed to start RocketMQ producer: %v", err)
			return
		}

		klog.Infof("RocketMQ producer started successfully")
	})
	return err
}

// InitConsumer initializes the RocketMQ consumer
func (c *RocketMQClient) InitConsumer() error {
	var err error
	opts := []consumer.Option{
		consumer.WithNameServer(c.config.RocketMQ.NameServer),
		consumer.WithGroupName(c.config.RocketMQ.GroupName),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromLastOffset),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithMaxReconsumeTimes(int32(c.config.RocketMQ.MaxReconsumeTimes)),
	}

	// Add ACL if credentials are provided
	if c.config.RocketMQ.AccessKey != "" && c.config.RocketMQ.SecretKey != "" {
		opts = append(opts, consumer.WithCredentials(primitive.Credentials{
			AccessKey: c.config.RocketMQ.AccessKey,
			SecretKey: c.config.RocketMQ.SecretKey,
		}))
	}

	c.consumer, err = rocketmq.NewPushConsumer(opts...)
	if err != nil {
		klog.Errorf("Failed to create RocketMQ consumer: %v", err)
		return err
	}

	// Subscribe to the order cancellation topic
	err = c.consumer.Subscribe(c.config.RocketMQ.OrderCancelTopic, consumer.MessageSelector{}, c.handleOrderCancellation)
	if err != nil {
		klog.Errorf("Failed to subscribe to topic %s: %v", c.config.RocketMQ.OrderCancelTopic, err)
		return err
	}

	if err = c.consumer.Start(); err != nil {
		klog.Errorf("Failed to start RocketMQ consumer: %v", err)
		return err
	}

	klog.Infof("RocketMQ consumer started successfully")
	return nil
}

// SendDelayedOrderCancellation sends a delayed message for order cancellation
func (c *RocketMQClient) SendDelayedOrderCancellation(ctx context.Context, orderID string) error {
	if err := c.InitProducer(); err != nil {
		return err
	}

	message := OrderCancelMessage{
		OrderID: orderID,
	}

	data, err := json.Marshal(message)
	if err != nil {
		klog.Errorf("Failed to marshal order cancellation message: %v", err)
		return err
	}

	msg := &primitive.Message{
		Topic: c.config.RocketMQ.OrderCancelTopic,
		Body:  data,
	}

	// Set delay level (1s 5s 10s 30s 1m 2m 3m 4m 5m 6m 7m 8m 9m 10m 20m 30m 1h 2h)
	// Since we need 30 minutes, we can use level 16 (30m)
	msg.WithDelayTimeLevel(16)

	res, err := c.producer.SendSync(ctx, msg)
	if err != nil {
		klog.Errorf("Failed to send delayed order cancellation message: %v", err)
		return err
	}

	klog.Infof("Sent delayed cancellation for order %s, result: %s", orderID, res.String())
	return nil
}

// handleOrderCancellation processes order cancellation messages
func (c *RocketMQClient) handleOrderCancellation(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	for _, msg := range msgs {
		var message OrderCancelMessage
		if err := json.Unmarshal(msg.Body, &message); err != nil {
			klog.Errorf("Failed to unmarshal order cancellation message: %v", err)
			continue
		}

		klog.Infof("Processing delayed order cancellation for order ID: %s", message.OrderID)

		// Check if we have a canceller configured
		if c.orderCanceller == nil {
			klog.Errorf("No order canceller configured, cannot process cancellation for order %s", message.OrderID)
			return consumer.ConsumeRetryLater, fmt.Errorf("no order canceller configured")
		}

		// Call the order canceller
		err := c.orderCanceller.CancelOrder(ctx, message.OrderID)
		if err != nil {
			klog.Errorf("Failed to cancel order %s: %v", message.OrderID, err)
			// If it's a transient error, we might want to retry
			return consumer.ConsumeRetryLater, err
		}

		klog.Infof("Successfully processed cancellation for order %s", message.OrderID)
	}

	return consumer.ConsumeSuccess, nil
}

// Close shuts down the RocketMQ client
func (c *RocketMQClient) Close() {
	if c.producer != nil {
		c.producer.Shutdown()
	}
	if c.consumer != nil {
		c.consumer.Shutdown()
	}
}
