package service

import (
	"context"
	"testing"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
)

func TestCancelOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCancelOrderService(ctx)
	// init req and assert value

	req := &payment.CancelOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
