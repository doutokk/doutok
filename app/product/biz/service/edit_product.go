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

type EditProductService struct {
	ctx context.Context
}

// NewEditProductService new EditProductService
func NewEditProductService(ctx context.Context) *EditProductService {
	return &EditProductService{ctx: ctx}
}

func (s *EditProductService) Run(req *product.EditProductReq) (resp *product.EditProductResp, err error) {
	product := req.Product

	// 数据库更新
	_, err = query.Product.WithContext(s.ctx).Where(query.Product.ID.Eq(uint(product.Id))).Updates(model.Product{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Picture:     product.Picture,
	})
	if err != nil {
		return
	}

	// Get the total count of products
	totalCount, err := query.Product.Count()
	if err != nil {
		return
	}
	// Clear the cache for the last page of each page size
	pageSizes := []int{10, 20, 50, 100}
	for _, pageSize := range pageSizes {
		lastPage := (int(totalCount) + pageSize - 1) / pageSize
		cacheKey := fmt.Sprintf(constants.ProductCategoryKeyPattern, product.Categories[0], lastPage, pageSize)
		redis.RedisClient.Del(s.ctx, cacheKey)
	}

	return
}
