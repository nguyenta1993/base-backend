package grpc

import (
	"base_service/config"
	grpcmetrics "base_service/internal/metrics/grpc"
	service "base_service/internal/service"

	pb "base_service/internal/api/grpc/proto_gen"

	grpcserver "github.com/gogovan-korea/ggx-kr-service-utils/grpc/server"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
)

type Server struct {
	*service.Service
	logger  logger.Logger
	cfg     *config.AppConfig
	metrics *grpcmetrics.GrpcMetrics
}

func NewServer(s *service.Service, logger logger.Logger, cfg *config.AppConfig, metrics *grpcmetrics.GrpcMetrics) *Server {
	return &Server{s, logger, cfg, metrics}
}

func (s *Server) Run() {
	grpcServer, grpcServerInstance := grpcserver.NewServer(
		s.logger,
		grpcserver.GrpcServerConfig(*s.cfg.GRPC),
	)
	pb.RegisterUserServiceServer(grpcServerInstance, s)
	grpcServer.Run()
}
