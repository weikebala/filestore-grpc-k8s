#!/bin/bash

kubectl delete -f svc_account.yaml 
kubectl delete -f svc_apigw.yaml
kubectl delete -f svc_dbproxy.yaml 
kubectl delete -f svc_upload.yaml 
kubectl delete -f svc_download.yaml
kubectl delete -f nginx.yaml
