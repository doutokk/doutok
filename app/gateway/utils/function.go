package utils

import (
	"context"
	"strconv"
)

func GetUserIdFromCtx(ctx context.Context) uint32 {
	userIDStr := ctx.Value("user_id").(string)
	userIDUint64, _ := strconv.ParseUint(userIDStr, 10, 32)
	return uint32(userIDUint64)
}
