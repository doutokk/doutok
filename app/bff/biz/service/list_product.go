package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	bff "github.com/doutokk/doutok/app/bff/hertz_gen/bff"
)

type ListProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewListProductService(Context context.Context, RequestContext *app.RequestContext) *ListProductService {
	return &ListProductService{RequestContext: RequestContext, Context: Context}
}

func (h *ListProductService) Run(req *bff.ListProductReq) (resp *bff.ListProductResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
