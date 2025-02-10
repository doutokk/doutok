package main

import (
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal"
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal/mysql"
	"github.com/PengJingzhao/douyin-commerce/app/cart/biz/dal/query"
	"github.com/PengJingzhao/douyin-commerce/app/cart/conf"
	"github.com/PengJingzhao/douyin-commerce/common/serversuite"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"net"
	"os"
)

func main() {
	dal.Init()
	query.SetDefault(mysql.DB)
	opts := kitexInit()

	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// registry
	opts = append(opts,
		server.WithSuite(serversuite.CommonServerSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
		}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	klog.SetOutput(os.Stdout)
	return
}
