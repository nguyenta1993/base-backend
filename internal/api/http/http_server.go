package http

import (
	"base_service/config"
	v1 "base_service/internal/api/http/v1"

	httpserver "github.com/gogovan-korea/ggx-kr-service-utils/http/server"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
)

type Server struct {
	logger      logger.Logger
	cfg         *config.AppConfig
	userHandler *v1.UserHandler
}

func NewServer(logger logger.Logger, cfg *config.AppConfig, userHandler *v1.UserHandler) *Server {
	return &Server{logger, cfg, userHandler}
}

func (s *Server) Run() {
	config := &httpserver.HttpServerConfig{
		Port:            s.cfg.Http.Port,
		Development:     s.cfg.Http.Development,
		ShutdownTimeout: s.cfg.Http.ShutdownTimeout,
		Resources:       s.cfg.Http.Resources,
		RateLimiting: &httpserver.RateLimitingConfig{
			RateFormat: s.cfg.Http.RateLimiting.RateFormat,
		},
	}
	httpServer, router := httpserver.NewServer(s.logger, *config)
	// In the future, if we have v2, v3..., we will add at here
	v1.MapRoutes(router, s.userHandler)
	httpServer.Run()
}
