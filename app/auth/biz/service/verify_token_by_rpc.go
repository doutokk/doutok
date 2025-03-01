package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
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
		klog.Warnf("Token validate success: %v %v", userId, token)
		userId = result.UserID
	} else {
		klog.Warnf("Token validate failed: %v %v", err, token)
		userId = 0
	}

	var role string
	// 说明 token 不合法或者没有 token
	if userId == 0 {
		role = "base"
	} else {
		role = strconv.Itoa(userId)
	}

	resp.Res = casbin.CheckAuthByRBAC(role, req.Uri, req.Method)
	resp.UserId = int32(userId)
	return resp, nil
}
