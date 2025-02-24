package utils

import (
	"context"
	"fmt"
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
	// 从请求的上下文中提取 metadata
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		// 如果没有 metadata，返回 0 或其他默认值
		return 0
	}

	// 获取 "userId" 字段
	userIds := md["user-id"]
	if len(userIds) == 0 {
		return 0
	}

	// 尝试将 userId 转换为 int 类型
	userIdStr := userIds[0]
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		// 如果转换失败，返回 0 或其他默认值
		fmt.Printf("Error converting userId to int: %v\n", err)
		return 0
	}

	// 成功转换为 int，返回 userId
	return userId
}

func GetUserIdFromCtx(ctx context.Context) int {
	return ctx.Value("user_id").(int)
}
