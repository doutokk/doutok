package product

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
)

func ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (resp *product.ListProductsResp, err error) {
	resp, err = defaultClient.ListProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *product.SearchProductsReq, callOptions ...callopt.Option) (resp *product.SearchProductsResp, err error) {
	resp, err = defaultClient.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func EditProduct(ctx context.Context, req *product.EditProductReq, callOptions ...callopt.Option) (resp *product.EditProductResp, err error) {
	resp, err = defaultClient.EditProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "EditProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProductBatch(ctx context.Context, req *product.GetProductBatchReq, callOptions ...callopt.Option) (resp *product.GetProductBatchResp, err error) {
	resp, err = defaultClient.GetProductBatch(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProductBatch call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateProduct(ctx context.Context, req *product.CreateProductReq, callOptions ...callopt.Option) (resp *product.CreateProductResp, err error) {
	resp, err = defaultClient.CreateProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
