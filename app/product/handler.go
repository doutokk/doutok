package main

import (
	"context"
	"github.com/doutokk/doutok/app/product/biz/service"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// GetProductBatch implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProductBatch(ctx context.Context, req *product.GetProductBatchReq) (resp *product.GetProductBatchResp, err error) {
	resp, err = service.NewGetProductBatchService(ctx).Run(req)

	return resp, err
}

// EditProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) EditProduct(ctx context.Context, req *product.EditProductReq) (resp *product.EditProductResp, err error) {
	resp, err = service.NewEditProductService(ctx).Run(req)

	return resp, err
}
