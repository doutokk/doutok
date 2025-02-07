package service

import (
	"context"
	"github.com/PengJingzhao/douyin-commerce/app/auth/biz/utils"
	auth "github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/auth"
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

	_, err = utils.ValidateJWT(token)
	if err != nil {
		return &auth.VerifyResp{Res: false}, err
	}

	return &auth.VerifyResp{Res: true}, nil
}
