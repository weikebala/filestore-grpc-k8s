#!/bin/bash
docker build -t weikebala/filestore_grpc_k8s_nginx -f ./Dockerfile .
docker push weikebala/filestore_grpc_k8s_nginx:latest
