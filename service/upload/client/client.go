package client

import (
	"filestore-grpc-k8s/config"
	upProto "filestore-grpc-k8s/service/upload/proto"
	"log"

	"google.golang.org/grpc"
)

func CreateUploadClient() upProto.UploadServiceClient {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial(config.UploadServiceIngress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// 创建Waiter服务的客户端
	return upProto.NewUploadServiceClient(conn)
}
