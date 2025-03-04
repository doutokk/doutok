package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/doutokk/doutok/app/product/biz/dal/redis"
	"github.com/doutokk/doutok/app/product/constants"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/product/biz/dal/model"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
}

func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Redis key todo：常量类
	cacheKey := fmt.Sprintf(constants.ProductCategoryKeyPattern, req.CategoryName, req.Page, req.PageSize)

	// 尝试从 Redis 获取缓存数据
	cachedData, err := redis.RedisClient.Get(s.ctx, cacheKey).Result()
	if err == nil {
		// 缓存命中，反序列化数据并返回
		var cachedResp product.ListProductsResp
		if err := json.Unmarshal([]byte(cachedData), &cachedResp); err == nil {
			return &cachedResp, nil
		}
	}

	// 缓存未命中，查询数据库
	p := query.Product
	c := query.ProductCategory
	var products []*model.Product

	if req.CategoryName != "" {
		cat, err := c.Where(c.Name.Eq(req.CategoryName)).Preload(c.Products).First()
		if err != nil {
			klog.Infof("error: %v", err)
			return &product.ListProductsResp{Item: make([]*product.Product, 0)}, nil
		}
		products = make([]*model.Product, len(cat.Products))
		for i, prod := range cat.Products {
			products[i] = &prod
		}
	} else {
		products, err = p.Preload(p.Categories).
			Limit(int(req.PageSize)).Offset(int(req.PageSize * int64(req.Page-1))).Find()
		if err != nil {
			return nil, err
		}
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

	// 序列化响应并存入 Redis
	data, err := json.Marshal(resp)
	if err == nil {
		redis.RedisClient.Set(s.ctx, cacheKey, data, time.Minute*10)
	}

	return
}
