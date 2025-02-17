package service

import (
	"context"
	"fmt"
	"github.com/doutokk/doutok/app/cart/biz/dal/model"
	"github.com/doutokk/doutok/app/cart/biz/dal/query"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	// TODO: add transaction
	ci := query.Q.CartItem
	item, err := query.Q.CartItem.GetByUserIdAndProductId(req.UserId, req.Item.ProductId)
	if err != nil && err.Error() == "record not found" {
		err = nil
		if req.Item.Quantity <= 0 {
			return nil, fmt.Errorf("quantity must be greater than 0")
		}
		err = query.Q.CartItem.Create(&model.CartItem{
			UserId:    req.UserId,
			ProductId: req.Item.ProductId,
			Quantity:  uint32(req.Item.Quantity),
		})
	}
	if err != nil {
		return nil, err
	}
	if int32(item.Quantity)+req.Item.Quantity <= 0 {
		return nil, fmt.Errorf("quantity must be greater than 0")
	}
	_, err = ci.Where(ci.ID.Eq(item.ID)).Update(ci.Quantity, int32(item.Quantity)+req.Item.Quantity)

	return
}
