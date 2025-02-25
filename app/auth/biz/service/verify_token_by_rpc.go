package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/auth/biz/utils"
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
	// Finish your business logic.
	token := req.Token

	result, err := utils.ValidateJWT(token)
	if err != nil {
		klog.Warnf("Token validate failed: %v %v", err, token)
		return &auth.VerifyResp{Res: false}, err
	}

	userId := result.UserID
	klog.Warnf("Token validate success: %v %v", userId, token)
	return &auth.VerifyResp{
		Res:    true,
		UserId: int32(userId),
	}, nil
}
