package main

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol"
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

var (
	textCnt = 0
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

// Casbin 中间件
func CasbinMiddleware(e *casbin.Enforcer) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// todo:不硬编码获取角色
		sub := "alice"
		obj := string(c.Request.RequestURI())
		act := string(c.Request.Method())
		hlog.Info("sub: ", sub, " obj: ", obj, " act: ", act)

		ok, err := e.Enforce(sub, obj, act)
		if err != nil {
			c.AbortWithError(consts.StatusInternalServerError, err)
			hlog.Error(err)
			return
		}

		if !ok {
			c.AbortWithMsg("Forbidden", consts.StatusForbidden)
			return
		}

		c.Next(ctx)
	}
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
	})

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong1"})
	})

	if err != nil {
		// 处理错误
	}

	// 鉴权中间件
	enforcer, err := casbin.NewEnforcer("conf/model.conf", "conf/policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	h.Use(CasbinMiddleware(enforcer))

	// 定义路由，匹配所有路径
	h.Any("/*path", func(ctx context.Context, c *app.RequestContext) {

		// 打印请求的 URI
		hlog.Info("path: ", c.Request.URI())
		serviceName := getTargetServiceName(c.Request.URI().String())
		proxy := proxyPool.GetProxy(serviceName)
		proxy.SetDirector(func(req *protocol.Request) {
			req.SetHost(proxyPool.GetHost(serviceName))
		})

		// 调用反向代理处理请求
		proxy.ServeHTTP(ctx, c)

		// todo:后端服务需要把错误返回

		// todo:鉴权

		if c.Response.StatusCode() != 200 {
			err = fmt.Errorf("请求失败，状态码：%d", c.Response.StatusCode())
		}
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
