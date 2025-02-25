package service

import (
	"context"
	"github.com/doutokk/doutok/app/auth/infra/casbin"
	auth "github.com/doutokk/doutok/rpc_gen/kitex_gen/auth"
	"strconv"
)

type CreateUserRoleService struct {
	ctx context.Context
}

// NewCreateUserRoleService new CreateUserRoleService
func NewCreateUserRoleService(ctx context.Context) *CreateUserRoleService {
	return &CreateUserRoleService{ctx: ctx}
}

// Run create note info
func (s *CreateUserRoleService) Run(req *auth.CreateUserRoleReq) (resp *auth.CreateUserRoleResp, err error) {
	// Finish your business logic.
	userId := strconv.Itoa(int(req.UserId))
	err = casbin.CreateUserRole(userId, req.Role)
	if err != nil {
		return nil, err
	}
	resp = &auth.CreateUserRoleResp{
		Res: true,
	}
	return
}
