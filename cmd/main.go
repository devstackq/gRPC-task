package main

import (
	"github.com/devstackq/gRPC-task/handler/grpc"
	"github.com/devstackq/gRPC-task/handler/rest"
)

func main() {
	go rest.RunRest()
	grpc.RunGrpc()
}

//protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=.   --openapiv2_out . api.proto
