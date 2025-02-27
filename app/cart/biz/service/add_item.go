package service

import (
	"context"
	"fmt"
	"github.com/doutokk/doutok/app/cart/biz/dal/model"
	"github.com/doutokk/doutok/app/cart/biz/dal/query"
	"github.com/doutokk/doutok/common/utils"
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
	// 校验商品数量是否大于0
	if req.Item.Quantity <= 0 {
		return nil, fmt.Errorf("quantity must be greater than 0")
	}
	userId := uint32(utils.GetUserId(s.ctx))

	// 获取用户的购物车商品
	ci := query.Q.CartItem
	item, err := ci.GetByUserIdAndProductId(userId, req.Item.ProductId)

	// 如果购物车没有该商品，创建新记录
	if err != nil && err.Error() == "record not found" {
		err = nil // 重置错误
		err = ci.Create(&model.CartItem{
			UserId:    userId,
			ProductId: req.Item.ProductId,
			Quantity:  uint32(req.Item.Quantity),
		})
		return nil, err
	}

	// 如果商品已经存在，更新数量
	if err == nil {
		newQuantity := int32(item.Quantity) + req.Item.Quantity
		if newQuantity <= 0 {
			return nil, fmt.Errorf("quantity must be greater than 0")
		}
		_, err = ci.Where(ci.ID.Eq(item.ID)).Update(ci.Quantity, newQuantity)
	}

	return nil, err
}
