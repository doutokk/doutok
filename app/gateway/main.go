package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/doutokk/doutok/app/gateway/conf"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/cors"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/hertz-contrib/registry/consul"
	"github.com/hertz-contrib/reverseproxy"
	"log"
	"os"
	"strings"
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
	h := server.New(
		server.WithHostPorts(":8887"),
		server.WithHandleMethodNotAllowed(true),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "hertz.test.demo",
			Addr:        utils.NewNetAddr("tcp", "localhost:8887"),
			Weight:      10,
			Tags:        nil,
		}),
	)
	registerMiddleware(h)

	h.Use(func(ctx context.Context, c *app.RequestContext) {
		hlog.Info("孩子们这里我做了一些权限校验的东西，并且测试一下请求的顺序，1")
		c.Next(ctx)
	})

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {

		c.JSON(consts.StatusOK, utils.H{"ping": "pong1"})
	})

	proxy, err := reverseproxy.NewSingleHostReverseProxy("http://10.21.32.14:8888/")

	h.Any("/*path", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping3": "pong1231231"})
		req := c.GetRequest()
		targetServiceUrl := getTargetServiceUrl(req.URI().String())
		hlog.Info("targetServiceUrl: ", targetServiceUrl)

		proxy.ServeHTTP(ctx, c)
		c.Next(ctx)
	})

	h.Spin()
}

// todo:往里面加多点好东西
func registerMiddleware(h *server.Hertz) {
	// access log
	logger := hertzzap.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelInfo)
	hlog.SetOutput(os.Stdout)

	// cores
	h.Use(cors.Default())
}

func getTargetServiceUrl(uri string) string {
	hlog.Info(uri)
	parts := strings.Split(uri, "/")
	targetServiceName := parts[3] + "-service"
	hlog.Info(targetServiceName)
	var targetUri string
	for i := 3; i < len(parts); i++ {
		targetUri += parts[i] + "/"
	}
	return "http://10.21.32.14" + ":8888/" + targetUri
}
