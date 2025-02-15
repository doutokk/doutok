package service

import (
	"context"
	"github.com/doutokk/doutok/app/frontend/infra/rpc"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	bff "github.com/doutokk/doutok/app/bff/hertz_gen/bff"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

var userClient = rpc.UserClient

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *bff.LoginReq) (resp *bff.LoginResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	request := &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	}

	response, err := userClient.Login(h.Context, request)
	resp = &bff.LoginResp{
		UserId: response.UserId,
	}

	return
}
