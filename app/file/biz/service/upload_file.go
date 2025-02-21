package service

import (
	"context"
	file "github.com/doutokk/doutok/rpc_gen/kitex_gen/file"
)

type UploadFileService struct {
	ctx context.Context
}

// NewUploadFileService new UploadFileService
func NewUploadFileService(ctx context.Context) *UploadFileService {
	return &UploadFileService{ctx: ctx}
}

// Run create note info
func (s *UploadFileService) Run(req *file.UploadFileReq) (resp *file.UploadFileResp, err error) {
	// Finish your business logic.

	return
}
