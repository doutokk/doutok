package service

import (
	"context"
	"github.com/PengJingzhao/douyin-commerce/app/product/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/app/product/biz/model"
	product "github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// 查询商品列表
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	var products []*model.Product
	query := mysql.DB

	// Filter by category_name if provided
	query = query.Joins("JOIN product_categories ON product_categories.id = products.id").
		Where("product_categories.name = ?", req.CategoryName)

	// Pagination
	offset := (req.PageSize - 1) * req.PageSize
	if err := query.Offset(int(offset)).Limit(int(req.PageSize)).Preload("Categories").Find(&products).Error; err != nil {
		return nil, err
	}

	// map
	respProducts := make([]*product.Product, len(products))
	for i, p := range products {
		respProducts[i] = &product.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			Categories:  mapCategoriesToString(p.Categories),
		}
	}
	return &product.ListProductsResp{Products: respProducts}, nil
}

// Helper function to map categories to a list of strings
func mapCategoriesToString(categories []model.ProductCategory) []string {
	strCategories := make([]string, len(categories))
	for i, c := range categories {
		strCategories[i] = c.Name
	}
	return strCategories
}
