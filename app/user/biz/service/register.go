package service

import (
	"context"
	"errors"
	"github.com/doutokk/doutok/app/user/biz/dal/model"
	"github.com/doutokk/doutok/app/user/biz/dal/query"
	"github.com/doutokk/doutok/app/user/infra/rpc"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/auth"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/user"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
}

// NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	u := query.User
	count, err := query.Q.User.Where(u.Email.Eq(req.Email)).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("email already exists")
	}

	nu := &model.User{
		Email:          req.Email,
		HashedPassword: string(hashedPassword),
	}

	tx := query.Q.Begin()
	err = query.Q.User.Create(nu)
	if err != nil {
		return
	}
	resp = &user.RegisterResp{
		UserId: int32(nu.ID),
	}

	roleReq := &auth.CreateUserRoleReq{
		UserId: int32(nu.ID),
		Role:   "user",
	}
	roleResp, err := rpc.AuthClient.CreateUserRole(s.ctx, roleReq)

	if err != nil || !roleResp.Res {
		tx.Rollback()
	}
	tx.Commit()

	return
}
