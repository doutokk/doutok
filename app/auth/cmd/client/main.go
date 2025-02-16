package main

import (
	"context"
	client2 "github.com/cloudwego/kitex/client"
	"github.com/doutokk/doutok/app/auth/conf"
	"github.com/doutokk/doutok/app/auth/kitex_gen/auth"
	"github.com/doutokk/doutok/app/auth/kitex_gen/auth/authservice"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
)

func main() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		log.Fatal(err)
	}
	cli, err := authservice.NewClient("auth", client2.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.DeliverTokenByRPC(context.TODO(), &auth.DeliverTokenReq{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
