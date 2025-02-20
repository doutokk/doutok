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
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/cors"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/hertz-contrib/registry/consul"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
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
			fmt.Println("无法写入文件:", errr)
			return
		}
		fmt.Println("文件已写入:", filePath)
	} else {
		fmt.Println("文件已存在:", filePath)
	}
}

func main() {
	// build a consul client

	writeFile("conf/model.conf", modelFile)
	writeFile("conf/policy.csv", policyFile)

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

	// 创建反向代理池子
	proxyPool.Init()

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
	registerMiddleware(h)

	h.Use(func(ctx context.Context, c *app.RequestContext) {
		hlog.Info("孩子们这里我做了一些权限校验的东西，并且测试一下请求的顺序，1")
	})

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"ping": "pong1"})
	})

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
		serviceName := proxyPool.GetTargetServiceName(c.Request.URI().String())
		proxy := proxyPool.GetProxy(serviceName)
		proxy.SetDirector(func(req *protocol.Request) {
			req.SetHost(proxyPool.GetHost(serviceName))
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
