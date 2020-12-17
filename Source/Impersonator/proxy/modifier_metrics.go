package proxy

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	requestHandlingTime = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "impersonator_proxy_handler_time",
		Help: "The time spent handling requests",
	})

	tenantRequestTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "impersonator_proxy_handler_tenant_requests_total",
		Help: "The total number of processed requests per tenant",
	}, []string{"method", "code", "tenant_id"})
)
