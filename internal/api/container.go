package api

import (
	"base_service/internal/api/grpc"
	"base_service/internal/api/http"
	"base_service/internal/api/kafka"
)

type ApiContainer struct {
	HttpServer *http.Server
	GrpcServer *grpc.Server
	Consumer   *kafka.Consumer
}

func NewApiContainer(http *http.Server, grpc *grpc.Server, consumer *kafka.Consumer) *ApiContainer {
	return &ApiContainer{HttpServer: http, GrpcServer: grpc, Consumer: consumer}
}
