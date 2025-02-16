package service

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/doutokk/doutok/app/cart/kitex_gen/cart"
	"github.com/doutokk/doutok/app/checkout/infra/rpc"
	checkout "github.com/doutokk/doutok/app/checkout/kitex_gen/checkout"
	"github.com/doutokk/doutok/app/order/kitex_gen/order"
	"github.com/doutokk/doutok/app/payment/kitex_gen/payment"
	"github.com/doutokk/doutok/app/product/kitex_gen/product"
	"github.com/doutokk/doutok/common/mocks"
	"github.com/golang/mock/gomock"
	"strconv"
	"testing"
)

func TestCheckout_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()
	// 初始化 mock 对象
	cartClient := mocks.NewMockcartClient(ctrl)
	productClient := mocks.NewMockproductClient(ctrl)
	orderClient := mocks.NewMockorderClient(ctrl)
	paymentClient := mocks.NewMockpaymentClient(ctrl)
	rpc.CartClient = cartClient
	rpc.ProductClient = productClient
	rpc.OrderClient = orderClient
	rpc.PaymentClient = paymentClient

	// 设置 mock 返回逻辑
	cartItems := []*cart.CartItem{
		{
			ProductId: 1,
			Quantity:  2,
		},
	}
	yourCart := &cart.Cart{
		UserId: 123,
		Items:  cartItems,
	}
	cartClient.EXPECT().
		GetCart(ctx, gomock.Any()).
		Return(&cart.GetCartResp{
			Cart: yourCart, // 返回的购物车内容
		}, nil)

	productClient.EXPECT().
		GetProduct(ctx, gomock.Any()).
		Return(&product.GetProductResp{
			Product: &product.Product{
				Id:    1,
				Price: 10,
			},
		}, nil)

	// 创建订单结果
	odResult := &order.OrderResult{}
	var orderId int = gofakeit.Number(100000, 999999)
	odResult.OrderId = strconv.Itoa(orderId) // 生成随机订单 ID
	orderClient.EXPECT().
		PlaceOrder(ctx, gomock.Any()).
		Return(&order.PlaceOrderResp{
			Order: odResult, // 返回的订单结果
		}, nil)

	// 处理付款
	paymentClient.EXPECT().
		Charge(ctx, gomock.Any()).
		Return(&payment.ChargeResp{
			TransactionId: "123", // 返回的交易 ID
		}, nil)

	// 标记订单已付款
	orderClient.EXPECT().
		MarkOrderPaid(ctx, gomock.Any())

	// 清空购物车
	cartClient.EXPECT().
		EmptyCart(ctx, gomock.Any())

	req := &checkout.CheckoutReq{
		UserId: 123,
		Email:  "test@example.com",
		Address: &checkout.Address{
			StreetAddress: "123 Street",
			City:          "City",
			State:         "State",
			Country:       "Country",
			ZipCode:       "12345",
		},
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          "4111111111111111",
			CreditCardExpirationMonth: 12,
			CreditCardExpirationYear:  2025,
			CreditCardCvv:             123,
		},
	}

	s := NewCheckoutService(ctx)
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	if resp.OrderId == strconv.Itoa(orderId) {
		t.Log("测试通过")
	}

	// todo: edit your unit test

}
