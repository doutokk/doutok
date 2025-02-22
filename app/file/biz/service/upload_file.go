package service

import (
	"context"
	"github.com/doutokk/doutok/app/file/biz/aliyun"
	file "github.com/doutokk/doutok/rpc_gen/kitex_gen/file"
	"time"
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
	policyResp := aliyun.GetPolicy(aliyun.GetPolicyReq{
		FileName:   req.FileName,
		ExpireTime: time.Now().Add(30 * time.Minute),
	})

	resp = &file.UploadFileResp{
		Key:                  policyResp.Key,
		Host:                 policyResp.Host,
		Policy:               policyResp.Policy,
		SecurityToken:        policyResp.SecurityToken,
		Signature:            policyResp.Signature,
		XOssCredential:       policyResp.SignatureCredential,
		XOssDate:             policyResp.Date,
		XOssSignatureVersion: policyResp.SignatureVersion,
	}

	return
}
