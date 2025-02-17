package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/kr/pretty"
	"testing"
)

func TestMain_Run(t *testing.T) {
	c, err := cartservice.NewClient("cart", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}
	resp, err := c.GetCart(context.TODO(), &cart.GetCartReq{
		UserId: 1,
	})
	if err != nil {
		return
	}

	_, _ = pretty.Println(resp.Cart)

}
