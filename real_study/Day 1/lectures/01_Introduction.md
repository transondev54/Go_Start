# Lesson 1: Introduction to Go

## 📖 Nội dung bài học

1. Go là gì?
2. Tại sao học Go?
3. Cài đặt Go
4. Kiểm tra cài đặt
5. First Go Program

---

## 1️⃣ GO LÀ GÌ?

### Định nghĩa

Go (hay Golang) là một **ngôn ngữ lập trình mã nguồn mở** được phát triển bởi **Google** từ năm 2007.

### Đặc điểm chính

- **Compiled language**: Biên dịch thành binary, chạy nhanh
- **Strongly typed**: Kiểu dữ liệu cứng nhắc (an toàn hơn)
- **Garbage collected**: Tự động quản lý bộ nhớ
- **Concurrent**: Hỗ trợ lập trình song song dễ dàng
- **Simple syntax**: Cú pháp đơn giản, dễ học

### Tại sao tên là Go?

- Ngắn gọn, dễ tìm trên Google
- Phát âm dễ (2 âm tiết)
- Tương tự như "Hello" trong lập trình

---

## 2️⃣ TẠI SAO HỌC GO?

### ✅ Ưu điểm

| Ưu điểm        | Chi tiết                                            |
| -------------- | --------------------------------------------------- |
| **Nhanh**      | Chạy nhanh như C/C++, nhưng code dễ viết như Python |
| **Đơn giản**   | Syntax gọn gàng, ít boilerplate code                |
| **Concurrent** | Goroutines làm lập trình song song dễ hơn           |
| **Popular**    | Được Google, Uber, Dropbox sử dụng                  |
| **Backend**    | Tuyệt vời cho backend, microservices                |
| **Cloud**      | Docker, Kubernetes viết bằng Go                     |

### 📊 So sánh với ngôn ngữ khác

```
          Speed    Simplicity   Concurrency
Go        ████████ ████████░░   ████████░░
Python    ██░░░░░░ ██████████   ██░░░░░░░░
Java      ████████ ██████░░░░   ████████░░
C++       ██████████ ██░░░░░░░  ██████░░░░
```

---

## 3️⃣ CÀI ĐẶT GO

### Bước 1: Download

- Vào https://golang.org/dl/
- Chọn phiên bản cho hệ điều hành của bạn
- Windows: `go1.x.x.windows-amd64.msi`

### Bước 2: Install

- **Windows**: Double-click file .msi, follow wizard
- **macOS**: Download .pkg hoặc dùng `brew install go`
- **Linux**: `sudo apt-get install golang-go` (Ubuntu/Debian)

### Bước 3: Kiểm tra cài đặt

Mở **Terminal/PowerShell** và gõ:

```bash
go version
```

**Output mong muốn:**

```
go version go1.21.0 windows/amd64
```

Nếu không nhận diện lệnh `go`, có nghĩa là cần restart terminal hoặc thêm vào PATH.

---

## 4️⃣ KIỂM TRA CÀI ĐẶT

### Kiểm tra $GOPATH

```bash
go env
```

Sẽ thấy danh sách các biến môi trường Go.

### Kiểm tra các lệnh Go

```bash
go help
```

Sẽ liệt kê các lệnh có sẵn:

- `go run` - Biên dịch và chạy
- `go build` - Biên dịch
- `go test` - Chạy test
- `go fmt` - Format code

---

## 5️⃣ FIRST GO PROGRAM

### Tạo thư mục dự án

```bash
mkdir hello_world
cd hello_world
```

### Tạo file `main.go`

```bash
code main.go  # hoặc dùng editor yêu thích
```

### Viết code Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Chạy chương trình

```bash
go run main.go
```

**Output:**

```
Hello, World!
```

---

## 📚 GIẢI THÍCH CODE

```go
package main                    // 1. Định nghĩa package

import "fmt"                    // 2. Import package fmt

func main() {                   // 3. Hàm chính (entry point)
    fmt.Println("Hello, World!") // 4. In ra màn hình
}
```

### Từng dòng:

1. **`package main`**: Mỗi file Go phải thuộc package. `main` là package đặc biệt (chứa hàm `main()`)

2. **`import "fmt"`**: Import package `fmt` (formatting & printing)

3. **`func main() { }`**: Hàm `main()` là entry point - nơi bắt đầu chương trình

4. **`fmt.Println(...)`**: In text với newline ở cuối

---

## 🎯 BÀI TẬP VỀ NHÀ

### Bài 1: Thay đổi message

Sửa `"Hello, World!"` thành tên của bạn, ví dụ: `"Xin chào, Ngọc!"`

### Bài 2: In nhiều dòng

```go
fmt.Println("Dòng 1")
fmt.Println("Dòng 2")
fmt.Println("Dòng 3")
```

### Bài 3: Khám phá

Thay `Println` bằng `Print` (không newline) và thấy khác gì:

```go
fmt.Print("Tôi ")
fmt.Print("yêu ")
fmt.Print("Go!")
```

---

## 💡 LƯU Ý

- Go là **case-sensitive**: `main` ≠ `Main`
- Package phải nằm trên cùng file
- Dấu `;` không bắt buộc (khác C++)
- Indentation không bắt buộc nhưng nên viết đẹp

---

## ❓ QUIZ (5 câu)

1. Go được phát triển bởi công ty nào?
2. Viết lệnh để cài đặt Go
3. Hàm nào là entry point của chương trình Go?
4. Package nào cung cấp hàm `Println`?
5. Sự khác nhau giữa `Print` và `Println`?

---

## 🎓 KẾT LUẬN

✅ Bạn đã nắm được:

- Tính chất của Go
- Cách cài đặt
- Viết chương trình đầu tiên

**Tiếp theo:** Lesson 2 - Hello World & Setup Project
