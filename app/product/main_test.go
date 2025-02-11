package main

import (
	"context"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/product"
	"github.com/PengJingzhao/douyin-commerce/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain_Run(t *testing.T) {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		t.Fatal(err)
	}

	c, err := productcatalogservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		t.Fatal(err)
	}

	resp1, err := c.GetProduct(context.TODO(), &product.GetProductReq{Id: 1})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, resp1)

	resp2, err := c.ListProducts(context.TODO(), &product.ListProductsReq{Page: 1, PageSize: 10})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, resp2)

	resp3, err := c.SearchProducts(context.TODO(), &product.SearchProductsReq{Query: "123"})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, resp3)
}
