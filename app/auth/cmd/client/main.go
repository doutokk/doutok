package main

import (
	"context"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/auth"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/auth/authservice"
	client2 "github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
)

func main() {
	r, err := consul.NewConsulResolver("10.21.32.14:8600")
	if err != nil {
		log.Fatal(err)
	}
	cli, err := authservice.NewClient("auth.server", client2.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.DeliverTokenByRPC(context.TODO(), &auth.DeliverTokenReq{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
