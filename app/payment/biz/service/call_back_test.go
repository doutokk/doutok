package service

import (
	"context"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestCallBack_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCallBackService(ctx)
	// init req and assert value

	req := &payment.CallBackReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
