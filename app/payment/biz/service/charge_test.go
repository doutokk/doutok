package service

import (
	"context"
	payment "github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/payment"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"testing"
)

func TestCharge_Run(t *testing.T) {

	ctx := context.Background()
	// init req and assert value

	// 连consul
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	cli, err := paymentservice.NewClient("payment", client.WithResolver(r))

	req := &payment.ChargeReq{
		UserId:  1,
		Amount:  12.32,
		OrderId: "ttest",
	}
	resp, err := cli.Charge(ctx, req)

	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

}
