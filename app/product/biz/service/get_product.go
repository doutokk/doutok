package service

import (
	"context"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
	product "github.com/doutokk/doutok/app/product/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
}

// NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	p := query.Product
	prod, err := query.Q.Product.Where(p.ID.Eq(uint(req.Id))).Preload(p.Categories).First()
	if err != nil {
		return
	}
	cats := make([]string, len(prod.Categories))
	for i, cat := range prod.Categories {
		cats[i] = cat.Name
	}
	resp = &product.GetProductResp{Product: &product.Product{
		Id:          uint32(prod.ID),
		Name:        prod.Name,
		Description: prod.Description,
		Picture:     prod.Picture,
		Price:       prod.Price,
		Categories:  cats,
	}}
	return
}
