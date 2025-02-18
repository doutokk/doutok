package proxyPool

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/doutokk/doutok/app/gateway/conf"
	"github.com/hertz-contrib/reverseproxy"
	"strings"
)

var (
	pool    map[string]*reverseproxy.ReverseProxy
	hostMap = make(map[string]string)
)

// 测试过，挂了重启也能连上
func GetProxy(name string) *reverseproxy.ReverseProxy {
	return pool[name]
}

func GetTargetServiceName(uri string) string {
	// eg: http://10.21.32.14:8887/user/login
	hlog.Info("req_uri:  " + uri)
	parts := strings.Split(uri, "/")
	targetServiceName := parts[3]
	doubleParts := strings.Split(targetServiceName, "?")
	targetServiceName = doubleParts[0] // 无论有没有问号，都取第一个元素
	return targetServiceName
}

func GetHost(name string) string {
	return hostMap[name]
}

func Init() {
	pool = make(map[string]*reverseproxy.ReverseProxy)
	hosts := conf.GetConf().Gateway.ServiceHost

	for name, host := range hosts {
		hostMap[name] = host
		url := "http://" + host
		proxy, _ := reverseproxy.NewSingleHostReverseProxy(url)

		proxy.SetErrorHandler(func(c *app.RequestContext, err error) {
			if strings.Contains(err.Error(), "connection refused") {
				c.String(503, "Service Unavailable: Backend service is unavailable.")
			} else if strings.Contains(err.Error(), "connection has been closed by peer") {
				c.String(502, "Bad Gateway: The connection to the backend service was closed unexpectedly.")
			} else {
				c.String(502, "Bad Gateway: An error occurred while communicating with the backend service.")
			}
		})

		if proxy == nil {
			panic(name + "proxy is nil")
		}
		pool[name] = proxy
	}

}
