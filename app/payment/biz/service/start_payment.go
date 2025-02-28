package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/payment/biz/fsm"
	"github.com/doutokk/doutok/app/payment/infra/rpc"
	"github.com/doutokk/doutok/common/utils"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
)

type StartPaymentService struct {
	ctx context.Context
}

// NewStartPaymentService new StartPaymentService
func NewStartPaymentService(ctx context.Context) *StartPaymentService {
	return &StartPaymentService{ctx: ctx}
}

// Run create note info
func (s *StartPaymentService) Run(req *payment.StartPaymentReq) (resp *payment.StartPaymentResp, err error) {
	// Finish your business logic.

	fmt.Printf("StartPaymentService is called with req: %+v\n", req)
	userId := utils.GetUserId(&s.ctx)
	oi := req.OrderId
	r, err := rpc.OrderClient.GetOrder(s.ctx, &order.GetOrderReq{Id: oi})
	if err != nil {
		return
	}
	if r.Order.UserId != uint32(userId) {
		klog.Warn("用户名和订单 id 不匹配")
		return nil, errors.New("")
	}
	var orderFSM *fsm.PayOrderFSM
	orderFSM, err = fsm.RestoreFromDB(oi)
	if err == nil && orderFSM.GetStatus() != fsm.CREATED {
		klog.Warn("订单状态不正确")
		return nil, errors.New("订单状态不正确")
	}

	if err != nil {
		amount := 0.0
		for _, item := range r.Order.OrderItems {
			amount += float64(item.Cost)
		}
		orderFSM, err = fsm.NewOrder(fsm.CreatePayOrderReq{
			UserId:  userId,
			OrderId: req.OrderId,
			Amount:  float32(amount),
		})
	}
	if err != nil {
		return nil, err
	}
	fmt.Printf("orderFSM: %+v\n", orderFSM)

	// 生成支付链接
	url, err := orderFSM.StartPayment(s.ctx)
	if err != nil {
		return nil, err
	}

	resp = &payment.StartPaymentResp{
		PaymentUrl: url,
	}

	return
}
