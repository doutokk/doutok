package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/doutokk/doutok/gateway/proto"
	cart_pb "github.com/doutokk/doutok/gateway/proto/cart"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterFileServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// err = cart_pb.RegisterCartServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	err = cart_pb.RegisterCartServiceHandlerFromEndpoint(ctx, mux, "cart-service:8888", opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()

	if err := run(); err != nil {
		panic(err)
	}
}
