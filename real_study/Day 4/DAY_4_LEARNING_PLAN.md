# 🌟 GO LEARNING - DAY 4 MASTER PLAN

**Thời gian:** 8 tiếng học (chia làm 4 session × 2 tiếng)
**Mục tiêu:** Xây dựng enterprise-grade applications với gRPC, event-driven architecture, observability, real-time systems & software architecture patterns

---

## 📅 LỊCH TRÌNH CHI TIẾT

### ⏰ Session 1 (0-2 giờ): gRPC & Event-Driven Architecture

- **Lesson 1:** gRPC & Protocol Buffers - Service definition, streaming
- **Lesson 2:** Event-Driven Architecture - Message queues, pub/sub, CQRS
- **Quiz:** 5 câu hỏi về gRPC & events

### ⏰ Session 2 (2-4 giờ): Database Patterns & Observability

- **Lesson 3:** Advanced Database Patterns - Migrations, transactions, CQRS
- **Lesson 4:** Observability & Monitoring - Structured logging, tracing, metrics
- **Mini Project 1:** gRPC Task Service

### ⏰ Session 3 (4-6 giờ): API Gateway & Real-time

- **Lesson 5:** API Gateway & Middleware - Rate limiting, auth middleware chains
- **Lesson 6:** WebSockets & Real-time - Bi-directional communication, SSE
- **Mini Project 2:** Real-time Chat Server

### ⏰ Session 4 (6-8 giờ): Advanced Testing & Architecture

- **Lesson 7:** Advanced Testing - Mocks, integration tests, table-driven tests
- **Lesson 8:** Software Architecture - DDD, hexagonal architecture, clean code
- **Mini Project 3:** Event-Driven Order System
- **Final Quiz:** 10 câu hỏi enterprise-level

---

## 📚 CÁC BÀI GIẢNG

| Số thứ tự | Bài giảng                      | Thời lượng | Mục tiêu                              |
| --------- | ------------------------------ | ---------- | ------------------------------------- |
| 1         | gRPC & Protocol Buffers        | 50 phút    | gRPC services & streaming             |
| 2         | Event-Driven Architecture      | 45 phút    | Pub/sub, message brokers, CQRS        |
| 3         | Advanced Database Patterns     | 45 phút    | Migrations, transactions, repository  |
| 4         | Observability & Monitoring     | 45 phút    | Logging, tracing, metrics             |
| 5         | API Gateway & Middleware       | 40 phút    | Rate limiting, auth, middleware chain |
| 6         | WebSockets & Real-time         | 45 phút    | WS protocol, SSE, fan-out             |
| 7         | Advanced Testing               | 45 phút    | Mocks, integration, table-driven      |
| 8         | Software Architecture Patterns | 60 phút    | DDD, hexagonal, clean architecture    |

---

## 🎯 KỸ NĂNG CẬP NHẬT ĐƯỢC

✅ gRPC service definition với Protocol Buffers
✅ Unary, server-streaming, client-streaming & bi-directional streaming
✅ Event-driven patterns (pub/sub, message queue)
✅ CQRS (Command Query Responsibility Segregation)
✅ Database migrations & schema management
✅ Advanced transaction patterns (saga, outbox)
✅ Structured logging với zerolog/zap
✅ Distributed tracing với OpenTelemetry
✅ Metrics collection & export (Prometheus)
✅ API Gateway patterns & rate limiting
✅ JWT authentication middleware
✅ WebSocket server & client
✅ Server-Sent Events (SSE)
✅ Mock generation & interface-based testing
✅ Integration test suites
✅ Domain-Driven Design (DDD) basics
✅ Hexagonal / Clean Architecture
✅ Event sourcing patterns

---

## 📋 MINI PROJECTS

### Project 1: gRPC Task Service (1.5 giờ)

- gRPC server với CRUD operations cho tasks
- Unary & server-streaming RPCs
- Interceptors (logging, auth, metrics)
- TLS / insecure connection support
- gRPC client với retry logic
- Kỹ năng: gRPC, Protobuf, Interceptors, Streaming

### Project 2: Real-time Chat Server (2 giờ)

- WebSocket-based chat server
- Multiple rooms / channels
- Message broadcast & fan-out pattern
- Presence tracking (online/offline)
- Message history (in-memory)
- Server-Sent Events fallback
- Kỹ năng: WebSocket, Concurrency, Real-time, Channels

### Project 3: Event-Driven Order System (2 giờ)

- Domain-driven order management
- Event bus (in-process pub/sub)
- CQRS: separate Command & Query models
- Saga pattern for distributed transactions
- Event sourcing basics (append-only log)
- Observability: logs, traces, metrics
- Kỹ năng: DDD, CQRS, Event Sourcing, Architecture

---

## 📊 TIẾN ĐỘ HỌC TẬP

```
Day 1: Basics          ████████████████████ 100%
Day 2: Intermediate    ████████████████████ 100%
Day 3: Advanced        ████████████████████ 100%
Day 4: Enterprise      ░░░░░░░░░░░░░░░░░░░░   0%  ← BẠN ĐANG Ở ĐÂY
```

---

## 🔗 PREREQUISITES

Trước khi bắt đầu Day 4, đảm bảo bạn đã nắm vững:

- ✅ Goroutines, channels & concurrency patterns (Day 3)
- ✅ Context & cancellation (Day 3)
- ✅ HTTP server & REST API (Day 2)
- ✅ Interfaces & error handling (Day 2)
- ✅ Testing cơ bản (Day 2)
- ✅ Docker basics (Day 3)
