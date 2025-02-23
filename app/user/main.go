// Code generated by kitex v0.9.1, Customize by suyiiyii at https://github.com/suyiiyii/cwgo-template
package main

import (
	"github.com/doutokk/doutok/app/user/infra/rpc"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/user/userservice"
	"github.com/joho/godotenv"
	"net"
	"os"

	"github.com/doutokk/doutok/app/user/biz/dal"
	"github.com/doutokk/doutok/app/user/biz/dal/mysql"
	"github.com/doutokk/doutok/app/user/biz/dal/query"
	"github.com/doutokk/doutok/app/user/conf"
	"github.com/doutokk/doutok/common/mtl"
	"github.com/doutokk/doutok/common/serversuite"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	_ = godotenv.Load()
	dal.Init()
	query.SetDefault(mysql.DB)
	mtl.InitTracing(serviceName, conf.GetConf().Kitex.OtlpAddr)
	mtl.InitMetric(serviceName, "8383", conf.GetConf().Registry.RegistryAddress[0])
	rpc.InitClient()
	opts := kitexInit()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

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
