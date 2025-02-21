package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func GetUserIdRequest(c *app.RequestContext) int {
	strId := c.Request.Header.Get("userID")
	userId, _ := strconv.Atoi(string(strId))
	return userId
}

func getUserIdFromCtx(ctx context.Context) int {
	return ctx.Value("userID").(int)
}
