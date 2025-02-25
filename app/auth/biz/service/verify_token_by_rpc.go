package service

import (
	"context"
	"github.com/doutokk/doutok/app/auth/biz/utils"
	"github.com/doutokk/doutok/app/auth/infra/casbin"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/auth"
)

type VerifyTokenByRPCService struct {
	ctx context.Context
} // NewVerifyTokenByRPCService new VerifyTokenByRPCService
func NewVerifyTokenByRPCService(ctx context.Context) *VerifyTokenByRPCService {
	return &VerifyTokenByRPCService{ctx: ctx}
}

// Run create note info
func (s *VerifyTokenByRPCService) Run(req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	token := req.Token

	result, err := utils.ValidateJWT(token)

	var role string
	// 说明token不合法或者没有token
	if err != nil {
		role = "tourist"
	} else {
		role = "test"
	}

	if !casbin.CheckAuthByRBAC(role, req.Uri, req.Method) {
		return &auth.VerifyResp{
			Res:    false,
			UserId: 0,
		}, nil
	}

	// 游客
	if result == nil {
		return &auth.VerifyResp{
			Res:    true,
			UserId: 0,
		}, nil
	}

	userId := result.UserID
	return &auth.VerifyResp{
		Res:    true,
		UserId: int32(userId),
	}, nil
}
