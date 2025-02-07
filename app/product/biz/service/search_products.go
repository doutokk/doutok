package service

import (
	"context"
	"github.com/PengJingzhao/douyin-commerce/app/product/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/app/product/biz/model"
	product "github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	// Finish your business logic.
	var products []model.Product

	// Search products by name or description
	if err := mysql.DB.Where("name LIKE ?", "%"+req.Query+"%").
		Or("description LIKE ?", "%"+req.Query+"%").
		Preload("Categories").Find(&products).Error; err != nil {
		return nil, err
	}

	// Map models.Product to product.Product
	respResults := make([]*product.Product, len(products))
	for i, p := range products {
		respResults[i] = &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			Categories:  mapCategoriesToString(p.Categories),
		}
	}

	return &product.SearchProductsResp{Results: respResults}, nil
}
