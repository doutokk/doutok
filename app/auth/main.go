package main

import (
	"github.com/doutokk/doutok/app/auth/biz/dal"
	"github.com/doutokk/doutok/app/auth/conf"
	"github.com/doutokk/doutok/app/auth/infra/casbin"
	"github.com/doutokk/doutok/common/mtl"
	"github.com/doutokk/doutok/common/serversuite"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/auth/authservice"
	"github.com/joho/godotenv"
	"net"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	_ = godotenv.Load()
	mtl.InitTracing(serviceName, conf.GetConf().Kitex.OtlpAddr)
	mtl.InitMetric(serviceName, "8383", conf.GetConf().Registry.RegistryAddress[0])
	opts := kitexInit()
	dal.Init()
	casbin.Init()

	svr := authservice.NewServer(new(AuthServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	klog.SetOutput(os.Stdout)

	// address
	address := conf.GetConf().Kitex.Address
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))
	opts = append(opts, server.WithSuite(serversuite.CommonServerSuite{CurrentServiceName: serviceName, RegistryAddr: conf.GetConf().Registry.RegistryAddress[0]}))
	return
}
