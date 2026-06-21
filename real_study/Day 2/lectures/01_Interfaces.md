# Lesson 1: Interfaces - Thiết kế Polymorphism

## 📖 Nội dung bài học

1. Interface là gì?
2. Định nghĩa interfaces
3. Implement interfaces (implicit)
4. Ví dụ thực tế
5. Empty interface

---

## 1️⃣ INTERFACE LÀ GÌ?

### Định nghĩa

**Interface** là một tập hợp các **method signatures** (chữ ký hàm). Nó định nghĩa hành vi mà một type phải thực hiện.

```go
// Định nghĩa interface
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### Đặc điểm

- ✅ Định nghĩa hợp đồng (contract)
- ✅ Cho phép polymorphism
- ✅ **Implicit implementation** - không cần khai báo `implements`
- ✅ Hỗ trợ composition tốt

---

## 2️⃣ ĐỊNH NGHĨA INTERFACES

### Cú pháp cơ bản

```go
type InterfaceName interface {
    Method1(param type) returnType
    Method2(param type) returnType
}
```

### Ví dụ

```go
// Định nghĩa interface Shape
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Định nghĩa interface Writer
type Writer interface {
    Write(data string) error
}
```

---

## 3️⃣ IMPLEMENT INTERFACES

### Cách implement

Một type **tự động implement** một interface nếu nó có tất cả các method của interface đó.

```go
// Rectangle type
type Rectangle struct {
    Width  float64
    Height float64
}

// Implement Shape interface
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Circle type
type Circle struct {
    Radius float64
}

// Implement Shape interface
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * 3.14 * c.Radius
}
```

### Sử dụng interface

```go
func PrintShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
    fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
    rect := Rectangle{Width: 5, Height: 3}
    circle := Circle{Radius: 4}

    PrintShapeInfo(rect)      // Polymorphism!
    PrintShapeInfo(circle)    // Cùng function, khác behavior
}
```

---

## 4️⃣ VÍ DỤ THỰC TẾ

### Database Interface

```go
// Define interface cho database operations
type Database interface {
    Query(sql string) ([]map[string]interface{}, error)
    Execute(sql string) error
    Close() error
}

// Implement MySQL
type MySQL struct {
    conn *sql.DB
}

func (m *MySQL) Query(sql string) ([]map[string]interface{}, error) {
    // Implementation
}

func (m *MySQL) Execute(sql string) error {
    // Implementation
}

func (m *MySQL) Close() error {
    return m.conn.Close()
}

// Implement SQLite
type SQLite struct {
    conn *sql.DB
}

func (s *SQLite) Query(sql string) ([]map[string]interface{}, error) {
    // Implementation
}

// ... rest of implementation

// Usage - không quan tâm database nào
func SaveData(db Database, data string) error {
    return db.Execute("INSERT INTO table VALUES (?)", data)
}
```

### Logger Interface

```go
type Logger interface {
    Log(message string)
    Error(message string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(message string) {
    fmt.Println("[LOG]", message)
}

func (cl ConsoleLogger) Error(message string) {
    fmt.Println("[ERROR]", message)
}

type FileLogger struct {
    file *os.File
}

func (fl FileLogger) Log(message string) {
    fl.file.WriteString("[LOG] " + message + "\n")
}

func (fl FileLogger) Error(message string) {
    fl.file.WriteString("[ERROR] " + message + "\n")
}

// Usage
func DoWork(logger Logger) {
    logger.Log("Starting work")
    // ... do something
    logger.Error("Error occurred")
}
```

---

## 5️⃣ EMPTY INTERFACE

### `interface{}` - Bất cứ cái gì

```go
// Empty interface chấp nhận bất cứ loại gì
func PrintAnything(val interface{}) {
    fmt.Println(val)
}

func main() {
    PrintAnything("hello")      // string
    PrintAnything(42)           // int
    PrintAnything(3.14)         // float
    PrintAnything([]int{1,2,3}) // slice
}
```

### Khi nào dùng

- ✅ Hàm generic (nhận nhiều loại)
- ✅ JSON parsing (unmarshaling)
- ✅ Type assertion cần thiết

---

## 📝 TÓM TẮT

| Khái niệm           | Mô tả                                  |
| ------------------- | -------------------------------------- |
| **Interface**       | Tập hợp method signatures              |
| **Implicit**        | Implement tự động khi có tất cả method |
| **Polymorphism**    | Cùng interface, behavior khác nhau     |
| **Empty interface** | `interface{}` - chấp nhận tất cả loại  |

---

## 💡 BEST PRACTICES

1. **Interfaces nhỏ** - định nghĩa interface với ít methods
2. **Reader, Writer** - sử dụng interfaces từ standard library
3. **Dependency injection** - pass interfaces, không concrete types
4. **Avoid empty interface** - sử dụng khi thực sự cần thiết

---

## 🎯 BƯỚC TIẾP THEO

- Đọc **Lesson 2** về Type Assertions
- Thử tạo interface cho mini project
- Thực hành implicit implementation
