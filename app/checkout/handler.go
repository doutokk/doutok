package main

import (
	"context"
	"github.com/PengJingzhao/douyin-commerce/app/checkout/biz/service"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/checkout"
)

// CheckoutServiceImpl implements the last service interface defined in the IDL.
type CheckoutServiceImpl struct{}

// Checkout implements the CheckoutServiceImpl interface.
func (s *CheckoutServiceImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	resp, err = service.NewCheckoutService(ctx).Run(req)

	return resp, err
}
