package main

import (
	"context"
	"douyin-commerce/auth_service/kitex_gen/auth"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// GenerateToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) GenerateToken(ctx context.Context, userId int64) (resp *auth.Token, err error) {
	// TODO: Your code here...
	return
}

// ValidateToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) ValidateToken(ctx context.Context, token string) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// RevokeToken implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RevokeToken(ctx context.Context, token string) (resp bool, err error) {
	// TODO: Your code here...
	return
}
