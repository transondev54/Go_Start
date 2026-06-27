# Lesson 4: Observability & Monitoring

## 📖 Nội dung bài học

1. Observability là gì? (Logs, Metrics, Traces)
2. Structured Logging với `zerolog`
3. Metrics với `prometheus/client_golang`
4. Distributed Tracing cơ bản với OpenTelemetry
5. Health checks & Readiness probes
6. Practical setup cho production service

---

## 1️⃣ OBSERVABILITY LÀ GÌ?

Observability là khả năng hiểu **trạng thái bên trong** hệ thống chỉ qua **output bên ngoài**.

### 3 Pillars of Observability

```
┌─────────────────────────────────────────────────────┐
│                   OBSERVABILITY                     │
│                                                     │
│   📋 LOGS        📊 METRICS       🔍 TRACES         │
│                                                     │
│  "Điều gì đã    "Hệ thống đang   "Request này      │
│   xảy ra?"       hoạt động        đi qua đâu?"     │
│                  tốt không?"                        │
│                                                     │
│  Structured     Time-series      Distributed       │
│  JSON logs      counters,        request flow      │
│                 histograms                          │
└─────────────────────────────────────────────────────┘
```

### Cài đặt dependencies

```bash
go get github.com/rs/zerolog
go get github.com/prometheus/client_golang/prometheus
go get go.opentelemetry.io/otel
go get go.opentelemetry.io/otel/trace
```

---

## 2️⃣ STRUCTURED LOGGING VỚI ZEROLOG

### Tại sao structured logging?

```go
// ❌ Unstructured log (khó parse, khó search)
log.Printf("User %s logged in from %s at %s", userID, ip, time.Now())

// ✅ Structured log (JSON, dễ query với Elasticsearch/Loki)
log.Info().
    Str("user_id", userID).
    Str("ip", ip).
    Str("event", "login").
    Msg("user logged in")
// Output: {"level":"info","user_id":"u123","ip":"1.2.3.4","event":"login","message":"user logged in","time":"..."}
```

### Setup zerolog

```go
package logger

import (
    "os"
    "time"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func Setup(env string) {
    zerolog.TimeFieldFormat = time.RFC3339

    if env == "development" {
        // Pretty print for dev
        log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
    } else {
        // JSON for production
        log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
    }

    zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

// Logger with service context
func WithService(service string) zerolog.Logger {
    return log.With().Str("service", service).Logger()
}
```

### Sử dụng zerolog

```go
// Basic logging
log.Info().Msg("server started")
log.Warn().Str("path", path).Msg("route not found")
log.Error().Err(err).Str("order_id", id).Msg("failed to process order")

// Structured context
logger := log.With().
    Str("request_id", reqID).
    Str("user_id", userID).
    Logger()

logger.Info().Msg("processing request")
logger.Debug().Interface("payload", body).Msg("request body")

// Fatal (logs then calls os.Exit(1))
log.Fatal().Err(err).Msg("cannot connect to database")
```

### HTTP Logging Middleware

```go
func LoggingMiddleware(logger zerolog.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            reqID := r.Header.Get("X-Request-ID")
            if reqID == "" {
                reqID = generateRequestID()
            }

            // Inject logger into context
            ctx := logger.With().Str("request_id", reqID).Logger().WithContext(r.Context())
            r = r.WithContext(ctx)

            // Wrap ResponseWriter to capture status code
            rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
            next.ServeHTTP(rw, r)

            zerolog.Ctx(r.Context()).Info().
                Str("method", r.Method).
                Str("path", r.URL.Path).
                Int("status", rw.status).
                Dur("duration", time.Since(start)).
                Msg("request completed")
        })
    }
}

type responseWriter struct {
    http.ResponseWriter
    status int
}

func (rw *responseWriter) WriteHeader(status int) {
    rw.status = status
    rw.ResponseWriter.WriteHeader(status)
}
```

---

## 3️⃣ METRICS VỚI PROMETHEUS

### Các loại metrics

