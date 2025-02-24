package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func GetUserIdRequest(c *app.RequestContext) int {
	strId := c.Request.Header.Get("user_id")
	userId, _ := strconv.Atoi(string(strId))
	return userId
}

func GetUserId(ctx context.Context) int {
	return 1
}

func GetUserIdFromCtx(ctx context.Context) int {
	return ctx.Value("user_id").(int)
}
