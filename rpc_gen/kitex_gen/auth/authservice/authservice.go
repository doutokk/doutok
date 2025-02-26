// Code generated by Kitex v0.9.1. DO NOT EDIT.

package authservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	auth "github.com/doutokk/doutok/rpc_gen/kitex_gen/auth"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"DeliverTokenByRPC": kitex.NewMethodInfo(
		deliverTokenByRPCHandler,
		newDeliverTokenByRPCArgs,
		newDeliverTokenByRPCResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"VerifyTokenByRPC": kitex.NewMethodInfo(
		verifyTokenByRPCHandler,
		newVerifyTokenByRPCArgs,
		newVerifyTokenByRPCResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"CreateUserRole": kitex.NewMethodInfo(
		createUserRoleHandler,
		newCreateUserRoleArgs,
		newCreateUserRoleResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	authServiceServiceInfo                = NewServiceInfo()
	authServiceServiceInfoForClient       = NewServiceInfoForClient()
	authServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return authServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return authServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return authServiceServiceInfoForClient
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
	serviceName := "AuthService"
	handlerType := (*auth.AuthService)(nil)
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
		"PackageName": "auth",
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

func deliverTokenByRPCHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.DeliverTokenReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).DeliverTokenByRPC(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeliverTokenByRPCArgs:
		success, err := handler.(auth.AuthService).DeliverTokenByRPC(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeliverTokenByRPCResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeliverTokenByRPCArgs() interface{} {
	return &DeliverTokenByRPCArgs{}
}

func newDeliverTokenByRPCResult() interface{} {
	return &DeliverTokenByRPCResult{}
}

type DeliverTokenByRPCArgs struct {
	Req *auth.DeliverTokenReq
}

func (p *DeliverTokenByRPCArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.DeliverTokenReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeliverTokenByRPCArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeliverTokenByRPCArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeliverTokenByRPCArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeliverTokenByRPCArgs) Unmarshal(in []byte) error {
	msg := new(auth.DeliverTokenReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeliverTokenByRPCArgs_Req_DEFAULT *auth.DeliverTokenReq

func (p *DeliverTokenByRPCArgs) GetReq() *auth.DeliverTokenReq {
	if !p.IsSetReq() {
		return DeliverTokenByRPCArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeliverTokenByRPCArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeliverTokenByRPCArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeliverTokenByRPCResult struct {
	Success *auth.DeliveryResp
}

var DeliverTokenByRPCResult_Success_DEFAULT *auth.DeliveryResp

func (p *DeliverTokenByRPCResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.DeliveryResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeliverTokenByRPCResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeliverTokenByRPCResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeliverTokenByRPCResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeliverTokenByRPCResult) Unmarshal(in []byte) error {
	msg := new(auth.DeliveryResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeliverTokenByRPCResult) GetSuccess() *auth.DeliveryResp {
	if !p.IsSetSuccess() {
		return DeliverTokenByRPCResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeliverTokenByRPCResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.DeliveryResp)
}

func (p *DeliverTokenByRPCResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeliverTokenByRPCResult) GetResult() interface{} {
	return p.Success
}

func verifyTokenByRPCHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.VerifyTokenReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).VerifyTokenByRPC(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *VerifyTokenByRPCArgs:
		success, err := handler.(auth.AuthService).VerifyTokenByRPC(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*VerifyTokenByRPCResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newVerifyTokenByRPCArgs() interface{} {
	return &VerifyTokenByRPCArgs{}
}

func newVerifyTokenByRPCResult() interface{} {
	return &VerifyTokenByRPCResult{}
}

type VerifyTokenByRPCArgs struct {
	Req *auth.VerifyTokenReq
}

func (p *VerifyTokenByRPCArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.VerifyTokenReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *VerifyTokenByRPCArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *VerifyTokenByRPCArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *VerifyTokenByRPCArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *VerifyTokenByRPCArgs) Unmarshal(in []byte) error {
	msg := new(auth.VerifyTokenReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var VerifyTokenByRPCArgs_Req_DEFAULT *auth.VerifyTokenReq

func (p *VerifyTokenByRPCArgs) GetReq() *auth.VerifyTokenReq {
	if !p.IsSetReq() {
		return VerifyTokenByRPCArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *VerifyTokenByRPCArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *VerifyTokenByRPCArgs) GetFirstArgument() interface{} {
	return p.Req
}

type VerifyTokenByRPCResult struct {
	Success *auth.VerifyResp
}

var VerifyTokenByRPCResult_Success_DEFAULT *auth.VerifyResp

func (p *VerifyTokenByRPCResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.VerifyResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *VerifyTokenByRPCResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *VerifyTokenByRPCResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *VerifyTokenByRPCResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *VerifyTokenByRPCResult) Unmarshal(in []byte) error {
	msg := new(auth.VerifyResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *VerifyTokenByRPCResult) GetSuccess() *auth.VerifyResp {
	if !p.IsSetSuccess() {
		return VerifyTokenByRPCResult_Success_DEFAULT
	}
	return p.Success
}

func (p *VerifyTokenByRPCResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.VerifyResp)
}

func (p *VerifyTokenByRPCResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *VerifyTokenByRPCResult) GetResult() interface{} {
	return p.Success
}

func createUserRoleHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(auth.CreateUserRoleReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(auth.AuthService).CreateUserRole(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateUserRoleArgs:
		success, err := handler.(auth.AuthService).CreateUserRole(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateUserRoleResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateUserRoleArgs() interface{} {
	return &CreateUserRoleArgs{}
}

func newCreateUserRoleResult() interface{} {
	return &CreateUserRoleResult{}
}

type CreateUserRoleArgs struct {
	Req *auth.CreateUserRoleReq
}

func (p *CreateUserRoleArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(auth.CreateUserRoleReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateUserRoleArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateUserRoleArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateUserRoleArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateUserRoleArgs) Unmarshal(in []byte) error {
	msg := new(auth.CreateUserRoleReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateUserRoleArgs_Req_DEFAULT *auth.CreateUserRoleReq

func (p *CreateUserRoleArgs) GetReq() *auth.CreateUserRoleReq {
	if !p.IsSetReq() {
		return CreateUserRoleArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateUserRoleArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateUserRoleArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateUserRoleResult struct {
	Success *auth.CreateUserRoleResp
}

var CreateUserRoleResult_Success_DEFAULT *auth.CreateUserRoleResp

func (p *CreateUserRoleResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(auth.CreateUserRoleResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateUserRoleResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateUserRoleResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateUserRoleResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateUserRoleResult) Unmarshal(in []byte) error {
	msg := new(auth.CreateUserRoleResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateUserRoleResult) GetSuccess() *auth.CreateUserRoleResp {
	if !p.IsSetSuccess() {
		return CreateUserRoleResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateUserRoleResult) SetSuccess(x interface{}) {
	p.Success = x.(*auth.CreateUserRoleResp)
}

func (p *CreateUserRoleResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateUserRoleResult) GetResult() interface{} {
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

func (p *kClient) DeliverTokenByRPC(ctx context.Context, Req *auth.DeliverTokenReq) (r *auth.DeliveryResp, err error) {
	var _args DeliverTokenByRPCArgs
	_args.Req = Req
	var _result DeliverTokenByRPCResult
	if err = p.c.Call(ctx, "DeliverTokenByRPC", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) VerifyTokenByRPC(ctx context.Context, Req *auth.VerifyTokenReq) (r *auth.VerifyResp, err error) {
	var _args VerifyTokenByRPCArgs
	_args.Req = Req
	var _result VerifyTokenByRPCResult
	if err = p.c.Call(ctx, "VerifyTokenByRPC", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateUserRole(ctx context.Context, Req *auth.CreateUserRoleReq) (r *auth.CreateUserRoleResp, err error) {
	var _args CreateUserRoleArgs
	_args.Req = Req
	var _result CreateUserRoleResult
	if err = p.c.Call(ctx, "CreateUserRole", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
