package service

import (
	"context"
	"fmt"
	"github.com/doutokk/doutok/app/product/biz/dal/model"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
	"github.com/doutokk/doutok/app/product/biz/dal/redis"
	"github.com/doutokk/doutok/app/product/constants"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

type CreateProductService struct {
	ctx context.Context
}

func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
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

	// 数据库总条数
	totalCount, err := p.Count()
	if err != nil {
		return
	}

	// 清除最后一页的缓存
	pageSizes := []int{10, 20, 50, 100}
	for _, pageSize := range pageSizes {
		lastPage := (int(totalCount) + pageSize - 1) / pageSize
		cacheKey := fmt.Sprintf(constants.ProductCategoryKeyPattern, req.Categories[0], lastPage, pageSize)
		redis.RedisClient.Del(s.ctx, cacheKey)
	}

	return
}
