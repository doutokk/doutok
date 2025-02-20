package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/doutokk/doutok/app/cart/conf"
	"github.com/doutokk/doutok/common/clientsuite"
	"github.com/doutokk/doutok/common/utils"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/order/orderservice"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	OrderClient   orderservice.Client
	PaymentClient paymentservice.Client
	once          sync.Once
	err           error
	commonSuite   client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr := conf.GetConf().Registry.RegistryAddress[0]
		serviceName := conf.GetConf().Kitex.Service
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: serviceName,
		})
		initPaymentClient()
		initCartClient()
		initProductClient()
		initOrderClient()
	})
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	utils.PassOrPanic(err)
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", commonSuite)
	utils.PassOrPanic(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	utils.PassOrPanic(err)
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment", commonSuite)
	utils.PassOrPanic(err)
}
