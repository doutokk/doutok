package service

import (
	"context"
	"github.com/doutokk/doutok/app/auth/biz/utils"
	auth "github.com/doutokk/doutok/rpc_gen/kitex_gen/auth"
	"time"
)

type DeliverTokenByRPCService struct {
	ctx context.Context
} // NewDeliverTokenByRPCService new DeliverTokenByRPCService
func NewDeliverTokenByRPCService(ctx context.Context) *DeliverTokenByRPCService {
	return &DeliverTokenByRPCService{ctx: ctx}
}

// 分发 token
func (s *DeliverTokenByRPCService) Run(req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// Finish your business logic.
	userId := req.UserId

	jwt, err := utils.GenerateJWT(int(userId), 24*time.Hour)
	if err != nil {
		return &auth.DeliveryResp{
			Token: "",
		}, err
	}

	return &auth.DeliveryResp{
		Token: jwt,
	}, nil
}
