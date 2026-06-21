# Lesson 7: Testing & Code Quality

## 📖 Nội dung bài học

1. Unit testing cơ bản
2. Table-driven tests
3. Test utilities
4. Benchmarking
5. Coverage

---

## 1️⃣ UNIT TESTING CƠ BẢN

### Viết test

```go
// calculator.go
func Add(a, b int) int {
    return a + b
}

func Subtract(a, b int) int {
    return a - b
}
```

```go
// calculator_test.go
package main

import \"testing\"

// Tên phải là Test[FunctionName]
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5

    if result != expected {
        t.Errorf(\"Add(2, 3) = %d, want %d\", result, expected)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    expected := 2

    if result != expected {
        t.Errorf(\"Subtract(5, 3) = %d, want %d\", result, expected)
    }
}
```

### Chạy tests

```bash
go test              # Chạy tất cả tests
go test -v           # Verbose output
go test -run TestAdd # Chạy specific test
```

---

## 2️⃣ TABLE-DRIVEN TESTS

### Ví dụ

```go
func TestAddMultiple(t *testing.T) {
    tests := []struct {
        a        int
        b        int
        expected int
    }{
        {2, 3, 5},
        {-1, 1, 0},
        {0, 0, 0},
        {100, -50, 50},
    }

    for _, test := range tests {
        result := Add(test.a, test.b)
        if result != test.expected {
            t.Errorf(\"Add(%d, %d) = %d, want %d\",
                test.a, test.b, result, test.expected)
        }
    }
}
```

### Với names

```go
func TestAddWithNames(t *testing.T) {
    tests := []struct {
        name     string
        a        int
        b        int
        expected int
    }{
        {\"positive numbers\", 2, 3, 5},
        {\"negative numbers\", -1, -1, -2},
        {\"mixed signs\", 10, -5, 5},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := Add(test.a, test.b)
            if result != test.expected {
                t.Errorf(\"got %d, want %d\", result, test.expected)
            }
        })
    }
}
```

---

## 3️⃣ TESTING FUNCTIONS WITH ERRORS

```go
func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf(\"division by zero\")
    }
    return a / b, nil
}

func TestDivideError(t *testing.T) {
    _, err := Divide(10, 0)
    if err == nil {
        t.Error(\"Expected error for division by zero\")
    }
}

func TestDivideSuccess(t *testing.T) {
    result, err := Divide(10, 2)
    if err != nil {
        t.Fatalf(\"Unexpected error: %v\", err)
    }

    if result != 5 {
        t.Errorf(\"got %f, want 5\", result)
    }
}
```

---

## 4️⃣ TESTING WITH STRUCTS

```go
type User struct {
    Name  string
    Email string
}

func NewUser(name, email string) (*User, error) {
    if name == \"\" {
        return nil, fmt.Errorf(\"name required\")
    }
    if !strings.Contains(email, \"@\") {
        return nil, fmt.Errorf(\"invalid email\")
    }
    return &User{name, email}, nil
}

func TestNewUserValid(t *testing.T) {
    user, err := NewUser(\"John\", \"john@example.com\")
    if err != nil {
        t.Fatalf(\"unexpected error: %v\", err)
    }

    if user.Name != \"John\" {
        t.Errorf(\"got %q, want John\", user.Name)
    }
}

func TestNewUserInvalid(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        wantErr bool
    }{
        {\"\", \"test@example.com\", true},
        {\"John\", \"invalid-email\", true},
        {\"John\", \"john@example.com\", false},
    }

    for _, test := range tests {
        _, err := NewUser(test.name, test.email)
        if (err != nil) != test.wantErr {
            t.Errorf(\"NewUser(%q, %q): got error %v, want error: %v\",
                test.name, test.email, err, test.wantErr)
        }
    }
}
```

---

## 5️⃣ BENCHMARKING

### Viết benchmarks

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

func BenchmarkAddLarge(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1000000, 2000000)
    }
}
```

### Chạy benchmarks

```bash
go test -bench=.         # Chạy tất cả benchmarks
go test -bench=Add       # Chạy specific
go test -benchmem=.      # Memory allocation info
```

---

## 6️⃣ COVERAGE

### Measure coverage

```bash
go test -cover                    # Percentage
go test -coverprofile=coverage.out
go tool cover -html=coverage.out  # View in browser
```

---

## 📝 TEST CHECKLIST

- ✅ Happy path - normal inputs
- ✅ Edge cases - empty, zero, nil
- ✅ Error cases - invalid inputs
- ✅ Boundary conditions
- ✅ Table-driven tests

---

## 💡 BEST PRACTICES

1. **Test names rõ ràng** - `TestAddPositiveNumbers`
2. **Arrange-Act-Assert** - Setup → Execute → Verify
3. **One assertion per test** - nếu có thể
4. **Use subtests** - `t.Run()` cho organization
5. **Aim for 80%+ coverage** - không 100% cũng OK

---

## 🎯 BƯỚC TIẾP THEO

- Viết tests cho mini project của bạn
- Đạt coverage ít nhất 70%
- Review code quality
