package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	hello "github.com/doutokk/doutok/app/gateway/hertz_gen/cwgo/http/hello"
)

type Method1Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMethod1Service(Context context.Context, RequestContext *app.RequestContext) *Method1Service {
	return &Method1Service{RequestContext: RequestContext, Context: Context}
}

func (h *Method1Service) Run(req *hello.HelloReq) (resp *hello.HelloResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
