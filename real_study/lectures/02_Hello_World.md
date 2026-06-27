# Lesson 2: Hello World & Project Setup

## 📖 Nội dung bài học

1. Tạo dự án Go đầu tiên
2. Cấu trúc thư mục
3. Go modules (go.mod)
4. Chạy & build chương trình
5. Sử dụng package khác nhau

---

## 1️⃣ TẠO DỰ ÁN GO

### Bước 1: Tạo thư mục dự án

```bash
mkdir my_first_app
cd my_first_app
```

### Bước 2: Khởi tạo module

```bash
go mod init my_first_app
```

**Kết quả:** Tạo file `go.mod`

```
module my_first_app

go 1.21
```

### Bước 3: Tạo file main.go

```bash
# Cách 1: Dùng editor
code main.go

# Cách 2: Command line
echo 'package main' > main.go
```

---

## 2️⃣ CẤU TRÚC THỪ MỤC

### Cấu trúc cơ bản

```
my_first_app/
├── go.mod              # Module definition
├── main.go             # Program entry point
└── README.md           # Documentation
```

### Cấu trúc lớn hơn

```
my_app/
├── go.mod
├── go.sum
├── main.go             # Entry point
├── README.md
├── cmd/                # Executable programs
│   └── cli/
│       └── main.go
├── pkg/                # Reusable packages
│   ├── utils/
│   │   └── utils.go
│   └── config/
│       └── config.go
└── internal/           # Internal packages
    └── database/
        └── db.go
```

### Quy tắc đặt tên:

- **`main.go`**: Entry point (phải chứa hàm `main()`)
- **`cmd/`**: Các executable programs
- **`pkg/`**: Các package cái reusable
- **`internal/`**: Các package nội bộ (không exported)

---

## 3️⃣ GO MODULES (go.mod)

### Tác dụng

- Quản lý version dependencies
- Giống như `package.json` (Node) hay `requirements.txt` (Python)

### File `go.mod` example

```
module github.com/username/my_app

go 1.21

require (
    github.com/some/package v1.2.3
    github.com/other/lib v2.0.0
)
```

### Lệnh quản lý modules

```bash
go mod init <module-name>      # Khởi tạo module
go mod tidy                    # Cập nhật dependencies
go mod download                # Download dependencies
go mod graph                   # Xem dependency graph
```

---

## 4️⃣ CHẠY & BUILD CHƯƠNG TRÌNH

### Chạy trực tiếp (không tạo file)

```bash
go run main.go
```

**Ưu điểm:** Nhanh, tiện cho development
**Nhược điểm:** Cần Go environment để chạy

### Build thành executable

```bash
go build
```

**Kết quả:**

- Windows: `my_first_app.exe`
- Linux/macOS: `my_first_app`

**Chạy file executable:**

```bash
./my_first_app      # Linux/macOS
my_first_app.exe    # Windows (hoặc .\my_first_app.exe)
```

### Build cho hệ điều hành khác

```bash
# Build cho Linux từ Windows
GOOS=linux GOARCH=amd64 go build

# Build cho macOS từ Windows
GOOS=darwin GOARCH=amd64 go build

# Build cho Windows từ Linux
GOOS=windows GOARCH=amd64 go build
```

---

## 5️⃣ SỬ DỤNG PACKAGE KHÁC NHAU

### Import một package

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
}
```

### Import nhiều package

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Pi =", math.Pi)
}
```

### Import package tùy chỉnh (custom)

**Cấu trúc:**

```
my_app/
├── go.mod
├── main.go
└── utils/
    └── helper.go
```

**File: `utils/helper.go`**

```go
package utils

func Greet(name string) string {
    return "Hello, " + name
}
```

**File: `main.go`**

```go
package main

import (
    "fmt"
    "my_app/utils"
)

func main() {
    msg := utils.Greet("Alice")
    fmt.Println(msg)
}
<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

```

### Import external packages

**Cài đặt từ GitHub:**

```bash
go get github.com/some/awesome-package
```

**Dùng trong code:**

```go
import "github.com/some/awesome-package"
```

---

## 📝 COMPLETE EXAMPLE

### main.go

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Biến
    name := "Ngọc"
    age := 25

    // In ra màn hình
    fmt.Println("=== Welcome to Go ===")
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)

    // Tính toán
    radius := 5.0
    area := math.Pi * radius * radius
    fmt.Printf("Diện tích hình tròn (r=%.1f): %.2f\n", radius, area)

    // Hết
    fmt.Println("=== End of Program ===")
}
```

**Chạy:**

```bash
go run main.go
```

**Output:**

```
=== Welcome to Go ===
Name: Ngọc
Age: 25
Diện tích hình tròn (r=5.0): 78.54
=== End of Program ===
```

---

## 🎯 BÀI TẬP

### Bài 1: Khởi tạo project

```bash
mkdir hello_project
cd hello_project
go mod init hello_project
```

### Bài 2: Tạo main.go với:

- In tên của bạn
- In năm sinh
- In năm hiện tại (2024)
- In tuổi (2024 - năm sinh)

**Gợi ý:**

```go
package main
import "fmt"

func main() {
    name := "Your Name"
    birthYear := 2000
    // ...
}
```

### Bài 3: Build & chạy executable

```bash
go build
./hello_project  # hoặc hello_project.exe
```

---

## 💡 LƯU Ý QUAN TRỌNG

1. **Package name**: Phải là `main` để chạy được
2. **Module name**: Có thể khác tên thư mục
3. **Naming convention**:
   - Package: lowercase, no underscores
   - Functions: PascalCase (exported), camelCase (unexported)
4. **Import**: Không dùng dấu `*` như Python

---

## 🐛 TROUBLESHOOTING

| Error                     | Nguyên nhân        | Giải pháp           |
| ------------------------- | ------------------ | ------------------- |
| `go: command not found`   | Go chưa cài        | Cài đặt Go          |
| `undefined: fmt`          | Quên import        | Thêm `import "fmt"` |
| `cannot find main module` | Quên `go mod init` | Chạy `go mod init`  |
| `no Go files in /path`    | Không có .go files | Tạo file main.go    |

---

## 📚 REFERENCES

- Go Documentation: https://golang.org/doc
- Go Tour: https://tour.golang.org
- Go Packages: https://pkg.go.dev

---

## ✅ CHECKLIST

- [ ] Cài đặt Go (Lesson 1)
- [ ] Tạo thư mục dự án
- [ ] Khởi tạo module (go mod init)
- [ ] Tạo main.go
- [ ] Chạy với `go run`
- [ ] Build executable
- [ ] Chạy executable

---

**Tiếp theo:** Lesson 3 - Variables and Data Types
