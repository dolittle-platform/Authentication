package proxy

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	totalRequestsServed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "impersonator_proxy_handler_requests_total",
		Help: "The total number of processed requests",
	}, []string{"method", "code"})
)
