package product

import (
	"context"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product/productcatalogservice"
)

type RPCClient interface {
	KitexClient() productcatalogservice.Client
	Service() string
	ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error)
	GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error)
	EditProduct(ctx context.Context, Req *product.EditProductReq, callOptions ...callopt.Option) (r *product.EditProductResp, err error)
	GetProductBatch(ctx context.Context, Req *product.GetProductBatchReq, callOptions ...callopt.Option) (r *product.GetProductBatchResp, err error)
	CreateProduct(ctx context.Context, Req *product.CreateProductReq, callOptions ...callopt.Option) (r *product.CreateProductResp, err error)
	DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productcatalogservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productcatalogservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productcatalogservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (r *product.ListProductsResp, err error) {
	return c.kitexClient.ListProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	return c.kitexClient.GetProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (r *product.SearchProductsResp, err error) {
	return c.kitexClient.SearchProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) EditProduct(ctx context.Context, Req *product.EditProductReq, callOptions ...callopt.Option) (r *product.EditProductResp, err error) {
	return c.kitexClient.EditProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProductBatch(ctx context.Context, Req *product.GetProductBatchReq, callOptions ...callopt.Option) (r *product.GetProductBatchResp, err error) {
	return c.kitexClient.GetProductBatch(ctx, Req, callOptions...)
}

func (c *clientImpl) CreateProduct(ctx context.Context, Req *product.CreateProductReq, callOptions ...callopt.Option) (r *product.CreateProductResp, err error) {
	return c.kitexClient.CreateProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteProduct(ctx context.Context, Req *product.DeleteProductReq, callOptions ...callopt.Option) (r *product.DeleteProductResp, err error) {
	return c.kitexClient.DeleteProduct(ctx, Req, callOptions...)
}
