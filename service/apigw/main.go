package main

import (
	"filestore-grpc-k8s/service/apigw/route"
)

func main() {
	r := route.Router()
	r.Run(":8080")
}
