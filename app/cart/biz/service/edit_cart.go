package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/cart/biz/dal/model"
	"github.com/doutokk/doutok/app/cart/biz/dal/query"
	"github.com/doutokk/doutok/app/cart/infra/rpc"
	"github.com/doutokk/doutok/common/utils"
	cart "github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
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
	userId := utils.GetUserId(&s.ctx)
	klog.Warn("userId", userId)
	for _, ids := range req.Items {

		_, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: ids.ProductId})
		if err != nil {
			return nil, fmt.Errorf("product %v not found", ids.ProductId)
		}

		item, innerErr := query.Q.CartItem.GetByUserIdAndProductId(userId, ids.ProductId)

		nowQuantity := ids.Quantity
		if innerErr != nil {
			innerErr = nil

			// 查数量是否小于 0 了
			if nowQuantity < 0 {
				return nil, fmt.Errorf("quantity must be greater than 0")
			}

			innerErr = query.Q.CartItem.Create(&model.CartItem{
				UserId:    userId,
				ProductId: ids.ProductId,
				Quantity:  uint32(nowQuantity),
			})
			if innerErr != nil {
				return nil, innerErr
			}

		} else {

			if ids.Quantity == 0 {
				// 删除
				_, innerErr := ci.Where(ci.UserId.Eq(userId), ci.ProductId.Eq(ids.ProductId)).Delete()
				if innerErr != nil {
					return nil, innerErr
				}
			} else {

				_, innerErr = ci.Where(ci.ID.Eq(item.ID)).Update(ci.Quantity, nowQuantity)
			}
		}
	}

	return
}
