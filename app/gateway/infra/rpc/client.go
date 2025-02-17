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
	"context"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/doutokk/doutok/app/gateway/infra/mtl"
	gatewayutils "github.com/doutokk/doutok/app/gateway/utils"
	"github.com/doutokk/doutok/common/clientsuite"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/order/orderservice"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/user/userservice"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
)

var (
	ProductClient  productcatalogservice.Client
	UserClient     userservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	err            error
	registryAddr   string
	commonSuite    client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = "consul:8500"
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: gatewayutils.ServiceName,
		})
		initProductClient()
		initUserClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
	})
}

func initProductClient() {
	var opts []client.Option

	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
		return circuitbreak.RPCInfo2Key(ri)
	})
	cbs.UpdateServiceCBConfig("shop-gateway/product/GetProduct", circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})

	opts = append(opts, commonSuite, client.WithCircuitBreaker(cbs), client.WithFallback(fallback.NewFallbackPolicy(fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
		methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
		if err == nil {
			return resp, err
		}
		if methodName != "ListProducts" {
			return resp, err
		}
		return &product.ListProductsResp{
			Item: []*product.Product{
				{
					Price:       6.6,
					Id:          3,
					Picture:     "/static/image/t-shirt.jpeg",
					Name:        "T-Shirt",
					Description: "CloudWeGo T-Shirt",
				},
			},
		}, nil
	}))))
	opts = append(opts, client.WithTracer(prometheus.NewClientTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))))

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	gatewayutils.MustHandleError(err)
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	gatewayutils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	gatewayutils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", commonSuite)
	gatewayutils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	gatewayutils.MustHandleError(err)
}
