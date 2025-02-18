// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"github.com/doutokk/doutok/app/gateway/conf"
	"github.com/doutokk/doutok/common/utils"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/payment/paymentservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/doutokk/doutok/common/clientsuite"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/order/orderservice"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product/productcatalogservice"
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
		initCartClient()
		initProductClient()
		initOrderClient()
		initPaymentClient()
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
