package client

import (
	"filestore-grpc-k8s/config"
	go_micro_service_user "filestore-grpc-k8s/service/account/proto"
	"log"

	"google.golang.org/grpc"
)

func CreateAccountClient() go_micro_service_user.UserServiceClient {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial(config.AccountServiceIngress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// 创建Waiter服务的客户端
	return go_micro_service_user.NewUserServiceClient(conn)
}
