// Code generated by Kitex v0.9.1. DO NOT EDIT.

package productcatalogservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	product "github.com/doutokk/doutok/app/product/kitex_gen/product"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"ListProducts": kitex.NewMethodInfo(
		listProductsHandler,
		newListProductsArgs,
		newListProductsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetProduct": kitex.NewMethodInfo(
		getProductHandler,
		newGetProductArgs,
		newGetProductResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"SearchProducts": kitex.NewMethodInfo(
		searchProductsHandler,
		newSearchProductsArgs,
		newSearchProductsResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	productCatalogServiceServiceInfo                = NewServiceInfo()
	productCatalogServiceServiceInfoForClient       = NewServiceInfoForClient()
	productCatalogServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return productCatalogServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "ProductCatalogService"
	handlerType := (*product.ProductCatalogService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "product",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func listProductsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.ListProductsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).ListProducts(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListProductsArgs:
		success, err := handler.(product.ProductCatalogService).ListProducts(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListProductsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListProductsArgs() interface{} {
	return &ListProductsArgs{}
}

func newListProductsResult() interface{} {
	return &ListProductsResult{}
}

type ListProductsArgs struct {
	Req *product.ListProductsReq
}

func (p *ListProductsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.ListProductsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListProductsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListProductsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListProductsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListProductsArgs) Unmarshal(in []byte) error {
	msg := new(product.ListProductsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListProductsArgs_Req_DEFAULT *product.ListProductsReq

func (p *ListProductsArgs) GetReq() *product.ListProductsReq {
	if !p.IsSetReq() {
		return ListProductsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListProductsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListProductsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListProductsResult struct {
	Success *product.ListProductsResp
}

var ListProductsResult_Success_DEFAULT *product.ListProductsResp

func (p *ListProductsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.ListProductsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListProductsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListProductsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListProductsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListProductsResult) Unmarshal(in []byte) error {
	msg := new(product.ListProductsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListProductsResult) GetSuccess() *product.ListProductsResp {
	if !p.IsSetSuccess() {
		return ListProductsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListProductsResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.ListProductsResp)
}

func (p *ListProductsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListProductsResult) GetResult() interface{} {
	return p.Success
}

func getProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.GetProductReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).GetProduct(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetProductArgs:
		success, err := handler.(product.ProductCatalogService).GetProduct(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetProductResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetProductArgs() interface{} {
	return &GetProductArgs{}
}

func newGetProductResult() interface{} {
	return &GetProductResult{}
}

type GetProductArgs struct {
	Req *product.GetProductReq
}

func (p *GetProductArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.GetProductReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetProductArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetProductArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetProductArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetProductArgs) Unmarshal(in []byte) error {
	msg := new(product.GetProductReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetProductArgs_Req_DEFAULT *product.GetProductReq

func (p *GetProductArgs) GetReq() *product.GetProductReq {
	if !p.IsSetReq() {
		return GetProductArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetProductArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetProductArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetProductResult struct {
	Success *product.GetProductResp
}

var GetProductResult_Success_DEFAULT *product.GetProductResp

func (p *GetProductResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.GetProductResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetProductResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetProductResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetProductResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetProductResult) Unmarshal(in []byte) error {
	msg := new(product.GetProductResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetProductResult) GetSuccess() *product.GetProductResp {
	if !p.IsSetSuccess() {
		return GetProductResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetProductResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.GetProductResp)
}

func (p *GetProductResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetProductResult) GetResult() interface{} {
	return p.Success
}

func searchProductsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(product.SearchProductsReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(product.ProductCatalogService).SearchProducts(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *SearchProductsArgs:
		success, err := handler.(product.ProductCatalogService).SearchProducts(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SearchProductsResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newSearchProductsArgs() interface{} {
	return &SearchProductsArgs{}
}

func newSearchProductsResult() interface{} {
	return &SearchProductsResult{}
}

type SearchProductsArgs struct {
	Req *product.SearchProductsReq
}

func (p *SearchProductsArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(product.SearchProductsReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *SearchProductsArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *SearchProductsArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *SearchProductsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *SearchProductsArgs) Unmarshal(in []byte) error {
	msg := new(product.SearchProductsReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SearchProductsArgs_Req_DEFAULT *product.SearchProductsReq

func (p *SearchProductsArgs) GetReq() *product.SearchProductsReq {
	if !p.IsSetReq() {
		return SearchProductsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SearchProductsArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *SearchProductsArgs) GetFirstArgument() interface{} {
	return p.Req
}

type SearchProductsResult struct {
	Success *product.SearchProductsResp
}

var SearchProductsResult_Success_DEFAULT *product.SearchProductsResp

func (p *SearchProductsResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(product.SearchProductsResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *SearchProductsResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *SearchProductsResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *SearchProductsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *SearchProductsResult) Unmarshal(in []byte) error {
	msg := new(product.SearchProductsResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SearchProductsResult) GetSuccess() *product.SearchProductsResp {
	if !p.IsSetSuccess() {
		return SearchProductsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SearchProductsResult) SetSuccess(x interface{}) {
	p.Success = x.(*product.SearchProductsResp)
}

func (p *SearchProductsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *SearchProductsResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ListProducts(ctx context.Context, Req *product.ListProductsReq) (r *product.ListProductsResp, err error) {
	var _args ListProductsArgs
	_args.Req = Req
	var _result ListProductsResult
	if err = p.c.Call(ctx, "ListProducts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetProduct(ctx context.Context, Req *product.GetProductReq) (r *product.GetProductResp, err error) {
	var _args GetProductArgs
	_args.Req = Req
	var _result GetProductResult
	if err = p.c.Call(ctx, "GetProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SearchProducts(ctx context.Context, Req *product.SearchProductsReq) (r *product.SearchProductsResp, err error) {
	var _args SearchProductsArgs
	_args.Req = Req
	var _result SearchProductsResult
	if err = p.c.Call(ctx, "SearchProducts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
