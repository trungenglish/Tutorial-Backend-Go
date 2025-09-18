package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"})

	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of HTTP request duration",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"})

	DBQueryDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "db_query_duration_seconds",
			Help:    "Histogram of DB query duration",
			Buckets: prometheus.DefBuckets},
	)

	CacheHitTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "cache_hit_total",
			Help: "Total number of cache hits"})

	CacheMissTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "cache_miss_total",
			Help: "Total number of cache misses"})
)

func InitMetrics() {
	prometheus.MustRegister(
		HttpRequestsTotal,
		HttpRequestDuration,
		DBQueryDuration,
		CacheHitTotal,
		CacheMissTotal)
}
