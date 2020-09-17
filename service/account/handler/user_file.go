package handler

import (
	"context"
	"encoding/json"

	"filestore-grpc-k8s/common"
	proto "filestore-grpc-k8s/service/account/proto"
	dbcli "filestore-grpc-k8s/service/dbproxy/client"
)

// UserFiles : 获取用户文件列表
func (u *User) UserFiles(ctx context.Context, req *proto.ReqUserFile) (res *proto.RespUserFile, err error) {
	dbResp, err := dbcli.QueryUserFileMetas(req.Username, int(req.Limit))
	res = &proto.RespUserFile{}
	if err != nil || !dbResp.Suc {
		res.Code = common.StatusServerError
		return
	}

	userFiles := dbcli.ToTableUserFiles(dbResp.Data)
	data, err := json.Marshal(userFiles)
	if err != nil {
		res.Code = common.StatusServerError
		return
	}

	res.FileData = data
	return
}

// UserFiles : 用户文件重命名
func (u *User) UserFileRename(ctx context.Context, req *proto.ReqUserFileRename) (res *proto.RespUserFileRename, err error) {
	dbResp, err := dbcli.RenameFileName(req.Username, req.Filehash, req.NewFileName)
	res = &proto.RespUserFileRename{}
	if err != nil || !dbResp.Suc {
		res.Code = common.StatusServerError
		return
	}

	userFiles := dbcli.ToTableUserFiles(dbResp.Data)
	data, err := json.Marshal(userFiles)
	if err != nil {
		res.Code = common.StatusServerError
		return
	}

	res.FileData = data
	return
}
