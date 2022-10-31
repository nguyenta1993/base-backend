package http

import (
	"base_service/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type HttpMetrics struct {
	SuccessHttpRequests prometheus.Counter
	ErrorHttpRequests   prometheus.Counter
	GetUserHttpRequests prometheus.Counter
}

func NewHttpMetrics(cfg *config.AppConfig) *HttpMetrics {
	return &HttpMetrics{
		SuccessHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: cfg.ServiceName,
			Name:      "success_http_requests_total",
			Help:      "The total number of success http requests",
		}),
		ErrorHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: cfg.ServiceName,
			Name:      "error_http_requests_total",
			Help:      "The total number of error http requests",
		}),
		GetUserHttpRequests: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: cfg.ServiceName,
			Name:      "get_user_http_requests_total",
			Help:      "The total number of get user http requests",
		}),
	}
}
