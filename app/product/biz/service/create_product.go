package service

import (
	"context"
	"github.com/doutokk/doutok/app/product/biz/dal/model"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

type CreateProductService struct {
	ctx context.Context
}

// NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// Finish your business logic.

	p := query.Q.Product
	cats := make([]model.ProductCategory, 0)
	for _, cat := range req.Categories {
		cats = append(cats, model.ProductCategory{
			Name: cat,
		})
	}
	m := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Picture:     req.Picture,
		Price:       req.Price,
		Categories:  cats,
	}
	err = p.Create(m)
	if err != nil {
		return
	}
	resp = &product.CreateProductResp{
		Id: uint32(m.ID),
	}

	return
}
