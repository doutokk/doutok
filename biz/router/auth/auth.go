package auth

import (
	auth "douyin-commerce/biz/handler/auth"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(r *server.Hertz) {

	authGroup := r.Group("/auth")

	authGroup.POST("/DeliverTokenByRpc", auth.DeliverTokenByRPC)
	authGroup.POST("/VerifyTokenByRpc", auth.VerifyTokenByRPC)
}
