package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/product/biz/dal/model"
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
	c := query.ProductCategory
	var products []*model.Product
	// 还是不太会用 gorm，这里的逻辑是如果请求中有分类名，就根据分类名查找商品，否则查找所有商品
	// 没搞定直接多对多查询的，所以先查出分类，再查出关联的商品
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

	return
}
