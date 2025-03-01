package serversuite

import (
	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/doutokk/doutok/common/mtl"

	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	registryconsul "github.com/kitex-contrib/registry-consul"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func LoggingMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		// 获取 RPC 方法信息
		ri := rpcinfo.GetRPCInfo(ctx)
		method := ri.To().Method()

		// 序列化请求
		reqJSON, err := json.MarshalIndent(req, "", "  ")
		if err != nil {
			fmt.Printf("[%s] Marshal request error: %v\n", method, err)
		} else {
			fmt.Printf("[%s] REQUEST:\n%s\n", method, string(reqJSON))
		}

		// 执行后续调用
		err = next(ctx, req, resp)

		// 序列化响应
		respJSON, marshalErr := json.MarshalIndent(resp, "", "  ")
		if marshalErr != nil {
			fmt.Printf("[%s] Marshal response error: %v\n", method, marshalErr)
		} else {
			if err != nil {
				fmt.Printf("[%s] RESPONSE ERROR: %v\n", method, err)
			} else {
				fmt.Printf("[%s] RESPONSE:\n%s\n", method, string(respJSON))
			}
		}

		return err
	}
}
func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithMiddleware(LoggingMiddleware),
	}

	// 注册到 consul
	r, err := registryconsul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))

	_ = provider.NewOpenTelemetryProvider(provider.WithSdkTracerProvider(mtl.TracerProvider), provider.WithEnableMetrics(false))

	opts = append(opts,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithSuite(tracing.NewServerSuite()),
		// 配置服务上传指标
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
	)

	return opts
}
