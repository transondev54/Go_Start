# Lesson 8: Deployment & Scaling

## 📖 Nội dung bài học

1. Build optimization
2. Cross-compilation
3. Docker containerization
4. Deployment strategies
5. Scaling patterns
6. Monitoring & logging

---

## 1️⃣ BUILD OPTIMIZATION

### Binary Size

```bash
# Standard build
go build -o app main.go
# Size: ~15MB

# Optimized build
go build -ldflags="-s -w" -o app main.go
# Size: ~12MB

# With UPX compression
upx app
# Size: ~4MB
```

### Flags Explained

```bash
# -s: disable symbol table
# -w: disable debug information
# -ldflags: linker flags

go build -ldflags="-s -w -X main.Version=1.0.0" -o app main.go
```

### Build for Multiple Platforms

```bash
# Build for different OS/Architecture
GOOS=linux GOARCH=amd64 go build -o app-linux
GOOS=darwin GOARCH=amd64 go build -o app-mac
GOOS=windows GOARCH=amd64 go build -o app-windows.exe

# List all supported platforms
go tool dist list
```

---

## 2️⃣ CROSS-COMPILATION

### Build Script

```bash
#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o bin/app-linux-amd64
GOOS=linux GOARCH=arm64 go build -o bin/app-linux-arm64
GOOS=darwin GOARCH=amd64 go build -o bin/app-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o bin/app-darwin-arm64
GOOS=windows GOARCH=amd64 go build -o bin/app-windows-amd64.exe
```

### Makefile

```makefile
.PHONY: build build-linux build-mac build-windows clean

APP_NAME = myapp
VERSION = 1.0.0

build:
	go build -ldflags="-s -w -X main.Version=$(VERSION)" -o bin/$(APP_NAME) main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP_NAME)-linux main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP_NAME)-mac main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP_NAME)-windows.exe main.go

clean:
	rm -f bin/*
```

---

## 3️⃣ DOCKER CONTAINERIZATION

### Dockerfile

```dockerfile
# Multi-stage build
FROM golang:1.21 AS builder

WORKDIR /app
COPY . .

RUN go build -ldflags="-s -w" -o app main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]
```

### Docker Compose

```yaml
version: "3.8"

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=db
      - DB_PORT=5432
    depends_on:
      - db

  db:
    image: postgres:15
    environment:
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=password
      - POSTGRES_DB=myapp
    volumes:
      - db_data:/var/lib/postgresql/data

volumes:
  db_data:
```

### Build & Run

```bash
# Build image
docker build -t myapp:latest .

# Run container
docker run -p 8080:8080 myapp:latest

# Using docker-compose
docker-compose up -d
```

---

## 4️⃣ DEPLOYMENT STRATEGIES

### Blue-Green Deployment

```yaml
# Nginx configuration
upstream blue {
    server app-blue:8080;
}

upstream green {
    server app-green:8080;
}

server {
    listen 80;

    # Route to current version
    location / {
        proxy_pass http://blue;
    }
}
```

### Rolling Deployment

```yaml
# Kubernetes rolling update
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  template:
    spec:
      containers:
        - name: app
          image: myapp:1.1.0
```

### Canary Deployment

```yaml
# Route percentage of traffic to new version
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: myapp
spec:
  hosts:
    - myapp
  http:
    - match:
        - uri:
            prefix: /
      route:
        - destination:
            host: myapp
            subset: v1
          weight: 90
        - destination:
            host: myapp
            subset: v2
          weight: 10
```

---

## 5️⃣ SCALING PATTERNS

### Horizontal Scaling

```go
// Load balancing across multiple instances
// Use reverse proxy (Nginx, HAProxy)

// OR use Go's built-in load balancing
func startInstance(port string) {
    http.HandleFunc("/", handleRequest)
    http.ListenAndServe(":"+port, nil)
}

// Run multiple instances
for i := 0; i < 4; i++ {
    go startInstance(strconv.Itoa(8080 + i))
}
```

### Vertical Scaling

```go
// Optimize single instance
// - Use worker pools
// - Cache frequently accessed data
// - Profile and optimize bottlenecks
// - Increase resources (CPU, memory)
```

### Database Scaling

```go
// Read replicas
dsn := "primary-db:5432"

// OR use connection pooling
connPool := &pgx.ConnPool{
    MaxConnections: 50,
}

// OR shard data
func getShardID(userID string) int {
    hash := fnv.New32a()
    hash.Write([]byte(userID))
    return int(hash.Sum32()) % numShards
}
```

---

## 6️⃣ MONITORING & LOGGING

### Structured Logging

```go
import "log/slog"

// Structured logging
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

logger.Info("request handled",
    "method", "GET",
    "path", "/api/users",
    "status", 200,
    "duration_ms", 45,
)
```

### Metrics Collection

```go
import "github.com/prometheus/client_golang/prometheus"

var (
    requestCount = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
        },
        []string{"method", "endpoint"},
    )

    requestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
        },
        []string{"method", "endpoint"},
    )
)

func recordMetrics(method, endpoint string, duration time.Duration) {
    requestCount.WithLabelValues(method, endpoint).Inc()
    requestDuration.WithLabelValues(method, endpoint).Observe(duration.Seconds())
}
```

### Health Checks

```go
func healthCheck(w http.ResponseWriter, r *http.Request) {
    if !isHealthy() {
        w.WriteHeader(http.StatusServiceUnavailable)
        json.NewEncoder(w).Encode(map[string]string{"status": "unhealthy"})
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func isHealthy() bool {
    // Check database connection
    // Check cache connectivity
    // Check external service availability
    return true
}
```

---

## 7️⃣ PRODUCTION CHECKLIST

- ✅ Build optimization flags
- ✅ Docker containerization
- ✅ Proper logging setup
- ✅ Metrics collection
- ✅ Health checks implemented
- ✅ Graceful shutdown
- ✅ Rate limiting
- ✅ Load balancing
- ✅ Database backups
- ✅ Error alerting
- ✅ Performance monitoring
- ✅ Log aggregation

---

## 8️⃣ EXAMPLE: COMPLETE DEPLOYMENT

```bash
# 1. Build
make build

# 2. Create Docker image
docker build -t myapp:1.0.0 .

# 3. Push to registry
docker push registry.example.com/myapp:1.0.0

# 4. Deploy using docker-compose
docker-compose up -d

# 5. Verify
curl http://localhost:8080/health

# 6. Monitor logs
docker logs -f myapp

# 7. Scale
docker-compose up -d --scale app=3
```

---

## 📝 EXERCISES

1. **Optimization**: Build optimized binary with reduced size
2. **Docker**: Create Dockerfile and docker-compose.yml
3. **Deployment**: Deploy app with health checks and monitoring

---

## 📚 RESOURCES

- [Go Build](https://pkg.go.dev/cmd/go)
- [Docker Documentation](https://docs.docker.com/)
- [Kubernetes](https://kubernetes.io/)
- [Prometheus](https://prometheus.io/)
