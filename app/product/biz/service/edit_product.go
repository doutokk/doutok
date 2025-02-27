package service

import (
	"context"
	"github.com/doutokk/doutok/app/product/biz/dal/model"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
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

	product := req.Product

	// 把更改插入数据库
	_, err = query.Product.WithContext(s.ctx).Where(query.Product.ID.Eq(uint(product.Id))).Updates(model.Product{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Picture:     product.Picture,
	})

	// 把整个对象跟数据库替换

	return
}
