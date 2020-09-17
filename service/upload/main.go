package main

import (
	"filestore-grpc-k8s/config"
	uploadProto "filestore-grpc-k8s/service/upload/proto"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	dbproxy "filestore-grpc-k8s/service/dbproxy/client"
	cfg "filestore-grpc-k8s/service/upload/config"
	"filestore-grpc-k8s/service/upload/route"
	upRpc "filestore-grpc-k8s/service/upload/rpc"
)

func startRPCService() {
	lis, err := net.Listen("tcp", config.UploadServicePort) //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer() //创建gRPC服务

	uploadProto.RegisterUploadServiceServer(s, &upRpc.Upload{})

	// 初始化dbproxy client
	dbproxy.Init()

	// 将监听交给gRPC服务处理
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startAPIService() {
	router := route.Router()
	router.Run(cfg.UploadServiceHost)
}

func main() {
	os.MkdirAll(config.TempLocalRootDir, 0777)
	os.MkdirAll(config.TempPartRootDir, 0777)

	// api 服务
	go startAPIService()

	// rpc 服务
	startRPCService()
}
