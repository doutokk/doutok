package service

import (
	"context"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
	"github.com/doutokk/doutok/app/product/infra"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

type DeleteProductService struct {
	ctx context.Context
}

// NewDeleteProductService new DeleteProductService
func NewDeleteProductService(ctx context.Context) *DeleteProductService {
	return &DeleteProductService{ctx: ctx}
}

// Run create note info
func (s *DeleteProductService) Run(req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	// Finish your business logic.

	p := query.Q.Product
	_, err = p.Where(p.ID.Eq(uint(req.Id))).Delete()
	if err != nil {
		return &product.DeleteProductResp{Success: false}, err
	}

	// 删除 Elasticsearch 中的商品
	if err := infra.DeleteProduct(s.ctx, req.Id); err != nil {
		return &product.DeleteProductResp{Success: false}, err
	}

	resp = &product.DeleteProductResp{
		Success: true,
	}
	return
}
