package utils

import (
	"context"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"strconv"
)

func GetUserId(ctx *context.Context) uint32 {

	md, ok := metadata.FromIncomingContext(*ctx)

	if !ok {
		return 0
	}
	userId, _ := strconv.Atoi(md.Get("user-id")[0])
	// 继续传
	*ctx = metadata.AppendToOutgoingContext(*ctx, "user-id", strconv.Itoa(userId))

	return uint32(userId)
}
