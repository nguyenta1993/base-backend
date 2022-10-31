package grpc

import (
	"base_service/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type GrpcMetrics struct {
	SuccessGrpcRequests prometheus.Counter
	ErrorGrpcRequests   prometheus.Counter
	GetUserGrpcRequests prometheus.Counter
	GrpcRequestsLatency *prometheus.HistogramVec
}

func NewGrpcMetrics(cfg *config.AppConfig) *GrpcMetrics {
	return &GrpcMetrics{
		SuccessGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: cfg.ServiceName,
			Name:      "success_grpc_requests_total",
			Help:      "The total number of success grpc requests",
		}),
		ErrorGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: cfg.ServiceName,
			Name:      "error_grpc_requests_total",
			Help:      "The total number of error grpc requests",
		}),
		GetUserGrpcRequests: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: cfg.ServiceName,
			Name:      "get_user_grpc_requests_total",
			Help:      "The total number of get user grpc requests",
		}),
		GrpcRequestsLatency: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: cfg.ServiceName,
			Name:      "grpc_duration_seconds",
			Help:      "Latency of grpc request in second",
			Buckets:   prometheus.LinearBuckets(0.01, 0.05, 10),
		}, []string{"method"}),
	}
}
