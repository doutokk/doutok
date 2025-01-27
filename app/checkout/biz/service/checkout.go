package service

import (
	"context"
	"fmt"
	"github.com/PengJingzhao/douyin-commerce/app/checkout/infra/rpc"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/cart"
	checkout "github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/checkout"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/order"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/payment"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
// todo: 事务？
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.

	// 1. 搞到购物车里的商品，计算价格
	getCartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.GetUserId(),
	})
	if err != nil {
		klog.Error()
	}

	cartItems := getCartResp.GetCart().Items
	var (
		oi    []*order.OrderItem
		price float32 = 0
	)
	for _, cartItem := range cartItems {
		productId := cartItem.ProductId
		productResp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: productId,
		})

		if err != nil {
			klog.Error()
			return
		}
		cost := productResp.GetProduct().Price * float32(cartItem.Quantity)
		price += cost
		oi = append(oi, &order.OrderItem{
			Item: &cart.CartItem{ProductId: cartItem.ProductId, Quantity: cartItem.Quantity},
			Cost: cost,
		})
	}

	// 2. 生成订单
	orderReq := &order.PlaceOrderReq{
		UserId:       req.GetUserId(),
		UserCurrency: "CNY",
		OrderItems:   oi,
		Email:        req.Email,
	}
	if req.Address != nil {
		addr := req.Address
		zipCodeInt, _ := strconv.Atoi(addr.ZipCode)
		orderReq.Address = &order.Address{
			StreetAddress: addr.StreetAddress,
			City:          addr.City,
			Country:       addr.Country,
			State:         addr.State,
			ZipCode:       int32(zipCodeInt),
		}
	}
	orderResp, err := rpc.OrderClient.PlaceOrder(s.ctx, orderReq)
	// todo: 这样能把错误日志收集到链路吗
	if err != nil {
		err = fmt.Errorf("PlaceOrder.err:%v", err)
		klog.Error(err)
		return
	}
	klog.Info("PlaceOrder.resp:", orderResp)

	// 3. 清空购物车
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{
		UserId: req.GetUserId(),
	})
	if err != nil {
		err = fmt.Errorf("EmptyCart.err:%v", err)
		klog.Error()
		return
	}

	// 4. 支付
	chargeReq := &payment.ChargeReq{
		Amount:  price,
		UserId:  req.GetUserId(),
		OrderId: orderResp.Order.OrderId,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			// 安全码？
			CreditCardCvv: req.CreditCard.CreditCardCvv,
		},
	}
	chargeresp, err := rpc.PaymentClient.Charge(s.ctx, chargeReq)
	if err != nil {
		err = fmt.Errorf("Charge.err:%v", err)
		klog.Error()
		return
	}
	klog.Info("Charge.resp:", chargeresp)

	// 5. 标注订单已支付
	_, err = rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{
		OrderId: orderResp.Order.OrderId,
		UserId:  req.GetUserId(),
	})

	if err != nil {
		err = fmt.Errorf("MarkOrderPaid.err:%v", err)
		klog.Error()
		return
	}

	return &checkout.CheckoutResp{
		OrderId: orderResp.Order.OrderId,
		// 交易生成的id
		TransactionId: chargeresp.TransactionId,
	}, nil
}
