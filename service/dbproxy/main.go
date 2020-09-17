package main

import (
	cfg "filestore-grpc-k8s/config"
	"google.golang.org/grpc"
	"log"
	"net"

	dbConn "filestore-grpc-k8s/service/dbproxy/conn"
	dbProxy "filestore-grpc-k8s/service/dbproxy/proto"
	dbRpc "filestore-grpc-k8s/service/dbproxy/rpc"
)

func startRpcService() {
	lis, err := net.Listen("tcp", cfg.DbproxyServicePort) //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer() //创建gRPC服务
	dbProxy.RegisterDBProxyServiceServer(s, &dbRpc.DBProxy{})

	// 初始化db connection
	dbConn.InitDBConn()

	// 将监听交给gRPC服务处理
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	startRpcService()
}
