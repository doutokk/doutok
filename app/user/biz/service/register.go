package service

import (
	"context"
	"errors"
	"github.com/PengJingzhao/douyin-commerce/app/auth/biz/model"
	"github.com/PengJingzhao/douyin-commerce/app/user/biz/dal/mysql"
	user "github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/user"
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
	var existingUser model.User
	if err = mysql.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		// 数据库查询到该邮箱
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
