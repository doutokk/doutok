package proxyPool

import (
	"github.com/doutokk/doutok/app/gateway/conf"
	"github.com/hertz-contrib/reverseproxy"
)

var (
	pool    map[string]*reverseproxy.ReverseProxy
	hostMap = make(map[string]string)
)

func GetProxy(name string) *reverseproxy.ReverseProxy {
	return pool[name]
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
		pool[name] = proxy
	}

}
