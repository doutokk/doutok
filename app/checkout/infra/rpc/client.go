package rpc

import (
	"github.com/doutokk/doutok/app/checkout/conf"
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
	registryAddr  string
	serviceName   string
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Registry.RegistryAddress[0]
		serviceName = conf.GetConf().Kitex.Service
		initCartClient()
		initProductClient()
		initOrderClient()
	})
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart")
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product")
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order")
}

func initPaymentClient() {
	PaymentClient, err = paymentservice.NewClient("payment")
}
