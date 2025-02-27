package service

import (
	"context"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

type GetProductBatchService struct {
	ctx context.Context
}

// NewGetProductBatchService new GetProductBatchService
func NewGetProductBatchService(ctx context.Context) *GetProductBatchService {
	return &GetProductBatchService{ctx: ctx}
}

// Run create note info
func (s *GetProductBatchService) Run(req *product.GetProductBatchReq) (resp *product.GetProductBatchResp, err error) {
	// Finish your business logic.

	p := query.Q.Product
	ids := make([]uint, 0, len(req.Ids))
	for _, id := range req.Ids {
		ids = append(ids, uint(id))
	}
	products, err := p.Where(p.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	resp = new(product.GetProductBatchResp)
	for _, m := range products {
		resp.Item = append(resp.Item, &product.Product{
			Id:          uint32(m.ID),
			Name:        m.Name,
			Description: m.Description,
			Picture:     m.Picture,
			Price:       m.Price,
		})
	}
	return
}
