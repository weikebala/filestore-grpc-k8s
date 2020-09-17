package main

import (
	config "filestore-grpc-k8s/config"
	"filestore-grpc-k8s/service/account/handler"
	accountProto "filestore-grpc-k8s/service/account/proto"
	dbproxy "filestore-grpc-k8s/service/dbproxy/client"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.AccountServicePort) //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer() //创建gRPC服务

	accountProto.RegisterUserServiceServer(s, &handler.User{})

	// 初始化dbproxy client
	dbproxy.Init()

	// 将监听交给gRPC服务处理
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
