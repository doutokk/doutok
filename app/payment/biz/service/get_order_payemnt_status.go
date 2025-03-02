package service

import (
	"context"
	"fmt"
	"github.com/doutokk/doutok/app/payment/biz/fsm"
	payment "github.com/doutokk/doutok/rpc_gen/kitex_gen/payment"
	"strings"
)

type GetOrderPayemntStatusService struct {
	ctx context.Context
}

// NewGetOrderPayemntStatusService new GetOrderPayemntStatusService
func NewGetOrderPayemntStatusService(ctx context.Context) *GetOrderPayemntStatusService {
	return &GetOrderPayemntStatusService{ctx: ctx}
}

// Run create note info
func (s *GetOrderPayemntStatusService) Run(req *payment.GetOrderPayemntStatusReq) (resp *payment.GetOrderPayemntStatusResp, err error) {
	// Finish your business logic.

	fmt.Printf("GetOrderPayemntStatusService is called with req: %+v\n", req)

	orderFSM, err := fsm.RestoreFromDB(req.OrderId)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			return &payment.GetOrderPayemntStatusResp{Status: "Uncreated"}, nil
		}
		return nil, err
	}
	resp = &payment.GetOrderPayemntStatusResp{
		Status: string(orderFSM.GetStatus()),
	}

	return
}
