package service

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/doutokk/doutok/app/file/biz/aliyun"
	"github.com/doutokk/doutok/app/file/biz/dal/model"
	"github.com/doutokk/doutok/app/file/biz/dal/query"
	file "github.com/doutokk/doutok/rpc_gen/kitex_gen/file"
	"github.com/google/uuid"
	"strings"
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

	split := strings.Split(req.FileName, ".")
	if len(split) < 2 {
		return nil, errors.New("不允许上传无扩展名文件")
	}
	fileExt := split[1]
	f := &model.File{
		UserId:         req.UserId,
		FileOriginName: req.FileName,
		Key:            uuid.NewString() + "." + fileExt,
		Usage:          "default",
	}
	err = query.Q.File.Create(f)
	if err != nil {
		klog.Errorf("create file record: %v", err)
		return nil, err
	}

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
