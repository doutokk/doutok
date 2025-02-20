package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func GetUserIdRequest(c *app.RequestContext) int {
	strId := c.GetHeader("userId")
	userId, _ := strconv.Atoi(string(strId))
	return userId
}
