#!/bin/bash

kubectl apply -f svc_account.yaml
kubectl apply -f svc_apigw.yaml
kubectl apply -f svc_dbproxy.yaml 
kubectl apply -f svc_upload.yaml 
kubectl apply -f svc_download.yaml
kubectl apply -f nginx.yaml

