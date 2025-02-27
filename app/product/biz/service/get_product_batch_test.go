package service

import (
	"context"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
	"testing"
)

func TestGetProductBatch_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetProductBatchService(ctx)
	// init req and assert value

	req := &product.GetProductBatchReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
