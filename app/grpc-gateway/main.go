package main

import (
	"context"
	"fmt"
	"github.com/doutokk/doutok/app/order/grpc-gateway/pb/cart"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Authentication middleware
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("asdhoasildjaskldjaslkdjaslkdjas")
	})
}

func run() (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(
		insecure.NewCredentials())}

	//err = cartpb.RegisterCartServiceHandlerFromEndpoint(ctx, mux, "10.21.32.14:8886", opts)
	err = cart.RegisterCartServiceHandlerFromEndpoint(ctx, mux, "cart-service:8888", opts)
	//err = userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "10.21.32.14:8888", opts)
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
