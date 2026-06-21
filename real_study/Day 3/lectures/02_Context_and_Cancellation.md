# Lesson 2: Context & Cancellation

## 📖 Nội dung bài học

1. Context là gì?
2. Creating contexts
3. Context values (request-scoped data)
4. Timeouts & deadlines
5. Cancellation patterns
6. Best practices

---

## 1️⃣ CONTEXT LÀ GÌ?

### Định nghĩa

**Context** là một package cung cấp cách để:

- ✅ Pass request-scoped values
- ✅ Set timeouts & deadlines
- ✅ Signal cancellation giữa goroutines

### Tại sao cần Context?

```go
// ❌ Sai: Không có cách để cancel request
func handleRequest(conn net.Conn) {
    data := readData(conn) // Có thể hang vĩnh viễn
    process(data)
}

// ✅ Đúng: Sử dụng context
func handleRequest(ctx context.Context, conn net.Conn) {
    data := readDataWithContext(ctx, conn) // Có thể cancel
    process(ctx, data)
}
```

### Interface Context

```go
type Context interface {
    // Deadline trả về deadline nếu có
    Deadline() (deadline time.Time, ok bool)

    // Done trả về channel được close khi context cancelled
    Done() <-chan struct{}

    // Err trả về error (Canceled hoặc DeadlineExceeded)
    Err() error

    // Value trả về value cho key
    Value(key interface{}) interface{}
}
```

---

## 2️⃣ CREATING CONTEXTS

### Background Context

```go
// Root context, không bao giờ cancelled
ctx := context.Background()
```

### TODO Context

```go
// Dùng khi chưa biết sẽ dùng context nào
ctx := context.TODO()
```

### Child Contexts

```go
// Cancel context
ctx, cancel := context.WithCancel(parent)
defer cancel()

// Timeout context (cancel sau 5 giây)
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()

// Deadline context (cancel tại time cụ thể)
deadline := time.Now().Add(5*time.Second)
ctx, cancel := context.WithDeadline(parent, deadline)
defer cancel()

// With values
ctx = context.WithValue(parent, "user_id", 123)
```

---

## 3️⃣ CONTEXT VALUES (Request-Scoped Data)

### Lưu Values

```go
// Lưu value vào context
ctx := context.WithValue(context.Background(), "user_id", 42)
ctx = context.WithValue(ctx, "request_id", "abc123")
```

### Truy Cập Values

```go
// Truy cập value
userId := ctx.Value("user_id").(int)        // Type assertion
requestId := ctx.Value("request_id").(string)

// Safer way
userId, ok := ctx.Value("user_id").(int)
if !ok {
    // Value not found or wrong type
}
```

### Best Practice: Use Custom Types for Keys

```go
// Define custom type
type contextKey string

const (
    userIDKey contextKey = "user_id"
    requestIDKey contextKey = "request_id"
)

// Lưu value
ctx := context.WithValue(context.Background(), userIDKey, 42)

// Truy cập value
userId := ctx.Value(userIDKey).(int)
```

### Example: Middleware with Context Values

```go
// Middleware: Thêm user info vào context
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract user từ request
        userId := getUserFromToken(r)

        // Tạo new context với user info
        ctx := context.WithValue(r.Context(), userIDKey, userId)

        // Pass context vào next handler
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Handler: Truy cập user từ context
func handleRequest(w http.ResponseWriter, r *http.Request) {
    userId := r.Context().Value(userIDKey).(int)
    // Handle request với user info
}
```

---

## 4️⃣ TIMEOUTS & DEADLINES

### Timeout Pattern

```go
// Tạo context với timeout 5 giây
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Gọi function với timeout
result, err := slowFunction(ctx)
if err == context.DeadlineExceeded {
    fmt.Println("Request timeout!")
}
```

### Deadline Pattern

```go
// Tạo deadline cụ thể
deadline := time.Now().Add(10*time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()

// Kiểm tra deadline
if deadline, ok := ctx.Deadline(); ok {
    fmt.Println("Deadline:", deadline)
}
```

### Example: HTTP Request with Timeout

