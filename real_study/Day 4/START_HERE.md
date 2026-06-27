# ⭐ START HERE - GO LEARNING DAY 4

## 🎯 Chào mừng đến Day 4!

Bạn đã hoàn thành **Day 1, 2 & 3** - Từ beginner đến advanced developer! 🎉

**Day 4** sẽ đưa bạn từ advanced → **enterprise/architect level**, với focus vào:

- ✅ **gRPC & Protobuf** - Modern service-to-service communication
- ✅ **Event-Driven Architecture** - Pub/sub, message queues, CQRS
- ✅ **Advanced Database Patterns** - Migrations, saga transactions
- ✅ **Observability** - Structured logging, distributed tracing, metrics
- ✅ **API Gateway** - Rate limiting, middleware chains, auth
- ✅ **WebSockets & Real-time** - Bi-directional communication
- ✅ **Advanced Testing** - Mocks, integration tests
- ✅ **Software Architecture** - DDD, hexagonal, clean architecture

---

## 📅 LỊCH TRÌNH NHANH

| Thời gian | Nội dung                                  | Thời lượng |
| --------- | ----------------------------------------- | ---------- |
| **0h**    | Lesson 1: gRPC & Protocol Buffers         | 1 giờ      |
| **1h**    | Lesson 2: Event-Driven Architecture       | 1 giờ      |
| **2h**    | Lesson 3: Advanced Database Patterns      | 1 giờ      |
| **3h**    | Mini Project 1: gRPC Task Service         | 1.5 giờ    |
| **4.5h**  | Lesson 4: Observability & Monitoring      | 1 giờ      |
| **5.5h**  | Lesson 5: API Gateway & Middleware        | 1 giờ      |
| **6.5h**  | Mini Project 2: Real-time Chat Server     | 2 giờ      |
| **8.5h**  | Lesson 6: WebSockets & Real-time          | 1 giờ      |
| **9.5h**  | Lesson 7: Advanced Testing                | 1 giờ      |
| **10.5h** | Lesson 8: Software Architecture Patterns  | 1.5 giờ    |
| **12h**   | Mini Project 3: Event-Driven Order System | 2 giờ      |
| **14h**   | Review & Polish                           | 1-2 giờ    |

---

## 📚 CÓ GÌ TRONG DAY 4

### 📖 8 Bài giảng

1. **gRPC & Protocol Buffers** - Service definition, streaming, interceptors
2. **Event-Driven Architecture** - Pub/sub, CQRS, outbox pattern
3. **Advanced Database Patterns** - Migrations, transactions, saga
4. **Observability & Monitoring** - zerolog, OpenTelemetry, Prometheus
5. **API Gateway & Middleware** - Rate limiting, JWT auth, middleware chains
6. **WebSockets & Real-time** - WS protocol, rooms, SSE
7. **Advanced Testing** - Mocks, testcontainers, fuzzing
8. **Software Architecture** - DDD, hexagonal, event sourcing

### 🏗️ 3 Mini Projects

| Project                       | Level      | Focus                                     |
| ----------------------------- | ---------- | ----------------------------------------- |
| **gRPC Task Service**         | ⭐⭐⭐⭐   | gRPC, protobuf, interceptors, streaming   |
| **Real-time Chat Server**     | ⭐⭐⭐⭐   | WebSocket, concurrency, fan-out, presence |
| **Event-Driven Order System** | ⭐⭐⭐⭐⭐ | DDD, CQRS, event sourcing, observability  |

---

## 🚀 GETTING STARTED

### Bước 1: Kiểm tra Prerequisites

Đảm bảo bạn đã:

- ✅ Hoàn thành Day 1, 2 & 3
- ✅ Hiểu goroutines, channels & concurrency patterns
- ✅ Biết HTTP server & REST API
- ✅ Cài đặt Go 1.21+
- ✅ Cài đặt `protoc` compiler (cho gRPC)
- ✅ Cài đặt Docker

### Bước 2: Cài đặt tools cần thiết

```bash
# Install protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Install mockgen
go install github.com/golang/mock/mockgen@latest

# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

### Bước 3: Đọc DAY_4_LEARNING_PLAN.md

Xem chi tiết kế hoạch học tập cho 8 giờ.

### Bước 4: Bắt đầu với Lesson 1

- Mở `lectures/01_gRPC_and_Protobuf.md`
- Đọc kỹ từng phần, chạy các ví dụ code

---

## 💡 MẸO HỌC HIỆU QUẢ

1. **Gõ code bằng tay** - Đừng copy-paste, hãy tự gõ để nhớ sâu hơn
2. **Chạy từng ví dụ** - Mỗi code snippet đều có thể chạy được
3. **Thử nghiệm biến thể** - Thay đổi parameters, thêm features nhỏ
4. **Đọc error messages** - Lỗi là bạn, đừng sợ lỗi
5. **Review lại Day 3** - Nếu cần, ôn lại context & concurrency trước

---

## 🎯 MỤC TIÊU CUỐI DAY 4

Sau khi hoàn thành, bạn có thể:

- [ ] Viết gRPC services với protobuf từ đầu
- [ ] Thiết kế event-driven systems với pub/sub
- [ ] Implement CQRS pattern trong Go
- [ ] Thêm observability (logs, traces, metrics) vào bất kỳ service nào
- [ ] Xây dựng real-time WebSocket server
- [ ] Viết comprehensive test suite với mocks
- [ ] Áp dụng DDD & hexagonal architecture
- [ ] Build & explain kiến trúc cho production system
