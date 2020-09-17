package rpc

import (
	"context"
	cfg "filestore-grpc-k8s/service/download/config"
	dlProto "filestore-grpc-k8s/service/download/proto"
)

// Dwonload :download结构体
type Download struct{}

// DownloadEntry : 获取下载入口
func (u *Download) DownloadEntry(ctx context.Context, req *dlProto.ReqEntry) (res *dlProto.RespEntry, err error) {
	res = &dlProto.RespEntry{}
	res.Entry = cfg.DownloadEntry
	return
}
