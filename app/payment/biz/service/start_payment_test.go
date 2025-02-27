package service

import (
	"context"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestStartPayment_Run(t *testing.T) {
	ctx := context.Background()
	s := NewStartPaymentService(ctx)
	// init req and assert value

	req := &payment.StartPaymentReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
