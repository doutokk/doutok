package service

import (
	"context"
	"github.com/doutokk/doutok/app/auth/biz/utils"
	"github.com/doutokk/doutok/app/auth/infra/casbin"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/auth"
	"strconv"
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
	resp = &auth.VerifyResp{}
	var userId int

	if result != nil && err == nil {
		userId = result.UserID
	} else {
		userId = 0
	}

	var role string
	// 说明token不合法或者没有token
	if userId == 0 {
		role = "base"
	} else {
		role = strconv.Itoa(userId)
	}

	resp.Res = casbin.CheckAuthByRBAC(role, req.Uri, req.Method)
	resp.UserId = int32(userId)
	return resp, nil
}
