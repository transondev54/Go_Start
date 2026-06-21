# MINI PROJECTS GUIDE - DAY 3

## 📋 Giới thiệu

Có 3 mini projects trong Day 3 Go Learning Plan:

1. **Microservice API** (1.5 tiếng) - Bình thường
2. **CLI Tool** (2 tiếng) - Khó
3. **Distributed Cache** (2 tiếng) - Khó (++advanced)

---

## 🎯 Quy trình chung

### Bước 1: Tạo thư mục

```bash
cd mini_projects/01_Microservice_API
mkdir api
cd api
go mod init microservice
code main.go
```

### Bước 2: Đọc README.md

Mỗi project có README chi tiết với:

- Mô tả & yêu cầu
- Architecture overview
- Implementation steps
- Ví dụ usage
- Performance targets
- Bonus features

### Bước 3: Thiết kế Architecture

Trước khi code:

- Vẽ design diagram
- Lên kế hoạch structure
- Xác định interfaces
- Phân tích concurrency

### Bước 4: Implement từ Bottom-up

- Bắt đầu từ core logic
- Thêm concurrency
- Integrate context
- Add error handling

### Bước 5: Test & Benchmark

- Unit tests
- Integration tests
- Benchmark performance
- Profile memory

### Bước 6: Optimize & Deploy

- Optimize hot paths
- Add monitoring
- Docker containerization
- Final polish

### Bước 7: Submit

- Lưu file `main.go`
- Test lại lần cuối
- Commit với git
- Submit để nhận feedback

---

## 📊 PROJECT COMPARISON

| Project           | Difficulty | Time | Core Skills                         |
| ----------------- | ---------- | ---- | ----------------------------------- |
| Microservice API  | ⭐⭐⭐     | 1.5h | Context, validation, HTTP, patterns |
| CLI Tool          | ⭐⭐⭐⭐   | 2h   | Concurrency, performance, design    |
| Distributed Cache | ⭐⭐⭐⭐⭐ | 2h   | Advanced patterns, reflection, sync |

---

## ✅ CHECKLIST cho mỗi project

- [ ] Tạo project directory
- [ ] `go mod init`
- [ ] Đọc README.md chi tiết
- [ ] Thiết kế architecture
- [ ] Viết main.go từ scratch
- [ ] Implement core features
- [ ] Add concurrency & context
- [ ] Code compile không error
- [ ] Chạy được chương trình
- [ ] Test với ít nhất 10 test cases
- [ ] Benchmark performance
- [ ] Xử lý errors & edge cases
- [ ] Add proper logging
- [ ] Code formatted & readable
- [ ] Add comments cho functions
- [ ] Write unit tests
- [ ] Commit với git
- [ ] Documentation complete

---

## 🔄 DAY 3 WORKFLOW

```
Lesson 1 (Concurrency) → Quick practice
      ↓
Lesson 2 (Context)     → Microservice API project start
      ↓
Lesson 3 (Reflection)  → Implement reflection features
      ↓
Lesson 4 (Generics)    → Add generic components
      ↓
Lesson 5 (Profiling)   → CLI Tool project + benchmark
      ↓
Lesson 6 (Optimization)→ Optimize both projects
      ↓
Lesson 7 (Security)    → Add security features
      ↓
Lesson 8 (Deployment)  → Distributed Cache + Docker
```

---

## 📈 PROJECT DIFFICULTY PROGRESSION

### Project 1: Microservice API

- **Focus**: HTTP, context, validation
- **Concurrency**: Moderate (request handling)
- **Architecture**: Simple REST endpoints
- **Learning**: Request-scoped values, graceful shutdown

### Project 2: CLI Tool

- **Focus**: Concurrency, performance, UX
- **Concurrency**: High (worker pools, batching)
- **Architecture**: Command pattern, subcommands
- **Learning**: Performance profiling, user feedback

### Project 3: Distributed Cache

- **Focus**: Advanced patterns, reflection, sync
- **Concurrency**: Very high (thread-safe access)
- **Architecture**: Complex with multiple components
- **Learning**: Eviction policies, TTL management, statistics

---

## 🧪 TESTING STRATEGY

### Unit Tests

```go
func TestCacheSet(t *testing.T) {
    // Test cache SET operation
}

func TestCacheGet(t *testing.T) {
    // Test cache GET operation
}
```

### Integration Tests

```go
// Test API endpoints
// Test concurrent cache access
// Test context cancellation
```

### Benchmarks

```bash
go test -bench=. -benchmem
```

---

## 📊 PERFORMANCE TARGETS

### Microservice API

- **Response time**: < 50ms (p50), < 100ms (p99)
- **Throughput**: > 1000 req/sec
- **Memory**: < 100MB

### CLI Tool

- **File processing**: > 100 files/sec
- **Memory efficiency**: < 500MB for large datasets
- **CPU**: < 80% utilization

### Distributed Cache

- **Get latency**: < 1ms
- **Set latency**: < 1ms
- **Throughput**: > 100k ops/sec
- **Memory efficiency**: High (configurable eviction)

---

## 🔍 CODE REVIEW CHECKLIST

Before submitting, check:

- ✅ **Functionality**: All features work as specified
- ✅ **Performance**: Meets or exceeds targets
- ✅ **Concurrency**: Safe concurrent access
- ✅ **Error Handling**: Graceful error handling
- ✅ **Code Quality**: Clean, readable, maintainable
- ✅ **Testing**: Comprehensive test coverage (> 70%)
- ✅ **Documentation**: Clear comments & README
- ✅ **Best Practices**: Follows Go conventions
- ✅ **Security**: Input validation, safe operations
- ✅ **Monitoring**: Logging, metrics, observability

---

## 📚 RESOURCES

### Profiling & Optimization

- `go test -bench` - Benchmark tests
- `pprof` - CPU & memory profiling
- `benchstat` - Statistical comparison
- `trace` - Execution trace analysis

### Tools

```bash
# Benchmarking
go test -bench=. -benchmem ./...

# Profiling
go tool pprof http://localhost:6060/debug/pprof/profile

# Code quality
go vet ./...
golangci-lint run
```

### Documentation

- [Go Best Practices](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Concurrent Go](https://go.dev/blog/concurrency-patterns)
