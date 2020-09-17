package rpc

import (
	"context"
	cfg "filestore-grpc-k8s/service/upload/config"
	upProto "filestore-grpc-k8s/service/upload/proto"
)

// Upload : upload结构体
type Upload struct{}

// UploadEntry : 获取上传入口
func (u *Upload) UploadEntry(ctx context.Context, req *upProto.ReqEntry) (res *upProto.RespEntry, err error) {
	res = &upProto.RespEntry{}
	res.Entry = cfg.UploadEntry
	return
}
