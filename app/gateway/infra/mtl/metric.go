package mtl

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/doutokk/doutok/common/utils"

	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/route"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hertz-contrib/registry/consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

func initMetric() route.CtxCallback {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	config := consulapi.DefaultConfig()
	config.Address = "consul:8500"
	consulClient, _ := consulapi.NewClient(config)
	r := consul.NewConsulRegister(consulClient)

	localIp := utils.MustGetLocalIPv4()
	ip, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", localIp, 8))
	if err != nil {
		hlog.Error(err)
	}
	registryInfo := &registry.Info{Addr: ip, ServiceName: "prometheus", Weight: 1}
	err = r.Register(registryInfo)
	if err != nil {
		hlog.Error(err)
	}

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(fmt.Sprintf(":%d", "8383"), nil) //nolint:errcheck
	return func(ctx context.Context) {
		r.Deregister(registryInfo) //nolint:errcheck
	}
}
