// Code generated by hertz generator.

package auth

import (
	"context"
	"douyin-commerce/biz/model/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

// 实现 token 分发
func DeliverToken(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	// 简单实现：用户 ID 转换成 token
	resp = &auth.DeliveryResp{
		Token: "testToken",
	}
	return resp, nil
}

// 实现 token 校验
func VerifyToken(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	// 假设 token 正确
	resp = &auth.VerifyResp{
		Res: req.Token != "",
	}
	return resp, nil
}

func DeliverTokenByRPC(ctx context.Context, c *app.RequestContext) {
	token, err := DeliverToken(ctx, &auth.DeliverTokenReq{})
	if err != nil {
		return
	}

	c.JSON(200, token)
}

func VerifyTokenByRPC(ctx context.Context, c *app.RequestContext) {

}
