package main

import (
	"filestore-grpc-k8s/config"
	dbproxy "filestore-grpc-k8s/service/dbproxy/client"
	cfg "filestore-grpc-k8s/service/download/config"
	dlProto "filestore-grpc-k8s/service/download/proto"
	"filestore-grpc-k8s/service/download/route"
	dlRpc "filestore-grpc-k8s/service/download/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func startRPCService() {
	lis, err := net.Listen("tcp", config.DownloadServicePort)
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer() //创建gRPC服务

	dlProto.RegisterDownloadServiceServer(s, &dlRpc.Download{})

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
	router.Run(cfg.DownloadServiceHost)
}

func main() {
	// api 服务
	go startAPIService()

	// rpc 服务
	startRPCService()
}
