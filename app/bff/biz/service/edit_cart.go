package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	bff "github.com/doutokk/doutok/app/bff/hertz_gen/bff"
)

type EditCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewEditCartService(Context context.Context, RequestContext *app.RequestContext) *EditCartService {
	return &EditCartService{RequestContext: RequestContext, Context: Context}
}

func (h *EditCartService) Run(req *bff.EditCartReq) (resp *bff.Null, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
