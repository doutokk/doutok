package service

import (
	"context"
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal/query"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/cart"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	ci := query.Q.CartItem

	_, err = ci.Where(ci.UserId.Eq(req.UserId)).Delete()
	if err != nil {
		return nil, err
	}

	return
}
