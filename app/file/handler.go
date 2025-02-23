package main

import (
	"context"
	file "github.com/doutokk/doutok/rpc_gen/kitex_gen/file"
	"github.com/doutokk/doutok/app/file/biz/service"
)

// FileServiceImpl implements the last service interface defined in the IDL.
type FileServiceImpl struct{}

// UploadFile implements the FileServiceImpl interface.
func (s *FileServiceImpl) UploadFile(ctx context.Context, req *file.UploadFileReq) (resp *file.UploadFileResp, err error) {
	resp, err = service.NewUploadFileService(ctx).Run(req)

	return resp, err
}
