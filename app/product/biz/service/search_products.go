package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/doutokk/doutok/app/product/biz/dal/redis"
	"github.com/doutokk/doutok/app/product/constants"
	"github.com/doutokk/doutok/app/product/infra"

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
	if req.Page <= 0 || req.PageSize <= 0 {
		req.Page = 1
		req.PageSize = 10
	}

	// 首页是调用的这里的接口，query是null，因此加上缓存
	if req.Query == "" {
		// Generate cache key
		cacheKey := fmt.Sprintf(constants.ProductKeyPattern, req.Page, req.PageSize)

		// Try to get cached response
		cachedResp, err := redis.RedisClient.Get(s.ctx, cacheKey).Result()
		if err == nil {
			// Unmarshal cached response
			var cachedResult product.SearchProductsResp
			if err := json.Unmarshal([]byte(cachedResp), &cachedResult); err == nil {
				return &cachedResult, nil
			}
		}

		// If cache miss or unmarshal error, proceed with search
		resp, err = infra.SearchProducts(s.ctx, req.Query, "", req.Page, int32(req.PageSize))
		if err != nil {
			return nil, err
		}

		// Cache the response
		respJSON, err := json.Marshal(resp)
		if err == nil {
			redis.RedisClient.Set(s.ctx, cacheKey, respJSON, constants.Expire)
		}

		return resp, nil
	}

	return infra.SearchProducts(s.ctx, req.Query, "", req.Page, int32(req.PageSize))
}
