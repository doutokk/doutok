package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	bff "github.com/doutokk/doutok/app/bff/hertz_gen/bff"
)

type PaymentService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPaymentService(Context context.Context, RequestContext *app.RequestContext) *PaymentService {
	return &PaymentService{RequestContext: RequestContext, Context: Context}
}

func (h *PaymentService) Run(req *bff.PaymentReq) (resp *bff.Null, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
