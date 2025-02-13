// Code generated by MockGen. DO NOT EDIT.
// Source: ../../rpc_gen/kitex_gen/product/productcatalogservice/client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	product "github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
	callopt "github.com/cloudwego/kitex/client/callopt"
	gomock "github.com/golang/mock/gomock"
)

// MockproductClient is a mock of Client interface.
type MockproductClient struct {
	ctrl     *gomock.Controller
	recorder *MockproductClientMockRecorder
}

// MockproductClientMockRecorder is the mock recorder for MockproductClient.
type MockproductClientMockRecorder struct {
	mock *MockproductClient
}

// NewMockproductClient creates a new mock instance.
func NewMockproductClient(ctrl *gomock.Controller) *MockproductClient {
	mock := &MockproductClient{ctrl: ctrl}
	mock.recorder = &MockproductClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockproductClient) EXPECT() *MockproductClientMockRecorder {
	return m.recorder
}

// GetProduct mocks base method.
func (m *MockproductClient) GetProduct(ctx context.Context, Req *product.GetProductReq, callOptions ...callopt.Option) (*product.GetProductResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, Req}
	for _, a := range callOptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProduct", varargs...)
	ret0, _ := ret[0].(*product.GetProductResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockproductClientMockRecorder) GetProduct(ctx, Req interface{}, callOptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, Req}, callOptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockproductClient)(nil).GetProduct), varargs...)
}

// ListProducts mocks base method.
func (m *MockproductClient) ListProducts(ctx context.Context, Req *product.ListProductsReq, callOptions ...callopt.Option) (*product.ListProductsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, Req}
	for _, a := range callOptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProducts", varargs...)
	ret0, _ := ret[0].(*product.ListProductsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockproductClientMockRecorder) ListProducts(ctx, Req interface{}, callOptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, Req}, callOptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockproductClient)(nil).ListProducts), varargs...)
}

// SearchProducts mocks base method.
func (m *MockproductClient) SearchProducts(ctx context.Context, Req *product.SearchProductsReq, callOptions ...callopt.Option) (*product.SearchProductsResp, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, Req}
	for _, a := range callOptions {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchProducts", varargs...)
	ret0, _ := ret[0].(*product.SearchProductsResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProducts indicates an expected call of SearchProducts.
func (mr *MockproductClientMockRecorder) SearchProducts(ctx, Req interface{}, callOptions ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, Req}, callOptions...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProducts", reflect.TypeOf((*MockproductClient)(nil).SearchProducts), varargs...)
}
