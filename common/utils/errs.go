package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func PassOrPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func serverError(ctx context.Context, c *app.RequestContext, err error, statusCode ...int) {
	code := consts.StatusInternalServerError
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	c.Response.SetStatusCode(code)
	c.Response.SetBody([]byte(err.Error()))
	c.Abort()
}
