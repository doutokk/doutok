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

type CancelService struct {
	ctx context.Context
}

// NewCancelService new CancelService
func NewCancelService(ctx context.Context) *CancelService {
	return &CancelService{ctx: ctx}
}

// Run create note info
func (s *CancelService) Run(req *payment.CancelPaymentReq) (resp *payment.CancelPaymentResp, err error) {
	// Finish your business logic.
	fmt.Printf("CancelService is called with req: %+v\n", req)
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
	if err != nil {
		klog.Warn("订单状态获取失败")
		return nil, errors.New("订单状态获取失败")
	}
	if err == nil && orderFSM.GetStatus() != fsm.PAYING {
		klog.Warn("订单状态不正确")
		return nil, errors.New("订单状态不正确")
	}
	err = orderFSM.PaymentFailed(s.ctx)
	if err != nil {
		klog.Warn("订单状态更新失败")
		return nil, errors.New("订单状态更新失败")
	}

	return &payment.CancelPaymentResp{
		Success: true,
	}, nil
}
