package service

import (
	"context"
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal/query"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	// TODO: add transaction
	ci := query.Q.CartItem
	items, err := ci.Where(ci.UserId.Eq(req.UserId)).Find()
	if err != nil {
		return nil, err
	}
	cartItems := make([]*cart.CartItem, 0, len(items))
	for _, item := range items {
		cartItems = append(cartItems, &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  int32(item.Quantity),
		})
	}
	resp = &cart.GetCartResp{
		Cart: &cart.Cart{
			UserId: req.UserId,
			Items:  cartItems,
		},
	}
	return
}
