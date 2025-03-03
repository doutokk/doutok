package service

import (
	"context"
	"errors"

	"github.com/doutokk/doutok/app/product/biz/dal/model"
	"github.com/doutokk/doutok/app/product/biz/dal/query"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type CreateProductService struct {
	ctx context.Context
}

// NewCreateProductService new CreateProductService
func NewCreateProductService(ctx context.Context) *CreateProductService {
	return &CreateProductService{ctx: ctx}
}

// Run create note info
func (s *CreateProductService) Run(req *product.CreateProductReq) (resp *product.CreateProductResp, err error) {
	// Finish your business logic.

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

	return resp, nil
}
