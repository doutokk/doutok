package service

import (
	"context"
	"github.com/doutokk/doutok/app/order/biz/dal/query"
	"github.com/doutokk/doutok/app/order/infra/rpc"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
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
	//userId := utils.GetUserId(&s.ctx)
	userId := 7
	// Finish your business logic.
	o := query.Q.Order
	oi := query.Q.OrderItem
	orders, err := query.Q.Order.Where(o.UserID.Eq(uint32(userId))).Find()
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
				Item: &order.CartItem{
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

	// 批量获取商品详细信息
	idsS := make(map[uint32]interface{}, 0)
	for _, orderi := range resp.Orders {
		for _, item := range orderi.OrderItems {
			idsS[item.Item.ProductId] = nil
		}
	}
	idsL := make([]uint32, 0)
	for id := range idsS {
		idsL = append(idsL, id)
	}
	// 获取商品详细信息
	products, err := rpc.ProductClient.GetProductBatch(s.ctx, &product.GetProductBatchReq{Ids: idsL})
	if err != nil {
		return nil, err
	}
	// 使用获取到的商品详细信息构造 map
	pm := make(map[uint32]*product.Product, len(products.Item))
	for _, p := range products.Item {
		pm[p.Id] = p
	}

	// 填充商品详细信息
	for _, orderi := range resp.Orders {
		for _, item := range orderi.OrderItems {
			p, ok := pm[item.Item.ProductId]
			if !ok {
				continue
			}
			item.Item.Description = p.Description
			item.Item.ProductName = p.Name
			item.Item.Img = p.Picture
			item.Item.Price = p.Price
			item.Cost = float32(item.Item.Quantity) * p.Price
		}
	}

	return
}
