# filestore-grpc-k8s
文件上传服务-grpc-k8s
go实现api、download、upload、dbproxy、account服务
服务框架基于grpc实现rpc通信。

部署
k8s部署。
nginx做k8s集群外部访问的代理，实现方案拉起一个nginx pod，在pod级别暴露端口给外网访问，nginx做反向代理，通过nginx访问k8s集群内部的服务，集群内部通过service名称访问（k8s需部署dns）
