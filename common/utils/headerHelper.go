package utils

import (
	"context"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"strconv"
)

func GetUserId(ctx context.Context) int {

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return 0
	}
	userId, _ := strconv.Atoi(md.Get("user-id")[0])

	return userId
}
