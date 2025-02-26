package service

import (
	"context"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
	"testing"
)

func TestEditProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEditProductService(ctx)
	// init req and assert value

	req := &product.EditProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
