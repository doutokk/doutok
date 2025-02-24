package service

import (
	"context"
	"fmt"
	"github.com/doutokk/doutok/app/cart/biz/dal/model"
	"github.com/doutokk/doutok/app/cart/biz/dal/query"
	"github.com/doutokk/doutok/common/utils"
	cart "github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"
)

type EditCartService struct {
	ctx context.Context
} // NewEditCartService new EditCartService
func NewEditCartService(ctx context.Context) *EditCartService {
	return &EditCartService{ctx: ctx}
}

// Run create note info
func (s *EditCartService) Run(req *cart.EditCartReq) (resp *cart.EditCartResp, err error) {
	// Finish your business logic.
	// TODO: add transaction
	ci := query.Q.CartItem
	userId := uint32(utils.GetUserId(s.ctx))
	for _, ids := range req.Items {
		item, innerErr := query.Q.CartItem.GetByUserIdAndProductId(userId, ids.ProductId)

		// todo:检查商品是否存在

		nowQuantity := int32(item.Quantity) + ids.Quantity
		if innerErr != nil {
			innerErr = nil

			// 查数量是否小于0了
			if nowQuantity < 0 {
				return nil, fmt.Errorf("quantity must be greater than 0")
			}

			innerErr = query.Q.CartItem.Create(&model.CartItem{
				UserId:    userId,
				ProductId: ids.ProductId,
				Quantity:  uint32(nowQuantity),
			})
		}
		if innerErr != nil {
			return nil, innerErr
		}

		_, innerErr = ci.Where(ci.ID.Eq(item.ID)).Update(ci.Quantity, nowQuantity)
	}

	return
}
