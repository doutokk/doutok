package proxyPool

import (
	"github.com/doutokk/doutok/app/gateway/conf"
	"github.com/hertz-contrib/reverseproxy"
)

var (
	pool   map[string]*reverseproxy.ReverseProxy
	suffix = conf.GetConf().Gateway.Suffix
)

func GetProxy(name string) *reverseproxy.ReverseProxy {
	return pool[name+suffix]
}

func Init() {
	pool = make(map[string]*reverseproxy.ReverseProxy)
	paths := conf.GetConf().Gateway.ServicePath

	for _, path := range paths {
		proxy, _ := reverseproxy.NewSingleHostReverseProxy(path)
		pool[path] = proxy
	}

}
