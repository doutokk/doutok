package service

import (
	"context"
	"errors"
	"github.com/doutokk/doutok/app/user/biz/dal/model"
	"github.com/doutokk/doutok/app/user/biz/dal/query"
	user "github.com/doutokk/doutok/rpc_gen/kitex_gen/user"

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
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password and confirm password not match")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u := &model.User{
		Email:          req.Email,
		HashedPassword: string(hashedPassword),
	}
	err = query.Q.User.Create(u)
	if err != nil {
		return
	}
	resp = &user.RegisterResp{
		UserId: int32(u.ID),
	}

	return
}
