package file

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	file "github.com/doutokk/doutok/rpc_gen/kitex_gen/file"
)

func UploadFile(ctx context.Context, req *file.UploadFileReq, callOptions ...callopt.Option) (resp *file.UploadFileResp, err error) {
	resp, err = defaultClient.UploadFile(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UploadFile call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func FrontendUploadFile(ctx context.Context, req *file.FrontendUploadFileReq, callOptions ...callopt.Option) (resp *file.FrontendUploadFileResp, err error) {
	resp, err = defaultClient.FrontendUploadFile(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "FrontendUploadFile call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
