package main

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cartpb "github.com/doutokk/doutok/app/order/grpc-gateway/pb/cart"
)

func run() (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
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
