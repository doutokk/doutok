package main

import (
	"context"
	"github.com/doutokk/doutok/app/cart/biz/service"
	cart "github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	resp, err = service.NewAddItemService(ctx).Run(req)

	return resp, err
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	resp, err = service.NewGetCartService(ctx).Run(req)

	return resp, err
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	resp, err = service.NewEmptyCartService(ctx).Run(req)

	return resp, err
}

// EditCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EditCart(ctx context.Context, req *cart.EditCartReq) (resp *cart.EditCartResp, err error) {
	resp, err = service.NewEditCartService(ctx).Run(req)

	return resp, err
}

// FrontendGetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) FrontendGetCart(ctx context.Context, req *cart.FrontendGetCartReq) (resp *cart.FrontendGetCartResp, err error) {
	resp, err = service.NewFrontendGetCartService(ctx).Run(req)

	return resp, err
}
