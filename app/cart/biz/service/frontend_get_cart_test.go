package service

import (
	"context"
	cart "github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"
	"testing"
)

func TestFrontendGetCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFrontendGetCartService(ctx)
	// init req and assert value

	req := &cart.FrontendGetCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