```go
package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
    // Counter: chỉ tăng (requests, errors)
    RequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "path", "status"},
    )

    // Histogram: phân phối (latency, size)
    RequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "HTTP request duration in seconds",
            Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5},
        },
        []string{"method", "path"},
    )

    // Gauge: có thể tăng/giảm (active connections, queue size)
    ActiveConnections = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "active_connections",
        Help: "Number of active connections",
    })

    // Custom business metric
    OrdersProcessed = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "orders_processed_total",
            Help: "Total orders processed",
        },
        []string{"status"}, // "success", "failed", "cancelled"
    )
)

func Register() {
    prometheus.MustRegister(
        RequestsTotal,
        RequestDuration,
        ActiveConnections,
        OrdersProcessed,
    )
}
```

### Metrics middleware

```go
func MetricsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

        metrics.ActiveConnections.Inc()
        defer metrics.ActiveConnections.Dec()

        next.ServeHTTP(rw, r)

        duration := time.Since(start).Seconds()
        status := strconv.Itoa(rw.status)

        metrics.RequestsTotal.WithLabelValues(r.Method, r.URL.Path, status).Inc()
        metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
    })
}

// Expose /metrics endpoint
import "github.com/prometheus/client_golang/prometheus/promhttp"

mux.Handle("/metrics", promhttp.Handler())
```

---

## 4️⃣ DISTRIBUTED TRACING (OPENTELEMETRY)

### Setup OpenTelemetry

```go
package telemetry

import (
    "context"

    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
    "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
    "go.opentelemetry.io/otel/sdk/resource"
)

func SetupTracer(serviceName string) (*trace.TracerProvider, error) {
    exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
    if err != nil {
        return nil, err
    }

    tp := trace.NewTracerProvider(
        trace.WithBatcher(exporter),
        trace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceName(serviceName),
        )),
    )

    otel.SetTracerProvider(tp)
    return tp, nil
}
```

### Sử dụng tracing

```go
var tracer = otel.Tracer("order-service")

func (s *OrderService) ProcessOrder(ctx context.Context, order Order) error {
    ctx, span := tracer.Start(ctx, "ProcessOrder")
    defer span.End()

    // Add attributes to span
    span.SetAttributes(
        attribute.String("order.id", order.ID),
        attribute.Float64("order.total", order.Total),
    )

    // Child span
    ctx, childSpan := tracer.Start(ctx, "ValidateOrder")
    if err := s.validate(ctx, order); err != nil {
        childSpan.RecordError(err)
        childSpan.SetStatus(codes.Error, err.Error())
        childSpan.End()
        return err
    }
    childSpan.End()

    return s.repo.Save(ctx, order)
}
```

---

## 5️⃣ HEALTH CHECKS

```go
type HealthStatus struct {
    Status   string            `json:"status"` // "ok" hoặc "degraded"
    Checks   map[string]string `json:"checks"`
    Uptime   string            `json:"uptime"`
    Version  string            `json:"version"`
}

type HealthChecker struct {
    db      *sql.DB
    startAt time.Time
    version string
}

func (h *HealthChecker) Handler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
    defer cancel()

    checks := make(map[string]string)
    overall := "ok"

    // Check DB
    if err := h.db.PingContext(ctx); err != nil {
        checks["database"] = "unhealthy: " + err.Error()
        overall = "degraded"
    } else {
        checks["database"] = "healthy"
    }

    status := HealthStatus{
        Status:  overall,
        Checks:  checks,
        Uptime:  time.Since(h.startAt).String(),
        Version: h.version,
    }

    w.Header().Set("Content-Type", "application/json")
    if overall != "ok" {
        w.WriteHeader(http.StatusServiceUnavailable)
    }
    json.NewEncoder(w).Encode(status)
}
```

---

## 🧠 QUIZ - 5 CÂU HỎI

1. Sự khác nhau giữa Counter, Gauge và Histogram trong Prometheus?
2. Tại sao structured logging (JSON) tốt hơn plain text logging?
3. Distributed trace và log thông thường khác nhau như thế nào?
4. Health check `/health` và readiness probe `/ready` khác nhau gì?
5. Nếu một span bị lỗi, bạn cần làm gì trong OpenTelemetry?

---

## 📌 KEY TAKEAWAYS

- Logs, Metrics, Traces = 3 pillars of observability
- Structured JSON logs → dễ query và index
- Prometheus counters/histograms/gauges cho từng loại metric
- OpenTelemetry là standard, exporter có thể swap (Jaeger, Zipkin, ...)
- Health check endpoint là must-have cho mọi production service
