package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/doutokk/doutok/app/payment/biz/dal"
	"github.com/doutokk/doutok/app/payment/biz/dal/mysql"
	"github.com/doutokk/doutok/app/payment/biz/dal/query"
	"github.com/doutokk/doutok/app/payment/conf"
	"github.com/doutokk/doutok/app/payment/infra/rpc"
	"github.com/doutokk/doutok/common/mtl"
	"github.com/doutokk/doutok/common/serversuite"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"net"
	"os"
)

var serviceName = conf.GetConf().Kitex.Service

func main() {
	_ = godotenv.Load()
	mtl.InitTracing(serviceName, conf.GetConf().Kitex.OtlpAddr)
	mtl.InitMetric(serviceName, "8383", conf.GetConf().Registry.RegistryAddress[0])
	dal.Init()
	rpc.InitClient()
	query.SetDefault(mysql.DB)
	opts := kitexInit()

	svr := paymentservice.NewServer(new(PaymentServiceImpl), opts...)

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
