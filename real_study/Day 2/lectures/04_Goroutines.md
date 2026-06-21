# Lesson 4: Goroutines & Concurrency Basics

## 📖 Nội dung bài học

1. Goroutines là gì?
2. Tạo & quản lý goroutines
3. Channels - giao tiếp giữa goroutines
4. Buffered vs Unbuffered channels
5. Select statement

---

## 1️⃣ GOROUTINES LÀ GÌ?

### Định nghĩa

**Goroutine** là lightweight thread được quản lý bởi Go runtime. Rất rẻ và dễ tạo.

### So sánh

|               | Thread    | Goroutine     |
| ------------- | --------- | ------------- |
| **Memory**    | ~1MB      | ~2KB          |
| **Số lượng**  | Hàng trăm | Hàng triệu    |
| **Startup**   | Chậm      | Rất nhanh     |
| **Giao tiếp** | Phức tạp  | Dễ (channels) |

---

## 2️⃣ TẠO GOROUTINES

### Cú pháp cơ bản

```go
// Tạo goroutine
go function()
```

### Ví dụ

```go
func greet(name string) {
    fmt.Println(\"Hello\", name)
}

func main() {
    go greet(\"Alice\")
    go greet(\"Bob\")

    // Main goroutine quá nhanh, có thể kết thúc trước
    // other goroutines hoàn thành
    time.Sleep(1 * time.Second)  // ❌ Ugly way
}
```

### Vấn đề: Race condition

```go
var count int

func increment() {
    for i := 0; i < 1000; i++ {
        count++  // ❌ Race condition!
    }
}

func main() {
    go increment()
    go increment()

    time.Sleep(1 * time.Second)
    fmt.Println(count)  // Output: 1000-2000? Unpredictable!
}
```

---

## 3️⃣ CHANNELS - GỬI & NHẬN

### Định nghĩa channels

```go
// Channel của int
var ch chan int

// Tạo channel
ch := make(chan int)
ch := make(chan string)
```

### Gửi & nhận

```go
ch := make(chan int)

// Gửi
ch <- 42

// Nhận
value := <-ch
```

### Ví dụ: Simple communication

```go
func main() {
    ch := make(chan string)

    go func() {
        ch <- \"Hello from goroutine\"  // Gửi
    }()

    message := <-ch  // Nhận
    fmt.Println(message)
}
```

---

## 4️⃣ BUFFERED VS UNBUFFERED CHANNELS

### Unbuffered channels

```go
ch := make(chan int)  // Unbuffered

// Sender block cho đến khi có receiver
go func() {
    ch <- 42  // Chờ có ai nhận
}()

value := <-ch  // Nhận
fmt.Println(value)
```

### Buffered channels

```go
ch := make(chan int, 3)  // Buffer size 3

// Sender không block nếu buffer chưa đầy
ch <- 1
ch <- 2
ch <- 3

// Gửi thêm sẽ block
go func() {
    ch <- 4  // Chờ đến khi có space
}()

v1 := <-ch
v2 := <-ch
```

---

## 5️⃣ RANGE & CLOSE

### Close channels

```go
func main() {
    ch := make(chan int)

    go func() {
        for i := 1; i <= 3; i++ {
            ch <- i
        }
        close(ch)  // Signal done
    }()

    for value := range ch {
        fmt.Println(value)
    }
    // Output: 1, 2, 3
}
```

### Kiểm tra closed channel

```go
value, ok := <-ch
if !ok {
    fmt.Println(\"Channel closed\")
}
```

---

## 6️⃣ SELECT STATEMENT

### Chờ multiple channels

```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- \"Result from ch1\"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- \"Result from ch2\"
    }()

    // Chờ cái nào đến trước
    select {
    case result := <-ch1:
        fmt.Println(result)  // Nhận được trước
    case result := <-ch2:
        fmt.Println(result)
    }
}
```

### Select với timeout

```go
select {
case result := <-ch:
    fmt.Println(result)
case <-time.After(3 * time.Second):
    fmt.Println(\"Timeout!\")
}
```

---

## 7️⃣ VÍ DỤ: ASYNC OPERATIONS

```go
func fetchURL(url string, results chan string) {
    resp, err := http.Get(url)
    if err != nil {
        results <- \"Error: \" + err.Error()
        return
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    results <- string(body)
}

func main() {
    urls := []string{
        \"https://example.com\",
        \"https://google.com\",
        \"https://github.com\",
    }

    results := make(chan string, len(urls))

    // Tạo goroutines cho mỗi URL
    for _, url := range urls {
        go fetchURL(url, results)
    }

    // Nhận kết quả
    for i := 0; i < len(urls); i++ {
        fmt.Println(<-results)
    }
}
```

---

## 📝 TÓM TẮT CONCURRENCY

```go
// Tạo goroutine
go function()

// Channel
ch := make(chan int)
ch <- value           // Gửi
value := <-ch         // Nhận

// Buffered
ch := make(chan int, 10)

// Select
select {
case v := <-ch1:
case v := <-ch2:
}
```

---

## ⚠️ GOTCHAS

1. **Goroutine leak** - goroutines chưa kết thúc khi program dừng
2. **Deadlock** - channels chờ forever
3. **Race conditions** - shared memory mà không sync

---

## 💡 BEST PRACTICES

1. **Luôn close channels** - từ sender side
2. **Tránh shared memory** - dùng channels to communicate
3. **Use select với timeout** - tránh infinite wait
4. **WaitGroup** - sync goroutines (học ở mini project)

---

## 🎯 BƯỚC TIẾP THEO

- Đọc **Lesson 5** về Database
- Thử tạo goroutines trong mini project
- Thực hành channels & communication
