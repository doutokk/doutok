package service

import (
	"context"
	"errors"
	"github.com/doutokk/doutok/app/order/biz/dal/query"
	"github.com/doutokk/doutok/app/order/infra/rpc"
	"github.com/doutokk/doutok/common/utils"
	order "github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
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
	if err != nil {
		return nil, errors.New("order not found")
	}

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

	// 批量获取商品详细信息
	idsS := make(map[uint32]interface{}, 0)
	for _, item := range resp.Order.OrderItems {
		idsS[item.Item.ProductId] = nil
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
	for _, item := range resp.Order.OrderItems {
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

	return
}
