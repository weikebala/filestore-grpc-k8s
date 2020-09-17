#!/bin/bash

#ROOT_DIR=$GOPATH/filestore-grpc-k8s
ROOT_DIR=/project/go/filestore-grpc-k8s
#ROOT_DIR=/data/imooc/src/filestore-grpc-k8s
services="
dbproxy
upload
download
account
apigw
"

# 编译service可执行文件
build_service() {
    rm -f ${ROOT_DIR}/deploy/bin/$1
    go build -o deploy/bin/$1 service/$1/main.go
    echo -e "\033[32m编译完成: \033[0m ${ROOT_DIR}/deploy/bin/$1"
}

# 打包镜像
build_image() {
    # 替换(hub.fileserver.com/filestore/$service), 自定义镜像名即可
    docker build -t weikebala/filestore_grpc_k8s_$1 -f ./service/$1/Dockerfile .
    echo -e "\033[32m镜像打包完成: \033[0m hub.fileserver.com/filestore/$1\n"
}

# pull镜像
pull_image() {
    # 替换(hub.fileserver.com/filestore/$service), 自定义镜像名即可
    docker push weikebala/filestore_grpc_k8s_$1:latest
    echo -e "\033[32m$1镜像pull完成: \033[0m \n"
}


# 切换到工程根目录
cd ${ROOT_DIR}

# 打包静态资源
mkdir ${ROOT_DIR}/assets -p && go-bindata -nometadata -pkg assets -o ${ROOT_DIR}/assets/asset.go static/...

# 执行编译service
mkdir -p ${ROOT_DIR}/deploy/bin && rm -f ${ROOT_DIR}/deploy/bin/*
for service in $services
do
    build_service $service
done

echo -e "\033[32m编译完毕, 开始构建docker镜像... \033[0m"

# 打包微服务镜像
cd ${ROOT_DIR}/deploy/
for service in $services
do
    build_image $service
done

echo -e "\033[32mdocker镜像构建完毕.\033[0m"


for service in $services
do
    pull_image $service
done
echo -e "\033[32mdocker镜像pull完毕.\033[0m"