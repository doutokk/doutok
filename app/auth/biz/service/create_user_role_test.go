package service

import (
	"context"
	auth "github.com/doutokk/doutok/rpc_gen/kitex_gen/auth"
	"testing"
)

func TestCreateUserRole_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateUserRoleService(ctx)
	// init req and assert value

	req := &auth.CreateUserRoleReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
