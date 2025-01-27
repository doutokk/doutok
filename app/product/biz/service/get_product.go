package service

import (
	"context"
	"errors"
	"github.com/PengJingzhao/douyin-commerce/app/product/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/app/product/biz/model"
	product "github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// 查询单个商品
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	var productModel model.Product

	// Find product by id
	if err := mysql.DB.Preload("Categories").First(&productModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Map product model to RPC product
	respProduct := &product.Product{
		Id:          uint32(productModel.ID),
		Name:        productModel.Name,
		Description: productModel.Description,
		Picture:     productModel.Picture,
		Price:       productModel.Price,
		Categories:  mapCategoriesToString(productModel.Categories),
	}
	return &product.GetProductResp{Product: respProduct}, nil
}
