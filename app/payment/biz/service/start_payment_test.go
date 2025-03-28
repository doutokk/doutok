package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"github.com/doutokk/doutok/app/payment/biz/dal"
	"github.com/doutokk/doutok/app/payment/biz/dal/mysql"
	"github.com/doutokk/doutok/app/payment/biz/dal/query"
	"github.com/doutokk/doutok/app/payment/infra/rpc"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
	"testing"
)

func TestStartPayment_Run(t *testing.T) {
	ctx := context.Background()
	// init req and assert value

	dal.Init()
	rpc.InitClient()
	query.SetDefault(mysql.DB)

	req := &payment.StartPaymentReq{
		OrderId: "fa3e17c8-ff69-c02e-8c0b-855e3ede9255",
	}
	ctx = metadata.AppendToOutgoingContext(ctx, "user-id", "7")
	s := NewStartPaymentService(ctx)
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
