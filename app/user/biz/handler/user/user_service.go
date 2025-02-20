// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/doutokk/doutok/app/user/biz/service"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/user"
)

// Register .
// @router /user/register [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	s := service.NewRegisterService(ctx)
	resp, err := s.Run(&req)

	if err != nil {
		c.JSON(consts.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /user/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	s := service.NewLoginService(ctx)
	resp, err := s.Run(&req)

	if err != nil {
		c.JSON(consts.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(consts.StatusOK, resp)
}
