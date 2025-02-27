package service

import (
	"context"
	"testing"

	"github.com/doutokk/doutok/app/order/biz/dal"
	"github.com/doutokk/doutok/app/order/biz/dal/mysql"
	"github.com/doutokk/doutok/app/order/biz/dal/query"
	"github.com/doutokk/doutok/app/order/infra/rpc"
	"github.com/doutokk/doutok/common/mocks"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/order"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
	"go.uber.org/mock/gomock"
)

func TestListOrder_Run(t *testing.T) {
	ctx := context.Background()
	dal.Init()
	// use go run cmd/gorm_gen/main.go to generate the code
	query.SetDefault(mysql.DB)
	s := NewListOrderService(ctx)
	// init req and assert value

	ctrl := gomock.NewController(t)

	productClient := mocks.NewMockproductClient(ctrl)
	rpc.ProductClient = productClient

	productClient.EXPECT().GetProductBatch(ctx, gomock.Any()).Return(&product.GetProductBatchResp{
		Item: []*product.Product{
			{
				Id:          21,
				Name:        "Babel 贴纸",
				Description: "babel",
				Picture:     "https://doutok.oss-cn-shenzhen.aliyuncs.com/babel.webp",
				Price:       1.9,
				Categories:  []string{},
			},
			{
				Id:          22,
				Name:        "Vscode 贴纸",
				Description: "vscode",
				Picture:     "https://doutok.oss-cn-shenzhen.aliyuncs.com/vscode.webp",
				Price:       1.9,
				Categories:  []string{},
			},
		},
	}, nil)

	req := &order.ListOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
