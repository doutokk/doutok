package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/checkout/infra/rpc"
	checkout "github.com/doutokk/doutok/app/checkout/kitex_gen/checkout"
	"github.com/doutokk/doutok/app/payment/kitex_gen/payment"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
	"strconv"
	"sync"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.

	// 1. 搞到购物车里的商品，计算价格
	getCartResp, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.GetUserId(),
	})
	if err != nil {
		klog.Error()
		return
	}
	if getCartResp.GetCart() == nil {
		err = fmt.Errorf("getCartResp.Cart is nil")
		klog.Error(err)
		return
	}

	// 计算价格
	var (
		oi    []*order.OrderItem
		price float32 = 0
	)
	cartItems := getCartResp.GetCart().Items
	for _, cartItem := range cartItems {
		productId := cartItem.ProductId
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: productId,
		})

		if resultErr != nil {
			err = resultErr
			klog.Error(err)
			return
		}
		if productResp.GetProduct() == nil {
			continue
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
		klog.Error("Charge.err:%v", err)
		return
	}
	klog.Info("Charge.resp:%s", chargeresp)

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	// 下面貌似可以后台异步，但是要考虑失败怎么补偿
	// 清空购物车
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, resultErr := rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{
			UserId: req.GetUserId(),
		})
		if resultErr != nil {
			errChan <- fmt.Errorf("EmptyCart.err:%v", resultErr)
		}
	}()

	// 标记已支付
	wg.Add(1)
	go func() {
		defer wg.Done()
		_, resultErr := rpc.OrderClient.MarkOrderPaid(s.ctx, &order.MarkOrderPaidReq{
			OrderId: orderResp.Order.OrderId,
			UserId:  req.GetUserId(),
		})
		if resultErr != nil {
			errChan <- fmt.Errorf("MarkOrderPaid.err:%v", resultErr)
		}
	}()

	wg.Wait()
	close(errChan)

	for resultErr := range errChan {
		if resultErr != nil {
			err = resultErr
			klog.Error(err)
			return nil, err
		}
	}

	return &checkout.CheckoutResp{
		OrderId: orderResp.Order.OrderId,
		// 交易生成的 id
		TransactionId: chargeresp.TransactionId,
	}, nil
}
