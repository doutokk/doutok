package main

import (
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/doutokk/doutok/app/cart/biz/dal"
	"github.com/doutokk/doutok/app/cart/biz/dal/mysql"
	"github.com/doutokk/doutok/app/cart/biz/dal/query"
	"github.com/doutokk/doutok/app/cart/conf"
	"github.com/doutokk/doutok/app/cart/infra/rpc"
	"github.com/doutokk/doutok/common/mtl"
	"github.com/doutokk/doutok/common/serversuite"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/cart/cartservice"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	// use go run cmd/gorm/main.go to migrate the database
	dal.Init()
	// use go run cmd/gorm_gen/main.go to generate the code
	query.SetDefault(mysql.DB)
	rpc.InitClient()
	mtl.InitTracing(serviceName, conf.GetConf().Kitex.OtlpAddr)
	mtl.InitMetric(serviceName, "8383", conf.GetConf().Registry.RegistryAddress[0])
	rpc.InitClient()
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
