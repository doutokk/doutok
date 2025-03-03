package service

import (
	"context"

	"github.com/doutokk/doutok/app/product/biz/dal/query"
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
	// Finish your business logic.
	if req.Page <= 0 || req.PageSize <= 0 {
		req.Page = 1
		req.PageSize = 10
	}
	p := query.Product

	// Get total count first
	total, err := query.Q.Product.Where(p.Name.Like("%" + req.Query + "%")).Count()
	if err != nil {
		return
	}

	prods, err := query.Q.Product.Where(p.Name.Like("%" + req.Query + "%")).Preload(p.Categories).
		Limit(int(req.PageSize)).Offset(int(req.PageSize * int64(req.Page-1))).Find()
	if err != nil {
		return
	}
	resp = &product.SearchProductsResp{
		Item:  make([]*product.Product, len(prods)),
		Total: int32(total),
	}
	for i, prod := range prods {
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
