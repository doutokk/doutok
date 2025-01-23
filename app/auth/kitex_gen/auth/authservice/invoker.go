// Code generated by Kitex v0.9.1. DO NOT EDIT.

package authservice

import (
	auth "github.com/PengJingzhao/douyin-commerce/app/auth/kitex_gen/auth"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler auth.AuthService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
