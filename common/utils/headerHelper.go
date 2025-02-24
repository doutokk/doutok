package utils

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"google.golang.org/grpc/metadata"
	"strconv"
)

func GetUserIdRequest(c *app.RequestContext) int {
	strId := c.Request.Header.Get("user_id")
	userId, _ := strconv.Atoi(string(strId))
	return userId
}

func GetUserId(ctx context.Context) int {
	md, _ := metadata.FromIncomingContext(ctx)
	userId, _ := md["user_id"]
	intUserID, _ := strconv.Atoi(userId[0])
	return intUserID
}

func GetUserIdFromCtx(ctx context.Context) int {
	return ctx.Value("user_id").(int)
}
