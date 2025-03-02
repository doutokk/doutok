package service

import (
	"context"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestCancel_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCancelService(ctx)
	// init req and assert value

	req := &payment.CancelPaymentReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
