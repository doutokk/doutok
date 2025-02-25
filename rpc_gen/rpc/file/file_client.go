package file

import (
	"context"
	file "github.com/doutokk/doutok/rpc_gen/kitex_gen/file"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	
)

type RPCClient interface {
	KitexClient() fileservice.Client
	Service() string
	UploadFile(ctx context.Context, Req *file.UploadFileReq, callOptions ...callopt.Option) (r *file.UploadFileResp, err error)
	FrontendUploadFile(ctx context.Context, Req *file.FrontendUploadFileReq, callOptions ...callopt.Option) (r *file.FrontendUploadFileResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := fileservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient fileservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() fileservice.Client {
	return c.kitexClient
}

func (c *clientImpl) UploadFile(ctx context.Context, Req *file.UploadFileReq, callOptions ...callopt.Option) (r *file.UploadFileResp, err error) {
	return c.kitexClient.UploadFile(ctx, Req, callOptions...)
}

func (c *clientImpl) FrontendUploadFile(ctx context.Context, Req *file.FrontendUploadFileReq, callOptions ...callopt.Option) (r *file.FrontendUploadFileResp, err error) {
	return c.kitexClient.FrontendUploadFile(ctx, Req, callOptions...)
}
