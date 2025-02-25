package main

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/doutokk/doutok/app/gateway/conf"
	"github.com/doutokk/doutok/app/gateway/infra/proxyPool"
	"github.com/doutokk/doutok/app/gateway/infra/rpc"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/auth"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/cors"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/hertz-contrib/registry/consul"
	"github.com/hertz-contrib/reverseproxy"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	//go:embed conf/model.conf
	modelFile []byte
	//go:embed conf/policy.csv
	policyFile []byte
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

// Casbin 中间件 todo:应该给 auth 做这件事
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

func writeFile(filePath string, byteFile []byte) {

	// 获取目录路径
	dirPath := filepath.Dir(filePath)

	// 检查目录是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		// 目录不存在，创建目录
		errr := os.MkdirAll(dirPath, 0755) // 0755 是权限
		if errr != nil {
			panic(errr)
		}
		fmt.Printf("目录 %s 创建成功\n", dirPath)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 如果文件不存在，则创建并写入嵌入的内容
		errr := ioutil.WriteFile(filePath, byteFile, 0644)
		if errr != nil {
			fmt.Println("无法写入文件：", errr)
			return
		}
		fmt.Println("文件已写入：", filePath)
	} else {
		fmt.Println("文件已存在：", filePath)
	}
}

func checkAuth(ctx context.Context, c *app.RequestContext) bool {

	// todo:测试，实际要在 auth 那里用 casbin
	if proxyPool.GetTargetServiceName(c.Request.URI().String()) == "user" {
		return true
	}

	// 获取请求头中的 Authorization 字段
	authorization := string(c.Request.Header.Peek("Authorization"))
	req := new(auth.VerifyTokenReq)
	req.Token = authorization

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
	writeFile("conf/model.conf", modelFile)
	writeFile("conf/policy.csv", policyFile)
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

	// 鉴权中间件
	enforcer, err := casbin.NewEnforcer("conf/model.conf", "conf/policy.csv")
	if err != nil {
		log.Fatal(err)
	}
	h.Use(CasbinMiddleware(enforcer))
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
