package service

import (
	"context"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
}

// NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	p := query.Product
	var q query.IProductDo
	if req.CategoryName != "" {
		q = query.Q.Product.Where(query.ProductCategory.Name.In(req.CategoryName)).Preload(p.Categories)
	} else {
		q = query.Q.Product.Preload(p.Categories)
	}
	q.Limit(int(req.PageSize)).Offset(int(req.PageSize * int64(req.Page-1)))
	products, err := q.Find()
	if err != nil {
		return
	}
	resp = &product.ListProductsResp{Item: make([]*product.Product, len(products))}
	for i, prod := range products {
		cats := make([]string, len(prod.Categories))
		for i, cat := range prod.Categories {
			cats[i] = cat.Name
		}
		resp.Item[i] = &product.Product{
			Id:          uint32(prod.ID),
			Name:        prod.Name,
			Description: prod.Description,
			Picture:     prod.Picture,
			Price:       prod.Price,
			Categories:  cats,
		}
	}

	return
}
