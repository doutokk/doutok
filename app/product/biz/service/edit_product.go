package service

import (
	"context"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

type EditProductService struct {
	ctx context.Context
}

// NewEditProductService new EditProductService
func NewEditProductService(ctx context.Context) *EditProductService {
	return &EditProductService{ctx: ctx}
}

// Run create note info
func (s *EditProductService) Run(req *product.EditProductReq) (resp *product.EditProductResp, err error) {
	// Finish your business logic.

	return
}
