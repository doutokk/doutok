package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	cartpb "github.com/doutokk/doutok/app/grpc-gateway/pb/cart"
)

// 添加认证中间件
func withAuth() runtime.ServeMuxOption {
	return runtime.WithMetadata(func(ctx context.Context, req *http.Request) metadata.MD {
		md := make(map[string]string)
		if auth := req.Header.Get("Authorization"); auth != "" {
			md["authorization"] = auth
		}
		return metadata.New(md)
	})
}

func run() (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 使用自定义选项创建 ServeMux
	mux := runtime.NewServeMux(withAuth())
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = cartpb.RegisterCartServiceHandlerFromEndpoint(ctx, mux, "cart-service:8888", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe("0.0.0.0:8081", mux)
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
