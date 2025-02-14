package service

import (
	"context"
	"errors"
	"github.com/doutokk/doutok/app/user/biz/dal/mysql"
	"github.com/doutokk/doutok/app/user/biz/dal/query"
	"github.com/doutokk/doutok/app/user/biz/model"
	user "github.com/doutokk/doutok/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// 用户注册
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	// 检验两次输入的密码是否一致
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("password do not match")
	}
	// 检查用户是否存在
	_, err = query.Q.User.GetOneByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already register")
	}

	// 创建用户
	newUser := model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	if err = mysql.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}
	//返回结果
	return &user.RegisterResp{
		UserId: int32(newUser.ID),
	}, nil
}
