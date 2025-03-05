package service

import (
	"context"
	"github.com/doutokk/doutok/app/product/biz/dal/model"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
	"github.com/doutokk/doutok/app/product/infra"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

type EditProductService struct {
	ctx context.Context
}

// NewEditProductService new EditProductService
func NewEditProductService(ctx context.Context) *EditProductService {
	return &EditProductService{ctx: ctx}
}

func (s *EditProductService) Run(req *product.EditProductReq) (resp *product.EditProductResp, err error) {
	prod := req.Product

	// 数据库更新
	_, err = query.Product.WithContext(s.ctx).Where(query.Product.ID.Eq(uint(prod.Id))).Updates(model.Product{
		Name:        prod.Name,
		Price:       prod.Price,
		Description: prod.Description,
		Picture:     prod.Picture,
	})
	if err != nil {
		return
	}

	// 更新 Elasticsearch 中的商品
	esProduct := &product.Product{
		Id:          prod.Id,
		Name:        prod.Name,
		Description: prod.Description,
		Picture:     prod.Picture,
		Price:       prod.Price,
		Categories:  prod.Categories,
	}
	if err := infra.UpdateProduct(s.ctx, esProduct); err != nil {
		return nil, err
	}
	return
}
