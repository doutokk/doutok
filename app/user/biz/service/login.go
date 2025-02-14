package service

import (
	"context"
	"errors"
	"github.com/doutokk/doutok/app/user/biz/dal/query"
	user "github.com/doutokk/doutok/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
}

// NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.

	u := query.User
	us, err := query.Q.User.Where(u.Email.Eq(req.Email)).First()
	if err != nil {
		return nil, errors.New("user not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(us.HashedPassword), []byte(req.Password)) != nil {
		return nil, errors.New("password is incorrect")
	}

	resp = &user.LoginResp{
		UserId: int32(us.ID),
	}

	return
}
