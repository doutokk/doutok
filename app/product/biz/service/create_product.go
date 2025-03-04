package service

import (
	"context"
	"fmt"

	"errors"

	"github.com/doutokk/doutok/app/product/biz/dal/model"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
	"github.com/doutokk/doutok/app/product/biz/dal/redis"
	"github.com/doutokk/doutok/app/product/constants"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type CreateProductService struct {
	ctx context.Context
}

func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	p := query.Q.Product
	pc := query.Q.ProductCategory
	cats := make([]model.ProductCategory, 0)

	for _, catName := range req.Categories {
		// 先查询分类是否存在
		existingCat, errCat := pc.Where(pc.Name.Eq(catName)).First()
		if errCat == nil {
			// 分类已存在，使用现有分类
			cats = append(cats, *existingCat)
		} else if errors.Is(errCat, gorm.ErrRecordNotFound) {
			// 分类不存在，创建新分类
			newCat := &model.ProductCategory{
				Name: catName,
			}
			errCreate := pc.Create(newCat)
			if errCreate != nil {
				return nil, errCreate
			}
			cats = append(cats, *newCat)
		} else {
			// 其他错误
			return nil, errCat
		}
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
		return nil, err
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
	return resp, nil
}
