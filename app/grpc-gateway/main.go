package main

import (
	"context"
	"net/http"

	cartpb "github.com/doutokk/doutok/app/grpc-gateway/pb/cart"
	orderpb "github.com/doutokk/doutok/app/grpc-gateway/pb/order"
	productpb "github.com/doutokk/doutok/app/grpc-gateway/pb/product"
	userpb "github.com/doutokk/doutok/app/grpc-gateway/pb/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 添加认证中间件
func CustomMatcher(key string) (string, bool) {
	switch key {
	case "User-Id":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func run() (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 使用自定义选项创建 ServeMux
	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(CustomMatcher),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = cartpb.RegisterCartServiceHandlerFromEndpoint(ctx, mux, "cart-service:8888", opts)
	err = orderpb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, "order-service:8888", opts)
	err = userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "user-service:8888", opts)
	err = productpb.RegisterProductCatalogServiceHandlerFromEndpoint(ctx, mux, "product-service:8888", opts)

	if err != nil {
		return err
	}

	return http.ListenAndServe("0.0.0.0:8888", mux)
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
