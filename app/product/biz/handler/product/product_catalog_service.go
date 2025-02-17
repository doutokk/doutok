// Code generated by hertz generator.

package product

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/product/biz/service"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type bffProduct struct {
	ProductId   int     `json:"product_id"`   // JSON tag: product_id
	ProductName string  `json:"product_name"` // JSON tag: product_name
	Price       float32 `json:"price"`        // JSON tag: price
	Description string  `json:"description"`  // JSON tag: description
	Img         string  `json:"img"`          // JSON tag: img
	Quantity    int     `json:"quantity"`     // JSON tag: quantity
}

func convertProductToBff(product *product.Product) *bffProduct {
	return &bffProduct{
		ProductId:   int(product.Id),
		ProductName: product.Name,
		Price:       product.Price,
		Description: product.Description,
		Img:         product.Picture,
		Quantity:    0,
	}
}

// ListProducts .
// @router /product [POST]
func ListProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.ListProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	svc := service.NewListProductsService(ctx)

	resp, err := svc.Run(&req)

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	var bffProducts []*bffProduct
	for _, product := range resp.Item {
		bffProducts = append(bffProducts, convertProductToBff(product))
		klog.Info(bffProducts[len(bffProducts)-1])
	}

	c.JSON(consts.StatusOK, bffProducts)
}

// GetProduct .
// @router /product/{id} [GET]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.GetProductReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	svc := service.NewGetProductService(ctx)
	resp, err := svc.Run(&req)

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	c.JSON(consts.StatusOK, convertProductToBff(resp.Product))
}

// SearchProducts .
// @router product [GET]
func SearchProducts(ctx context.Context, c *app.RequestContext) {
	var err error
	var req product.SearchProductsReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	svc := service.NewSearchProductsService(ctx)
	resp, err := svc.Run(&req)

	// 把resp里的products转为bffProduct返回
	var bffProducts []*bffProduct
	for _, product := range resp.Item {
		bffProducts = append(bffProducts, convertProductToBff(product))
	}

	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	c.JSON(consts.StatusOK, bffProducts)
}
