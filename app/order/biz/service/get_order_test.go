package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/doutokk/doutok/app/payment/biz/dal"
	"github.com/doutokk/doutok/app/payment/infra/rpc"
	order "github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
	"testing"
)

func TestGetOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetOrderService(ctx)
	// init req and assert value

	dal.Init()
	rpc.InitClient()

	req := &order.GetOrderReq{
		Id: "fa3e17c8-ff69-c02e-8c0b-855e3ede9255",
	}
	metadata.AppendToOutgoingContext(ctx, "user-id", "7")
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
