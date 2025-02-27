package cart

import (
	"context"
	cart "github.com/doutokk/doutok/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	
)

type RPCClient interface {
	KitexClient() cartservice.Client
	Service() string
	AddItem(ctx context.Context, Req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error)
	GetCart(ctx context.Context, Req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error)
	EmptyCart(ctx context.Context, Req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error)
	EditCart(ctx context.Context, Req *cart.EditCartReq, callOptions ...callopt.Option) (r *cart.EditCartResp, err error)
	FrontendGetCart(ctx context.Context, Req *cart.FrontendGetCartReq, callOptions ...callopt.Option) (r *cart.FrontendGetCartResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := cartservice.NewClient(dstService, opts...)
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
	kitexClient cartservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() cartservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AddItem(ctx context.Context, Req *cart.AddItemReq, callOptions ...callopt.Option) (r *cart.AddItemResp, err error) {
	return c.kitexClient.AddItem(ctx, Req, callOptions...)
}

func (c *clientImpl) GetCart(ctx context.Context, Req *cart.GetCartReq, callOptions ...callopt.Option) (r *cart.GetCartResp, err error) {
	return c.kitexClient.GetCart(ctx, Req, callOptions...)
}

func (c *clientImpl) EmptyCart(ctx context.Context, Req *cart.EmptyCartReq, callOptions ...callopt.Option) (r *cart.EmptyCartResp, err error) {
	return c.kitexClient.EmptyCart(ctx, Req, callOptions...)
}

func (c *clientImpl) EditCart(ctx context.Context, Req *cart.EditCartReq, callOptions ...callopt.Option) (r *cart.EditCartResp, err error) {
	return c.kitexClient.EditCart(ctx, Req, callOptions...)
}

func (c *clientImpl) FrontendGetCart(ctx context.Context, Req *cart.FrontendGetCartReq, callOptions ...callopt.Option) (r *cart.FrontendGetCartResp, err error) {
	return c.kitexClient.FrontendGetCart(ctx, Req, callOptions...)
}
