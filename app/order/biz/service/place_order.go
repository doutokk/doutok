package service

import (
	"context"
	"github.com/doutokk/doutok/app/order/biz/dal/model"
	"github.com/doutokk/doutok/app/order/biz/dal/query"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
	"github.com/hashicorp/go-uuid"
)

type PlaceOrderService struct {
	ctx context.Context
}

// NewPlaceOrderService new PlaceOrderService
func NewPlaceOrderService(ctx context.Context) *PlaceOrderService {
	return &PlaceOrderService{ctx: ctx}
}

// Run create note info
func (s *PlaceOrderService) Run(req *order.PlaceOrderReq) (resp *order.PlaceOrderResp, err error) {
	// Finish your business logic.
	// TODO: 事务支持
	// 插入订单
	generateUUID, err := uuid.GenerateUUID()

	ord := &model.Order{
		OrderID:       generateUUID,
		UserID:        req.UserId,
		UserCurrency:  req.UserCurrency,
		Email:         req.Email,
		StreetAddress: req.Address.StreetAddress,
		City:          req.Address.City,
		State:         req.Address.State,
		Country:       req.Address.Country,
		ZipCode:       req.Address.ZipCode,
		PaidStatus:    false,
	}
	err = query.Q.Order.Create(ord)
	if err != nil {
		return nil, err
	}

	// 插入物品列表
	items := make([]*model.OrderItem, 0)
	for _, item := range req.OrderItems {
		// 插入订单物品
		orderItem := &model.OrderItem{
			OrderID:   generateUUID,
			ProductID: item.Item.ProductId,
			Quantity:  item.Item.Quantity,
			Cost:      float64(item.Cost),
		}
		items = append(items, orderItem)
	}
	err = query.Q.OrderItem.Create(items...)
	if err != nil {
		return nil, err
	}
	resp = &order.PlaceOrderResp{
		Order: &order.OrderResult{
			OrderId: generateUUID,
		},
	}
	return
}
