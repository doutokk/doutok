package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/user"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/user/userservice"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain_Run(t *testing.T) {
	// create the client
	r, err2 := consul.NewConsulResolver("127.0.0.1:8500")
	if err2 != nil {
		panic(err2)
	}
	c, err := userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		t.Fatal(err)
	}

	// call the server method
	resp, err := c.Login(context.TODO(), &user.LoginReq{Email: "p.vucvinyih@qq.com", Password: "abc"})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, resp.UserId)

	// call the register method
	register, err := c.Register(context.Background(), &user.RegisterReq{Email: "p.vucvinyih@qq.com", Password: "abc", ConfirmPassword: "abc"})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, register.UserId)

}
