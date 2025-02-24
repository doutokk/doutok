package service

import (
	"context"
	"github.com/doutokk/doutok/app/order/biz/dal/query"
	"github.com/doutokk/doutok/common/utils"
	order "github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
)

type GetOrderService struct {
	ctx context.Context
} // NewGetOrderService new GetOrderService
func NewGetOrderService(ctx context.Context) *GetOrderService {
	return &GetOrderService{ctx: ctx}
}

// Run create note info
func (s *GetOrderService) Run(req *order.GetOrderReq) (resp *order.GetOrderResp, err error) {
	userId := utils.GetUserId(s.ctx)
	resp = new(order.GetOrderResp)
	o := query.Q.Order
	oi := query.Q.OrderItem
	oneOrder, err := o.Where(o.OrderID.Eq(req.Id)).Where(o.UserID.Eq(uint32(userId))).First()

	orderItems, err := oi.Where(oi.OrderID.Eq(oneOrder.OrderID)).Find()
	orderItemsResp := make([]*order.OrderItem, 0)
	for _, orderItem := range orderItems {

		if orderItem == nil {
			continue
		}

		orderItemsResp = append(orderItemsResp, &order.OrderItem{
			Item: &order.CartItem{
				ProductId: orderItem.ProductID,
				Quantity:  orderItem.Quantity,
			},
			Cost: float32(orderItem.Cost),
		})
	}
	resp.Order = &order.Order{
		OrderId:      oneOrder.OrderID,
		UserId:       oneOrder.UserID,
		UserCurrency: oneOrder.UserCurrency,
		Email:        oneOrder.Email,
		Address: &order.Address{
			StreetAddress: oneOrder.StreetAddress,
			City:          oneOrder.City,
			State:         oneOrder.State,
			Country:       oneOrder.Country,
			ZipCode:       oneOrder.ZipCode,
		},
		OrderItems: orderItemsResp,
	}
	return
}
