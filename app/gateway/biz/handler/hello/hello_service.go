package hello

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/doutokk/doutok/app/gateway/biz/service"
	"github.com/doutokk/doutok/app/gateway/biz/utils"
	hello "github.com/doutokk/doutok/app/gateway/hertz_gen/cwgo/http/hello"
)

// Method1 .
// @router /hello [GET]
func Method1(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hello.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &hello.HelloResp{}
	resp, err = service.NewMethod1Service(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
