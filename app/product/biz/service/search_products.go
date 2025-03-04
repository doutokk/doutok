package service

import (
	"context"
	"github.com/doutokk/doutok/app/product/infra"

	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
}

// NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	if req.Page <= 0 || req.PageSize <= 0 {
		req.Page = 1
		req.PageSize = 10
	}
	return infra.SearchProducts(s.ctx, req.Query, "", req.Page, int32(req.PageSize))
}
