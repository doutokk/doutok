package utils

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/metadata"
	"strconv"
)

func GetUserId(ctx context.Context) int {

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return 0
	}
	if len(md.Get("user-id")) == 0 {
		klog.Errorf("user-id not found in header")
		panic("user-id not found in header")
	}
	userId, _ := strconv.Atoi(md.Get("user-id")[0])

	return userId
}
