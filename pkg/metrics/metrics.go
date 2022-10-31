package metrics

import (
	"base_service/config"

	"github.com/gin-gonic/gin"
	"github.com/gogovan-korea/ggx-kr-service-utils/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func RunMetrics(logger logger.Logger, cfg *config.AppConfig) func() {
	return func() {
		gin.SetMode(gin.ReleaseMode)
		metricsServer := gin.New()

		metricsServer.GET(cfg.Metrics.PrometheusPath, prometheusHandler())
		logger.Info("Metrics server is running on port", zap.String("Metrics port", cfg.Metrics.PrometheusPort))
		if err := metricsServer.Run(cfg.Metrics.PrometheusPort); err != nil {
			// If service uses both of http & grpc, it probaly happens error here(already bind the same port)
			// It's still good to go
			logger.Error("metricsServer.Run", zap.Error(err))
		}
	}
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
