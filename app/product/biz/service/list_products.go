package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/product/biz/dal/model"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
	"github.com/doutokk/doutok/app/product/biz/dal/redis"
	"github.com/doutokk/doutok/app/product/constants"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
}

func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Redis key
	cacheKey := fmt.Sprintf(constants.ProductCategoryKeyPattern, req.CategoryName, req.Page, req.PageSize)

	// Try to get cached data from Redis
	cachedData, err := redis.RedisClient.Get(s.ctx, cacheKey).Result()
	if err == nil {
		// Cache hit, deserialize data and return
		var cachedResp product.ListProductsResp
		if err := json.Unmarshal([]byte(cachedData), &cachedResp); err == nil {
			return &cachedResp, nil
		}
	}

	// Cache miss, query the database
	p := query.Product
	c := query.ProductCategory
	var products []*model.Product
	var total int64

	if req.CategoryName != "" {
		cat, err := c.Where(c.Name.Eq(req.CategoryName)).Preload(c.Products).First()
		if err != nil {
			klog.Infof("error: %v", err)
			return &product.ListProductsResp{
				Item:  make([]*product.Product, 0),
				Total: 0,
			}, nil
		}
		products = make([]*model.Product, len(cat.Products))
		for i, prod := range cat.Products {
			products[i] = &prod
		}
		total = int64(len(cat.Products))
	} else {
		// Get total count first
		total, err = p.Count()
		if err != nil {
			return nil, err
		}

		products, err = p.Preload(p.Categories).
			Limit(int(req.PageSize)).Offset(int(req.PageSize * int64(req.Page-1))).Find()
		if err != nil {
			return nil, err
		}
	}

	resp = &product.ListProductsResp{
		Item:  make([]*product.Product, len(products)),
		Total: int32(total),
	}
	for i, prod := range products {
		cats := make([]string, len(prod.Categories))
		for j, cat := range prod.Categories {
			cats[j] = cat.Name
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

	// Serialize response and store in Redis
	data, err := json.Marshal(resp)
	if err == nil {
		redis.RedisClient.Set(s.ctx, cacheKey, data, constants.Expire)
	}

	return resp, nil
}
