package service

import (
	"context"
	"github.com/doutokk/doutok/common/utils"
	file "github.com/doutokk/doutok/rpc_gen/kitex_gen/file"
)

type FrontendUploadFileService struct {
	ctx context.Context
}

// NewFrontendUploadFileService new FrontendUploadFileService
func NewFrontendUploadFileService(ctx context.Context) *FrontendUploadFileService {
	return &FrontendUploadFileService{ctx: ctx}
}

// Run create note info
func (s *FrontendUploadFileService) Run(req *file.FrontendUploadFileReq) (resp *file.FrontendUploadFileResp, err error) {
	// Finish your business logic.

	u := NewUploadFileService(s.ctx)
	userId := utils.GetUserId(&s.ctx)
	resp1, err := u.Run(&file.UploadFileReq{
		UserId:   uint64(userId),
		FileName: req.FileName,
	})
	if err != nil {
		return
	}
	resp = &file.FrontendUploadFileResp{
		Key:                  resp1.Key,
		Host:                 resp1.Host,
		Policy:               resp1.Policy,
		SecurityToken:        resp1.SecurityToken,
		Signature:            resp1.Signature,
		XOssCredential:       resp1.XOssCredential,
		XOssDate:             resp1.XOssDate,
		XOssSignatureVersion: resp1.XOssSignatureVersion,
	}
	return
}
