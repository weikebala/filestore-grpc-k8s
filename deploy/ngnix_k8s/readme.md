因为虚拟机k8s是1.5.6版本，需要转换
kubectl convert -f ingress.yaml --output-version=rbac.authorization.k8s.io/v1alpha1 | kubectl apply -f -