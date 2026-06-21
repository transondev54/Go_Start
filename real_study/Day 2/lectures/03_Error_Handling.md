# Lesson 3: Error Handling - Custom Errors & Error Wrapping

## 📖 Nội dung bài học

1. Go error model
2. Returning errors
3. Custom errors
4. Error wrapping (Go 1.13+)
5. Error handling patterns

---

## 1️⃣ GO ERROR MODEL

### Error là interface

```go
// Tính định nghĩa error từ standard library
type error interface {
    Error() string
}
```

### Convention

- ✅ Return `error` là return value cuối cùng
- ✅ Không panic trừ khi nghiêm trọng
- ✅ Return `nil` nếu thành công

---

## 2️⃣ RETURNING ERRORS

### Ví dụ cơ bản

```go
import \"errors\"

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New(\"division by zero\")
    }
    return a / b, nil
}

func main() {
    result, err := Divide(10, 2)
    if err != nil {
        fmt.Println(\"Error:\", err)
        return
    }
    fmt.Println(\"Result:\", result)
}
```

### fmt.Errorf - Format error messages

```go
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf(\"cannot divide %.2f by 0\", a)
    }
    return a / b, nil
}
```

---

## 3️⃣ CUSTOM ERRORS

### Định nghĩa custom error type

```go
type ValidationError struct {
    Field   string
    Message string
}

// Implement error interface
func (e ValidationError) Error() string {
    return fmt.Sprintf(\"validation error on %s: %s\", e.Field, e.Message)
}

func ValidateEmail(email string) error {
    if !strings.Contains(email, \"@\") {
        return ValidationError{
            Field:   \"email\",
            Message: \"invalid email format\",
        }
    }
    return nil
}

func main() {
    err := ValidateEmail(\"notanemail\")
    if err != nil {
        fmt.Println(err)  // validation error on email: invalid email format
    }
}
```

### Lợi ích custom errors

- ✅ Type-specific handling
- ✅ Thêm context
- ✅ Metadata cho debugging

---

## 4️⃣ ERROR WRAPPING (Go 1.13+)

### Wrapping errors

```go
import \"fmt\"

func ReadFile(filename string) ([]byte, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf(\"failed to read file %s: %w\", filename, err)
    }
    return data, nil
}

func main() {
    _, err := ReadFile(\"missing.txt\")
    if err != nil {
        fmt.Println(err)
        // Output: failed to read file missing.txt: open missing.txt: no such file or directory
    }
}
```

### Unwrap errors

```go
import \"errors\"

func ReadFile(filename string) ([]byte, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, fmt.Errorf(\"read %s: %w\", filename, err)
    }
    return data, nil
}

func main() {
    _, err := ReadFile(\"missing.txt\")

    // Check underlying error
    if errors.Is(err, os.ErrNotExist) {
        fmt.Println(\"File doesn't exist\")
    }

    // Get underlying error
    originalErr := errors.Unwrap(err)
    fmt.Println(\"Original error:\", originalErr)
}
```

---

## 5️⃣ ERROR HANDLING PATTERNS

### Pattern 1: Early Return

```go
func ProcessData(data string) error {
    if data == \"\" {
        return errors.New(\"data is empty\")
    }

    if err := ValidateData(data); err != nil {
        return fmt.Errorf(\"validation failed: %w\", err)
    }

    if err := SaveData(data); err != nil {
        return fmt.Errorf(\"save failed: %w\", err)
    }

    return nil
}
```

### Pattern 2: Custom Error Type

```go
type UserNotFoundError struct {
    UserID int
}

func (e UserNotFoundError) Error() string {
    return fmt.Sprintf(\"user %d not found\", e.UserID)
}

func GetUser(id int) (*User, error) {
    if id <= 0 {
        return nil, UserNotFoundError{UserID: id}
    }
    // ... fetch user
    return user, nil
}

func main() {
    _, err := GetUser(0)
    if _, ok := err.(UserNotFoundError); ok {
        fmt.Println(\"User not found\")
    }
}
```

### Pattern 3: Sentinel Errors

```go
var (
    ErrInvalidInput = errors.New(\"invalid input\")
    ErrNotFound     = errors.New(\"not found\")
    ErrUnauthorized = errors.New(\"unauthorized\")
)

func FindUser(id int) (*User, error) {
    if id <= 0 {
        return nil, ErrInvalidInput
    }
    // ...
}

func main() {
    _, err := FindUser(0)
    if err == ErrInvalidInput {
        fmt.Println(\"Invalid input provided\")
    }
}
```

---

## 📝 TÓM TẮT ERROR HANDLING

| Phương pháp         | Khi dùng                 |
| ------------------- | ------------------------ |
| **errors.New()**    | Lỗi đơn giản             |
| **fmt.Errorf()**    | Lỗi với format           |
| **Custom type**     | Lỗi phức tạp             |
| **Wrapping (%w)**   | Preserve original error  |
| **Sentinel errors** | Defined, reusable errors |

---

## 💡 BEST PRACTICES

1. **Luôn check error** - không bỏ qua
2. **Wrap with context** - thêm thông tin
3. **Use custom types** - khi cần type-specific handling
4. **Avoid panic** - chỉ dùng cho unrecoverable errors
5. **Log errors** - không silent failures

---

## 🎯 BƯỚC TIẾP THEO

- Đọc **Lesson 4** về Goroutines
- Thêm error handling vào mini project
- Tạo custom error types cho ứng dụng của bạn
