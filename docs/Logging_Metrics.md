# Logging và Metrics trong dự án Go

## 1. Logging

### Khái niệm
Logging là quá trình ghi lại các thông tin, sự kiện, lỗi hoặc trạng thái của ứng dụng trong quá trình chạy. Log giúp bạn:
- Theo dõi hoạt động của hệ thống.
- Phát hiện và xử lý lỗi.
- Phân tích hiệu năng và hành vi ứng dụng.

### Cách khởi tạo
Bạn có thể sử dụng package `log` mặc định của Go hoặc các thư viện nâng cao như `zap`, `logrus`, hoặc `zerolog` để log ra JSON.

#### Sử dụng zerolog (logging JSON)
Cài đặt zerolog:
```bash
go get github.com/rs/zerolog
```
Khởi tạo logger với zerolog:
```go
import (
	"os"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}) // log ra console (có thể bỏ để log ra JSON thuần)
}
```
Nếu muốn log ra JSON thuần:
```go
log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
```

### Cách sử dụng
Ghi log thông tin:
```go
log.Info().Msg("Ứng dụng đã khởi động")
```
Ghi log lỗi:
```go
log.Error().Err(err).Msg("Lỗi kết nối DB")
```
Ghi log với field:
```go
log.Info().Str("module", "main").Msg("Khởi động thành công")
```

#### Lợi ích logging JSON
- Dễ tích hợp với các hệ thống log tập trung (ELK, Grafana, Loki...)
- Dễ dàng phân tích, lọc, tìm kiếm log theo field
- Chuẩn hóa format log cho microservices

### So sánh
- `log` mặc định: đơn giản, log ra text, không có field.
- `zerolog`: log ra JSON, hỗ trợ field, hiệu năng cao, phù hợp cho hệ thống lớn.

## 2. Metrics

### Khái niệm
Metrics là các số liệu định lượng về hoạt động của hệ thống, ví dụ: số lượng request, thời gian xử lý, số lần cache hit/miss... Metrics giúp bạn:
- Giám sát hiệu năng hệ thống.
- Phân tích bottleneck.
- Đưa ra quyết định tối ưu hóa.

### Cách khởi tạo
Dự án sử dụng Prometheus client cho Go (`github.com/prometheus/client_golang/prometheus`). Khởi tạo các metrics trong file `service/metrics/metrics.go`:

```go
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
```

Khởi tạo metrics:
```go
func InitMetrics() {
	prometheus.MustRegister(
		HttpRequestsTotal,
		HttpRequestDuration,
		DBQueryDuration,
		CacheHitTotal,
		CacheMissTotal)
}
```

### Cách sử dụng

- Đếm số lượng request:
```go
metrics.HttpRequestsTotal.WithLabelValues(path, method, status).Inc()
```

- Đo thời gian xử lý request:
```go
timer := prometheus.NewTimer(metrics.HttpRequestDuration.WithLabelValues(method, path))
defer timer.ObserveDuration()
```

- Đo thời gian query DB:
```go
timer := prometheus.NewTimer(metrics.DBQueryDuration)
defer timer.ObserveDuration()
```

- Đếm cache hit/miss:
```go
metrics.CacheHitTotal.Inc()   // Khi cache hit
metrics.CacheMissTotal.Inc()  // Khi cache miss
```

### Export metrics cho Prometheus
Thường bạn sẽ expose endpoint `/metrics` để Prometheus scrape:
```go
import "github.com/prometheus/client_golang/prometheus/promhttp"
http.Handle("/metrics", promhttp.Handler())
```

## 3. Tổng kết

- Logging giúp ghi lại hoạt động và lỗi của hệ thống.
- Metrics giúp giám sát hiệu năng và hành vi hệ thống.
- Sử dụng đúng cách sẽ giúp bạn vận hành, giám sát và tối ưu hóa hệ thống hiệu quả.
