package service

import (
	"context"
	"testing"
	file "github.com/doutokk/doutok/rpc_gen/kitex_gen/file"
)

func TestUploadFile_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUploadFileService(ctx)
	// init req and assert value

	req := &file.UploadFileReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
