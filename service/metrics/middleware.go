package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		// tính toán metrics
		status := strconv.Itoa(c.Writer.Status())
		path := c.FullPath()
		method := c.Request.Method
		duration := time.Since(start).Seconds()

		HttpRequestsTotal.WithLabelValues(method, path, status).Inc()
		HttpRequestDuration.WithLabelValues(method, path).Observe(duration)
	}
}
