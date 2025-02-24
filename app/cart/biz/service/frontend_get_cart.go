package service

import (
	"context"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"

	"github.com/doutokk/doutok/app/cart/infra/rpc"
	cart "github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"
)

type FrontendGetCartService struct {
	ctx context.Context
}

// NewFrontendGetCartService new FrontendGetCartService
func NewFrontendGetCartService(ctx context.Context) *FrontendGetCartService {
	return &FrontendGetCartService{ctx: ctx}
}

// Run create note info
func (s *FrontendGetCartService) Run(req *cart.FrontendGetCartReq) (resp *cart.FrontendGetCartResp, err error) {
	// Finish your business logic.
	cs := GetCartService{s.ctx}
	cartResp, err := cs.Run(&cart.GetCartReq{
		UserId: 1, //TODO: get user_id from token
	})
	if err != nil {
		return nil, err
	}
	items := make([]*cart.FrontendItem, len(cartResp.Cart.Items))
	for _, item := range cartResp.Cart.Items {
		productResp, err :=
			rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: item.ProductId})
		if err != nil {
			return nil, err
		}
		items = append(items, &cart.FrontendItem{
			ProductId:   productResp.Product.Id,
			ProductName: productResp.Product.Name,
			Price:       productResp.Product.Price,
			Description: productResp.Product.Description,
			Img:         productResp.Product.Picture,
			Quantity:    item.Quantity,
		})
	}
	resp = &cart.FrontendGetCartResp{
		Items: items,
	}
	return
}