```go
func fetchURL(url string) (string, error) {
    // Timeout 10 giây
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Tạo request với context
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return "", err
    }

    // Thực thi request
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err // Có thể là deadline exceeded
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    return string(body), nil
}
```

---

## 5️⃣ CANCELLATION PATTERNS

### Cancel Function

```go
// Tạo cancel context
ctx, cancel := context.WithCancel(context.Background())

// Cancel context
cancel()

// Kiểm tra cancellation
select {
case <-ctx.Done():
    fmt.Println("Context cancelled:", ctx.Err())
}
```

### Example: Worker Pool with Cancellation

```go
func workerPool(ctx context.Context, numWorkers int, jobs <-chan Job) {
    for i := 0; i < numWorkers; i++ {
        go func(id int) {
            for {
                select {
                case job := <-jobs:
                    processJob(ctx, job)
                case <-ctx.Done():
                    fmt.Printf("Worker %d stopped\n", id)
                    return
                }
            }
        }(i)
    }
}
```

### Example: Graceful Shutdown

```go
func main() {
    // Create cancel context
    ctx, cancel := context.WithCancel(context.Background())

    // Handle signals
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        <-sigCh
        fmt.Println("Shutting down...")
        cancel() // Cancel context
    }()

    // Start server with context
    server := &http.Server{Addr: ":8080"}
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Println("Server error:", err)
        }
    }()

    // Wait for cancellation
    <-ctx.Done()

    // Graceful shutdown
    shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := server.Shutdown(shutdownCtx); err != nil {
        log.Println("Shutdown error:", err)
    }
}
```

---

## 6️⃣ BEST PRACTICES

### 1. Always Use Context

```go
// ✅ Đúng: Hàm nhận context
func doWork(ctx context.Context) error {
    // Check context
}

// ❌ Sai: Hàm không nhận context
func doWork() error {
    // Không thể cancel hoặc timeout
}
```

### 2. Pass Context as First Parameter

```go
// ✅ Đúng
func fetchData(ctx context.Context, url string) ([]byte, error) {
    // ...
}

// ❌ Sai
func fetchData(url string, ctx context.Context) ([]byte, error) {
    // ...
}
```

### 3. Don't Store Context in Structs

```go
// ❌ Sai
type MyStruct struct {
    ctx context.Context // Lỗi!
}

// ✅ Đúng
func (m *MyStruct) doWork(ctx context.Context) {
    // Nhận context từ parameter
}
```

### 4. Check Context Before Long Operations

```go
// ✅ Đúng: Check context regularly
func processItems(ctx context.Context, items []Item) error {
    for _, item := range items {
        select {
        case <-ctx.Done():
            return ctx.Err()
        default:
            process(item)
        }
    }
    return nil
}
```

### 5. Use context.Value for Request-Scoped Data Only

```go
// ✅ Đúng: Request-scoped data
type key string
const userKey key = "user"

ctx := context.WithValue(ctx, userKey, user)
user := ctx.Value(userKey)

// ❌ Sai: Configuration hoặc static data
ctx := context.WithValue(ctx, "config", globalConfig)
```

---

## 💡 EXAMPLES

### Example: API Handler with Timeout & Validation

```go
func handleRequest(ctx context.Context, req Request) (Response, error) {
    // Timeout 5 giây
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    // Thêm request ID
    requestID := generateID()
    ctx = context.WithValue(ctx, "request_id", requestID)

    // Validate input
    if err := validateRequest(ctx, req); err != nil {
        return Response{}, err
    }

    // Fetch data
    data, err := fetchData(ctx, req.URL)
    if err != nil {
        return Response{}, err
    }

    return Response{Data: data}, nil
}
```

---

## 📝 EXERCISES

1. **Timeout**: Tạo function gọi API với timeout 5 giây
2. **Cancellation**: Implement worker pool có thể cancel
3. **Values**: Tạo middleware thêm user info vào context
4. **Graceful Shutdown**: Implement HTTP server với graceful shutdown

---

## 📚 RESOURCES

- [Context Package](https://pkg.go.dev/context)
- [Using Context to Avoid Leaks](https://go.dev/blog/context)
- [Context Patterns](https://www.digitalocean.com/community/tutorials/how-to-use-context-in-go)
