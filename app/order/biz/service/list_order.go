package service

import (
	"context"
	"github.com/PengJingzhao/douyin-commerce/app/order/biz/dal/query"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/cart"

	order "github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/order"
)

type ListOrderService struct {
	ctx context.Context
}

// NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	o := query.Q.Order
	oi := query.Q.OrderItem
	orders, err := query.Q.Order.Where(o.UserID.Eq(req.UserId)).Find()
	if err != nil {
		return nil, err
	}
	resp = new(order.ListOrderResp)
	for _, m := range orders {
		orderItems, err := query.Q.OrderItem.Where(oi.OrderID.Eq(m.OrderID)).Find()
		if err != nil {
			return nil, err
		}
		orderItemsResp := make([]*order.OrderItem, 0)
		for _, orderItem := range orderItems {
			orderItemsResp = append(orderItemsResp, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: orderItem.ProductID,
					Quantity:  orderItem.Quantity,
				},
				Cost: float32(orderItem.Cost),
			})
		}
		resp.Orders = append(resp.Orders, &order.Order{
			OrderId:      m.OrderID,
			UserId:       m.UserID,
			UserCurrency: m.UserCurrency,
			Email:        m.Email,
			Address: &order.Address{
				StreetAddress: m.StreetAddress,
				City:          m.City,
				State:         m.State,
				Country:       m.Country,
				ZipCode:       m.ZipCode,
			},
			OrderItems: orderItemsResp,
		})
	}
	return
}
