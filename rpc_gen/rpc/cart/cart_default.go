package cart

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	cart "github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"
)

func AddItem(ctx context.Context, req *cart.AddItemReq, callOptions ...callopt.Option) (resp *cart.AddItemResp, err error) {
	resp, err = defaultClient.AddItem(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddItem call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCart(ctx context.Context, req *cart.GetCartReq, callOptions ...callopt.Option) (resp *cart.GetCartResp, err error) {
	resp, err = defaultClient.GetCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func EmptyCart(ctx context.Context, req *cart.EmptyCartReq, callOptions ...callopt.Option) (resp *cart.EmptyCartResp, err error) {
	resp, err = defaultClient.EmptyCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "EmptyCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func EditCart(ctx context.Context, req *cart.EditCartReq, callOptions ...callopt.Option) (resp *cart.EditCartResp, err error) {
	resp, err = defaultClient.EditCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "EditCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FrontendGetCart(ctx context.Context, req *cart.FrontendGetCartReq, callOptions ...callopt.Option) (resp *cart.FrontendGetCartResp, err error) {
	resp, err = defaultClient.FrontendGetCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FrontendGetCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
