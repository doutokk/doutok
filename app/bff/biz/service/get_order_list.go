package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	bff "github.com/doutokk/doutok/app/bff/hertz_gen/bff"
)

type GetOrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetOrderListService(Context context.Context, RequestContext *app.RequestContext) *GetOrderListService {
	return &GetOrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *GetOrderListService) Run(req *bff.Null) (resp *bff.GetOrderListResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
