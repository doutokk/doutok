package service

import (
	"context"
	"errors"
	"github.com/PengJingzhao/douyin-commerce/app/user/biz/dal/query"
	user "github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	// 通过邮件查找用户
	existingUser, err := query.Q.User.GetOneByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}
	//if err := mysql.DB.Where("email=?", req.Email).First(&existingUser).Error; err != nil {
	//	return nil, errors.New("invalid email or password")
	//}

	// 比对密码是否正确
	if existingUser.Password != req.Password {
		return nil, errors.New("password error")
	}
	return &user.LoginResp{UserId: int32(existingUser.ID)}, nil
}
