# Lesson 1: Advanced Concurrency Patterns

## 📖 Nội dung bài học

1. Review: Goroutines & Channels cơ bản
2. Worker Pool pattern
3. Fan-in & Fan-out patterns
4. Pipeline pattern
5. Practical examples & best practices

---

## 1️⃣ REVIEW: GOROUTINES & CHANNELS

### Goroutines

**Goroutine** là lightweight thread được quản lý bởi Go runtime.

```go
// Tạo goroutine
go doSomething()

// Goroutine with anonymous function
go func() {
    fmt.Println("Goroutine running")
}()
```

### Channels

**Channel** là cách để goroutines giao tiếp an toàn.

```go
// Tạo channel
ch := make(chan int)

// Gửi dữ liệu
ch <- value

// Nhận dữ liệu
value := <-ch

// Buffered channel
ch := make(chan int, 10)
```

---

## 2️⃣ WORKER POOL PATTERN

### Định nghĩa

**Worker Pool** là pattern để quản lý số lượng goroutines, tránh tạo quá nhiều.

### Vấn đề

```go
// ❌ Sai: Tạo quá nhiều goroutines
for i := 0; i < 1000000; i++ {
    go processItem(items[i]) // Quá nhiều goroutines!
}
```

### Giải pháp

```go
// ✅ Đúng: Worker pool với số lượng cố định
numWorkers := 10
jobs := make(chan Item, 100)
results := make(chan Result, 100)

// Khởi động workers
for w := 0; w < numWorkers; w++ {
    go worker(w, jobs, results)
}

// Gửi jobs
for i, item := range items {
    jobs <- item
}
close(jobs)

// Nhận results
for i := 0; i < len(items); i++ {
    result := <-results
    processResult(result)
}
```

### Cấu trúc Worker

```go
func worker(id int, jobs <-chan Item, results chan<- Result) {
    for job := range jobs {
        // Process job
        result := processItem(job)
        results <- result
    }
}
```

### Ưu điểm

✅ Kiểm soát resource usage
✅ Tránh goroutine leak
✅ Dễ scale up/down
✅ Predictable performance

---

## 3️⃣ FAN-IN & FAN-OUT PATTERNS

### Fan-Out Pattern

**Fan-Out** là gửi một task cho nhiều worker.

```go
// Fan-out: Một input → Nhiều output channels
func fanOut(input <-chan Item, numWorkers int) []<-chan Result {
    outputs := make([]<-chan Result, numWorkers)

    for i := 0; i < numWorkers; i++ {
        out := make(chan Result)
        outputs[i] = out

        go func(out chan Result) {
            for item := range input {
                out <- processItem(item)
            }
            close(out)
        }(out)
    }

    return outputs
}
```

### Fan-In Pattern

**Fan-In** là nhận từ nhiều channels và gộp thành một.

```go
// Fan-in: Nhiều input → Một output
func fanIn(inputs ...<-chan Result) <-chan Result {
    var wg sync.WaitGroup
    out := make(chan Result)

    // Khởi động goroutine cho mỗi input
    for _, in := range inputs {
        wg.Add(1)
        go func(in <-chan Result) {
            defer wg.Done()
            for result := range in {
                out <- result
            }
        }(in)
    }

    // Đóng output khi tất cả goroutines kết thúc
    go func() {
        wg.Wait()
        close(out)
    }()

    return out
}
```

### Ví dụ: Fan-Out + Fan-In

```go
// 1. Fan-out: Gửi items cho 5 workers
inputs := fanOut(itemsChan, 5)

// 2. Fan-in: Nhận results từ tất cả workers
results := fanIn(inputs...)

// 3. Xử lý results
for result := range results {
    fmt.Println(result)
}
```

---

## 4️⃣ PIPELINE PATTERN

### Định nghĩa

**Pipeline** là chuỗi các stage, mỗi stage nhận input từ stage trước.

### Cấu trúc

```go
// Stage 1: Generate numbers
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Stage 2: Square numbers
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Stage 3: Print results
func print(in <-chan int) {
    for n := range in {
        fmt.Println(n)
    }
}

// Sử dụng pipeline
func main() {
    numbers := generate(2, 3, 4, 5)
    squares := square(numbers)
    print(squares)
}
```

### Multi-Stage Pipeline

```go
func main() {
    // Pipeline: generate → filter → transform → aggregate

    numbers := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

    // Filter: chỉ lấy số chẵn
    evens := filter(numbers, func(n int) bool {
        return n%2 == 0
    })

    // Transform: nhân đôi
    doubled := transform(evens, func(n int) int {
        return n * 2
    })

    // Print results
    for result := range doubled {
        fmt.Println(result) // 4, 8, 12, 16, 20
    }
}
```

---

## 5️⃣ BEST PRACTICES

### 1. Đóng Channels Đúng Cách

```go
// ✅ Đúng: Sender đóng channel
func sender(out chan int) {
    defer close(out) // Chỉ sender đóng
    for i := 0; i < 5; i++ {
        out <- i
    }
}

// ❌ Sai: Receiver đóng channel (panic!)
ch := make(chan int)
go sender(ch)
close(ch) // panic: close of closed channel
```

### 2. Tránh Deadlock

```go
// ❌ Sai: Deadlock
func deadlock() {
    ch := make(chan int) // Unbuffered channel
    ch <- 1 // Blocked, no one to receive!
    <-ch
}

// ✅ Đúng: Sử dụng buffered channel hoặc goroutine
func notDeadlock() {
    ch := make(chan int, 1) // Buffered
    ch <- 1
    fmt.Println(<-ch)
}
```

### 3. Sử dụng Select cho Multiple Channels

```go
func multiplexChannels(ch1, ch2 <-chan int) {
    for {
        select {
        case val := <-ch1:
            fmt.Println("From ch1:", val)
        case val := <-ch2:
            fmt.Println("From ch2:", val)
        case <-time.After(1 * time.Second):
            fmt.Println("Timeout!")
            return
        }
    }
}
```

### 4. Cleanup Goroutines

```go
// ✅ Đúng: Đợi tất cả goroutines kết thúc
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        doWork()
    }()
}

wg.Wait() // Đợi tất cả kết thúc
```

---

## 💡 EXAMPLES

### Example 1: Concurrent File Processing

```go
func processFiles(filePaths []string, numWorkers int) {
    jobs := make(chan string, len(filePaths))
    results := make(chan error, len(filePaths))

    // Start workers
    for i := 0; i < numWorkers; i++ {
        go func() {
            for path := range jobs {
                results <- processFile(path)
            }
        }()
    }

    // Send jobs
    for _, path := range filePaths {
        jobs <- path
    }
    close(jobs)

    // Collect results
    for i := 0; i < len(filePaths); i++ {
        if err := <-results; err != nil {
            log.Println("Error:", err)
        }
    }
}
```

---

## 📝 EXERCISES

1. **Worker Pool**: Tạo worker pool xử lý 1000 items với 10 workers
2. **Fan-In/Out**: Implement fan-out → transform → fan-in pipeline
3. **Pipeline**: Tạo 4-stage pipeline: generate → filter → map → print

---

## 📚 RESOURCES

- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Advanced Go Concurrency](https://go.dev/blog/io2013-talk-concurrency)
- [Channel Patterns](https://pkg.go.dev/sync)
