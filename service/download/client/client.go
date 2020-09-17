package client

import (
	"filestore-grpc-k8s/config"
	go_micro_service_download "filestore-grpc-k8s/service/download/proto"
	"log"

	"google.golang.org/grpc"
)

func CreateDownloadClient() go_micro_service_download.DownloadServiceClient {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial(config.DownloadServiceIngress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// 创建Waiter服务的客户端
	return go_micro_service_download.NewDownloadServiceClient(conn)
}
