# Project 1: Microservice API with Context

## 🎯 Mục tiêu

Xây dựng một **RESTful API microservice** với:

- ✅ Multiple endpoints
- ✅ Request context handling (timeouts, cancellation)
- ✅ Worker pool for concurrent processing
- ✅ Graceful shutdown
- ✅ Input validation
- ✅ JSON request/response handling

---

## 📋 Yêu cầu

### Core Features

1. **API Endpoints**
   - `GET /health` - Health check
   - `POST /tasks` - Create task
   - `GET /tasks` - List all tasks
   - `GET /tasks/{id}` - Get task by ID
   - `PUT /tasks/{id}` - Update task
   - `DELETE /tasks/{id}` - Delete task

2. **Request Context**
   - Implement timeout for all requests (5 seconds)
   - Add request ID tracking
   - Support cancellation via context

3. **Worker Pool**
   - Use worker pool to process heavy operations
   - Configurable number of workers (default: 5)
   - Queue management

4. **Data Models**

   ```go
   type Task struct {
       ID        string    `json:"id"`
       Title     string    `json:"title"`
       Status    string    `json:"status"` // pending, in-progress, completed
       CreatedAt time.Time `json:"created_at"`
       UpdatedAt time.Time `json:"updated_at"`
   }
   ```

5. **Error Handling**
   - Proper HTTP status codes
   - JSON error responses
   - Request validation

6. **Graceful Shutdown**
   - Handle SIGINT/SIGTERM
   - Wait for in-flight requests
   - Close resources properly

---

## 🏗️ Architecture

```
api/
├── main.go              # Entry point
├── handlers.go          # HTTP handlers
├── models.go            # Data structures
├── service.go           # Business logic
├── worker.go            # Worker pool
└── middleware.go        # Context middleware
```

---

## 📝 Implementation Steps

### Step 1: Define Models

```go
type Task struct {
    ID        string    `json:"id"`
    Title     string    `json:"title"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type CreateTaskRequest struct {
    Title string `json:"title"`
}
```

### Step 2: Implement Service Layer

```go
type TaskService struct {
    tasks map[string]*Task
    mu    sync.RWMutex
}

func (s *TaskService) GetTask(ctx context.Context, id string) (*Task, error) {
    // Implement with context handling
}
```

### Step 3: Create Worker Pool

```go
type WorkerPool struct {
    jobs    chan Job
    workers int
}

func (wp *WorkerPool) Start() {
    for i := 0; i < wp.workers; i++ {
        go wp.worker()
    }
}
```

### Step 4: Implement HTTP Handlers

```go
func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
    var req CreateTaskRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Validate input
    if err := validateTaskInput(req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Create task
    task, err := h.service.CreateTask(r.Context(), req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(task)
}
```

### Step 5: Add Context Middleware

```go
func contextMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Add request ID
        requestID := generateRequestID()
        ctx := context.WithValue(r.Context(), "request_id", requestID)

        // Add timeout
        ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
        defer cancel()

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

### Step 6: Implement Graceful Shutdown

```go
func main() {
    server := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        <-sigCh
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        defer cancel()
        server.Shutdown(ctx)
    }()

    server.ListenAndServe()
}
```

---

## ✅ Test Cases

```bash
# 1. Health check
curl http://localhost:8080/health

# 2. Create task
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go"}'

# 3. List tasks
curl http://localhost:8080/tasks

# 4. Get specific task
curl http://localhost:8080/tasks/{id}

# 5. Update task
curl -X PUT http://localhost:8080/tasks/{id} \
  -H "Content-Type: application/json" \
  -d '{"title": "Learn Go", "status": "in-progress"}'

# 6. Delete task
curl -X DELETE http://localhost:8080/tasks/{id}

# 7. Test timeout (should fail after 5s)
curl http://localhost:8080/slow-endpoint
```

---

## 📊 Performance Targets

- **Response time**: < 50ms (p50), < 100ms (p99)
- **Throughput**: > 1000 req/sec
- **Memory**: < 100MB
- **Error rate**: < 0.1%

---

## 🌟 Bonus Features

- [ ] Request logging middleware
- [ ] Rate limiting
- [ ] CORS support
- [ ] OpenAPI/Swagger documentation
- [ ] Database persistence (SQLite)
- [ ] Caching layer
- [ ] Metrics collection (Prometheus)
- [ ] Tracing support

---

## 📚 Resources

- [Net/HTTP Package](https://pkg.go.dev/net/http)
- [Context Package](https://pkg.go.dev/context)
- [Best Practices](https://go.dev/doc/effective_go)

---

## 💡 Learning Goals

✨ Understand context handling in HTTP requests
✨ Implement worker pools for concurrent operations
✨ Design RESTful APIs with proper error handling
✨ Graceful shutdown patterns
✨ Request-scoped values and timeouts
