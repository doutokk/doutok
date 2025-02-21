package service

import (
	"context"
	cart "github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"
	"testing"
)

func TestEditCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEditCartService(ctx)
	// init req and assert value

	req := &cart.EditCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
