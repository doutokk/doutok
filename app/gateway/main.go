package main

import (
	"context"
	"github.com/doutokk/doutok/app/product/conf"
	"log"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
)

func main() {
	// build a consul client
	config := consulapi.DefaultConfig()
	config.Address = conf.GetConf().Registry.RegistryAddress[0]
	consulClient, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	// build a consul register with the consul client
	r := consul.NewConsulRegister(consulClient)

	// run Hertz with the consul register
	addr := "127.0.0.1:8887"
	h := server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "hertz.test.demo",
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}),
	)
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong1"})
	})

	h.Any("/", func(ctx context.Context, c *app.RequestContext) {
		req := c.GetRequest()
		targetServiceUrl := getTargetServiceUrl(req.URI().String())
		c.Redirect(consts.StatusMovedPermanently, []byte(targetServiceUrl))
	})

	h.Spin()
}

func getTargetServiceUrl(uri string) string {
	parts := strings.Split(uri, "/")
	//targetServiceName := parts[1] + "-service"
	var targetUri string
	for i := 2; i < len(parts); i++ {
		targetUri += parts[i] + "/"
	}
	return "10.21.32.14" + ":8888/" + targetUri
}
