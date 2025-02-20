package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/doutokk/doutok/app/auth/kitex_gen/auth/authservice"
	"github.com/doutokk/doutok/app/checkout/conf"
	"github.com/doutokk/doutok/common/clientsuite"
	"github.com/doutokk/doutok/common/utils"
	"sync"
)

var (
	AuthClient  authservice.Client
	once        sync.Once
	err         error
	commonSuite client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr := conf.GetConf().Registry.RegistryAddress[0]
		serviceName := conf.GetConf().Kitex.Service
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: serviceName,
		})
		initAuthClient()
	})
}

func initAuthClient() {
	AuthClient, err = authservice.NewClient("auth", commonSuite)
	utils.PassOrPanic(err)
}
