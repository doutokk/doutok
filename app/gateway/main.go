package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/doutokk/doutok/app/gateway/conf"
	"github.com/doutokk/doutok/app/gateway/infra/rpc"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/auth"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/cors"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/hertz-contrib/registry/consul"
	"github.com/hertz-contrib/reverseproxy"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 自定义错误类型
type BackendError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Body       string `json:"body"`
	Errors     string `json:"errors,omitempty"` // 如果没有错误，则忽略该字段
}

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

func checkAuth(ctx context.Context, c *app.RequestContext) bool {
	// 获取请求头中的 Authorization 字段
	authorization := string(c.Request.Header.Peek("Authorization"))
	req := new(auth.VerifyTokenReq)
	req.Token = authorization
	req.Method = string(c.Method())
	req.Uri = getPath(c.Request.URI().String())
	// 验证 token
	resp, err := rpc.AuthClient.VerifyTokenByRPC(ctx, req)
	if err != nil || !resp.Res {
		return false
	}
	userId := int(resp.UserId)
	c.Request.Header.Set("User-Id", fmt.Sprintf("%d", userId))
	return true
}

func allowCors(c *app.RequestContext) {
	// 允许跨域请求
	// 设置允许携带凭证（cookies）
	c.Header("Access-Control-Allow-Credentials", "true")

	// 获取请求的来源地址
	origin := c.Request.Header.Get("Origin")

	// 根据需求放宽或限制允许的源，这里假设只允许特定源
	if origin != "" {
		c.Header("Access-Control-Allow-Origin", origin)
	}

	// 设置允许的方法
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	// 设置允许的请求头
	c.Header("Access-Control-Allow-Headers", "Authorization, transfer, session, Content-Type, Accept, Origin, X-Requested-With, token, id, X-Custom-Header, X-Cookie, Connection, User-Agent, Cookie")

	// 设置预检请求的有效期
	c.Header("Access-Control-Max-Age", "3600")

	// 允许浏览器读取的所有响应头
	c.Header("Access-Control-Expose-Headers", "*")
}

func main() {
	rpc.InitClient()

	ip, err := GetOutboundIP()
	port := conf.GetConf().Kitex.Address
	addr := fmt.Sprintf("%s"+port, ip.String())

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
		server.WithHostPorts(port),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "gateway",
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}),
	)

	// todo:为什么我 h.Use 的没被执行
	h.Use(
		cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowMethods: []string{
				"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS",
			},
			AllowHeaders: []string{
				"Content-Type",
				"Authorization",
				"X-Token",
				"Refer",
				"Origin",
				"Zy-Cookie",
			},
			ExposeHeaders: []string{
				"Content-Length",
				"Access-Control-Allow-Origin",
				"Access-Control-Allow-Headers",
				"Access-Control-Request-Headers",
				"Access-Control-Expose-Headers",
				"Content-Type",
				"Zy-Cookie",
			},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong1"})
	})

	if err != nil {
		log.Fatal(err)
	}
	registerMiddleware()
	targetHost := conf.GetConf().Gateway.GrpcGatewayAddr
	proxy, _ := reverseproxy.NewSingleHostReverseProxy(targetHost)
	// 定义路由，匹配所有路径
	h.Any("/*path", func(ctx context.Context, c *app.RequestContext) {
		defer allowCors(c)
		if string(c.Method()) == "OPTIONS" {
			c.SetStatusCode(204)
			return
		}

		allow := checkAuth(ctx, c)
		if !allow {
			c.JSON(401, utils.H{"error": "Unauthorized"})
			return
		}

		path := getPath(c.Request.URI().String())
		proxy.SetDirector(func(req *protocol.Request) {
			req.SetHost(targetHost + path)
			req.SetRequestURI("http://" + targetHost + path)
		})

		// 调用反向代理处理请求
		proxy.ServeHTTP(ctx, c)

		// todo:后端服务需要把错误返回

		// todo:鉴权
		if c.Response.StatusCode() != 200 {
			// 读取响应 Body
			bodyBytes := c.Response.Body()
			bodyString := string(bodyBytes)

			// 读取 Errors
			errorsString := c.Errors.String()

			// 创建自定义错误
			backendErr := &BackendError{
				StatusCode: c.Response.StatusCode(),
				Message:    "Backend service returned an error",
				Body:       bodyString,
				Errors:     errorsString,
			}

			jsonErr, _ := json.Marshal(backendErr)
			// 将错误存储到 context 中
			c.Response.SetBody(jsonErr)
		}
	})

	h.Spin()
}

func getPath(rawURL string) string {
	// 按 "://" 分割，去掉协议头
	parts := strings.SplitN(rawURL, "://", 2)
	if len(parts) > 1 {
		rawURL = parts[1]
	}

	// 按 "/" 分割，提取第一个 "/" 及后面的部分
	index := strings.Index(rawURL, "/")
	return rawURL[index:]
}

// todo:往里面加多点好东西
func registerMiddleware() {
	// access log
	logger := hertzzap.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(hlog.LevelInfo)
	hlog.SetOutput(os.Stdout)
}
