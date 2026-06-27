# MINI PROJECTS GUIDE - DAY 4

## 📋 Giới thiệu

Có 3 mini projects trong Day 4 Go Learning Plan:

1. **gRPC Task Service** (1.5 tiếng) - Trung bình-Khó
2. **Real-time Chat Server** (2 tiếng) - Khó
3. **Event-Driven Order System** (2 tiếng) - Khó (++ enterprise)

---

## 🎯 Quy trình chung

### Bước 1: Tạo thư mục & khởi tạo module

```bash
# Ví dụ với project 1
cd mini_projects/01_gRPC_Task_Service
mkdir grpc_task_service
cd grpc_task_service
go mod init grpc_task_service
```

### Bước 2: Đọc README.md

Mỗi project có README chi tiết với:

- Mô tả & yêu cầu đầy đủ
- Architecture diagram
- API/Interface definitions
- Implementation steps có hướng dẫn
- Ví dụ usage & test commands
- Bonus features để thách thức thêm

### Bước 3: Thiết kế trước khi code

- Xác định interfaces & contracts
- Phân chia layers (domain, application, infrastructure)
- Thiết kế data models
- Lên kế hoạch event flow / gRPC service definitions

### Bước 4: Implement từng layer

- **Domain layer** trước: models, interfaces, business logic
- **Application layer**: use cases, event handlers, commands/queries
- **Infrastructure layer**: gRPC server, DB, WebSocket
- **Presentation layer**: wiring, main.go, config

### Bước 5: Test từng component

```bash
# Unit tests
go test ./...

# Race condition detection
go test -race ./...

# Benchmark
go test -bench=. ./...

# Integration test (nếu có)
go test -tags=integration ./...
```

### Bước 6: Add Observability

- Thêm structured logging (zerolog hoặc zap)
- Thêm metrics (counters, histograms)
- Thêm basic tracing

### Bước 7: Finalize & Review

- Chạy `go vet ./...`
- Chạy `golangci-lint run`
- Kiểm tra không có data races
- Đọc lại code và refactor nếu cần

---

## 📊 PROJECT COMPARISON

| Project                   | Difficulty | Time | Core Skills                              |
| ------------------------- | ---------- | ---- | ---------------------------------------- |
| gRPC Task Service         | ⭐⭐⭐⭐   | 1.5h | gRPC, Protobuf, Interceptors, Streaming  |
| Real-time Chat Server     | ⭐⭐⭐⭐   | 2h   | WebSocket, Concurrency, Fan-out, Rooms   |
| Event-Driven Order System | ⭐⭐⭐⭐⭐ | 2h   | DDD, CQRS, Event Sourcing, Observability |

---

## 🔧 DEPENDENCIES THƯỜNG DÙNG

### gRPC Project

```go
// go.mod dependencies
require (
    google.golang.org/grpc v1.60.0
    google.golang.org/protobuf v1.31.0
)
```

### WebSocket Project

```go
require (
    github.com/gorilla/websocket v1.5.1
)
```

### Observability

```go
require (
    go.uber.org/zap v1.26.0
    go.opentelemetry.io/otel v1.21.0
    github.com/prometheus/client_golang v1.17.0
)
```

### Testing

```go
require (
    github.com/stretchr/testify v1.8.4
    github.com/golang/mock v1.6.0
)
```

---

## ✅ CHECKLIST TRƯỚC KHI SUBMIT

- [ ] Code compile thành công (`go build ./...`)
- [ ] Tất cả tests pass (`go test ./...`)
- [ ] Không có data race (`go test -race ./...`)
- [ ] Không có lỗi lint nghiêm trọng
- [ ] README.md được cập nhật với hướng dẫn chạy
- [ ] Các edge cases đã được xử lý
- [ ] Graceful shutdown được implement
