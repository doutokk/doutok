package service

import (
	"context"
	file "github.com/doutokk/doutok/rpc_gen/kitex_gen/file"
	"testing"
)

func TestFrontendUploadFile_Run(t *testing.T) {
	ctx := context.Background()
	s := NewFrontendUploadFileService(ctx)
	// init req and assert value

	req := &file.FrontendUploadFileReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
