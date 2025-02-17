package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/doutokk/doutok/app/gateway/conf"
	"github.com/doutokk/doutok/app/gateway/infra/proxyPool"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/cors"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/hertz-contrib/registry/consul"
	"log"
	"net"
	"os"
	"strings"
)

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

func main() {
	// build a consul client
	ip, err := GetOutboundIP()
	addr := fmt.Sprintf("%s:8887", ip.String())

	// 创建反向代理池子
	proxyPool.Init()

	if err != nil {
		log.Fatal(err)
	}

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
		server.WithRegistry(r, &registry.Info{
			ServiceName: "gateway",
			Addr:        utils.NewNetAddr("tcp", addr),
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

	h.Any("/*path", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping3": "pong1231231"})
		req := c.GetRequest()
		proxy := proxyPool.GetProxy(getTargetServiceName(req.URI().String()))
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

func getTargetServiceName(uri string) string {
	// eg: http://10.21.32.14:8887/user/login
	hlog.Info("req_uri:  " + uri)
	parts := strings.Split(uri, "/")
	targetServiceName := parts[3]
	return targetServiceName
}
