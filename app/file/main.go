// Code generated by kitex v0.9.1, Customize by suyiiyii at https://github.com/suyiiyii/cwgo-template
package main

import (
	"net"
	"os"

	"github.com/doutokk/doutok/app/file/conf"
	"github.com/doutokk/doutok/app/file/biz/dal"
	"github.com/doutokk/doutok/app/file/biz/dal/mysql"
	"github.com/doutokk/doutok/app/file/biz/dal/query"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/file/fileservice"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

func main() {
	// use `go run cmd/gorm/main.go` to migrate the database
	dal.Init()
	// use `go run cmd/gorm_gen/main.go` to generate the code
	query.SetDefault(mysql.DB)
	opts := kitexInit()

	svr := fileservice.NewServer(new(FileServiceImpl), opts...)

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

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	klog.SetOutput(os.Stdout)
	return
}
