# Lesson 3: Variables and Data Types

## 📖 Nội dung bài học

1. Khai báo biến
2. Kiểu dữ liệu cơ bản
3. Zero values
4. Type conversion
5. Constants

---

## 1️⃣ KHAI BÁO BIẾN

### Cách 1: Var keyword (dài dòng)

```go
var name string
var age int
var height float64

name = "Alice"
age = 25
height = 1.65
```

### Cách 2: Var với initialization

```go
var name string = "Bob"
var age int = 30
var height float64 = 1.75
```

### Cách 3: Type inference (Go tự suy kiểu)

```go
var name = "Charlie"      // Go biết là string
var age = 35              // Go biết là int
var height = 1.85         // Go biết là float64
```

### Cách 4: Short declaration `:=` (phổ biến nhất)

```go
name := "David"
age := 40
height := 1.90
```

**Lưu ý:** Chỉ dùng trong hàm, không dùng ở package level

### Cách 5: Multiple variables

```go
// Cách 1
var a, b, c int
a, b, c = 1, 2, 3

// Cách 2
x, y, z := 10, 20, 30

// Cách 3
var (
    name = "Eve"
    age = 28
    city = "Hanoi"
)
```

### Swap giá trị

```go
x, y := 5, 10
x, y = y, x  // x=10, y=5
```

---

## 2️⃣ KIỂU DỮ LIỆU CƠ BẢN

### A. Numbers - Integers

```go
var i int = 42              // Platform-dependent (32 or 64 bit)
var i8 int8 = 127           // -128 to 127
var i16 int16 = 32767       // -32,768 to 32,767
var i32 int32 = 2147483647  // ~2 billion
var i64 int64 = 9223372036854775807  // ~9 quintillion

// Unsigned (không âm)
var u uint = 42
var u8 uint8 = 255
var u16 uint16 = 65535
var u32 uint32 = 4294967295
var u64 uint64 = 18446744073709551615

// Aliases
byte  // = uint8
rune  // = int32 (dùng cho Unicode)
```

**Khi nào dùng cái nào:**

- `int` / `uint`: Dùng mặc định
- `int64`: Làm việc với database, timestamps
- `uint8`: Làm việc với bytes/binary data

### B. Numbers - Floating Point

```go
var f32 float32 = 3.14
var f64 float64 = 3.14159265359

// Special values
inf := math.Inf(1)     // +Infinity
negInf := math.Inf(-1) // -Infinity
nan := math.NaN()      // Not-a-Number
```

**Khác biệt:**

- `float32`: 7 chữ số thập phân chính xác
- `float64`: 15-17 chữ số (dùng mặc định)

### C. Numbers - Complex

```go
var c1 complex64 = 1 + 2i
var c2 complex128 = 3 + 4i

// Hàm
real(c1)  // 1
imag(c1)  // 2
```

### D. Strings

```go
var name string = "Alice"
var empty string = ""

// String literals
raw := `Line 1
Line 2\n is literal`

// Escape sequences
escaped := "Line 1\nLine 2\tTabbed"
```

**Cặp ký tự:**

- `"..."`: Cho escape sequences
- `` `...` ``: Raw string (không escape)

**Thao tác với string:**

```go
s := "Hello"
len(s)      // 5
s[0]        // 'H' (byte)
s[1:4]      // "ell" (substring)
```

### E. Boolean

```go
var active bool = true
var inactive bool = false
```

**Operators:**

```go
true && false   // false (AND)
true || false   // true (OR)
!true           // false (NOT)
```

---

## 3️⃣ ZERO VALUES

Nếu không khởi tạo, Go gán giá trị mặc định (zero value):

```go
var i int       // 0
var f float64   // 0.0
var s string    // "" (empty string)
var b bool      // false
var p *int      // nil (null pointer)
```

---

## 📝 COMPLETE EXAMPLE

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Integers
    age := 25
    count := int64(1000000)

    // Floats
    pi := math.Pi
    height := 1.75

    // String
    name := "Ngọc"

    // Boolean
    isStudent := true

    // Print
    fmt.Println("=== Variables Demo ===")
    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Height: %.2f m\n", height)
    fmt.Printf("Is Student: %v\n", isStudent)
    fmt.Printf("Pi: %.4f\n", pi)
    fmt.Printf("Count: %d\n", count)
}
```

**Output:**

```
=== Variables Demo ===
Name: Ngọc
Age: 25
Height: 1.75 m
Is Student: true
Pi: 3.1416
Count: 1000000
```

---

## 4️⃣ TYPE CONVERSION

Go **không** cho phép implicit type conversion - phải explicit:

```go
var i int = 42
var f float64

// WRONG: f = i  // Compilation error

// CORRECT: Convert to float64
f = float64(i)
fmt.Println(f)  // 42

// String to int
s := "100"
n, err := strconv.Atoi(s)  // "100" -> 100

// Int to string
i := 42
s := strconv.Itoa(i)  // 42 -> "42"

// Using fmt.Sprint
i := 42
s := fmt.Sprint(i)  // 42 -> "42"
```

---

## 5️⃣ CONSTANTS

### Khai báo const

```go
const Pi = 3.14159
const MaxRetry = 5
const Message = "Hello, World!"

// Multiple constants
const (
    KB = 1000
    MB = 1000 * KB
    GB = 1000 * MB
)
```

### Untyped vs Typed constants

```go
// Untyped (linh hoạt hơn)
const x = 42
var i int = x      // OK
var f float64 = x  // OK

// Typed (cứng nhắc hơn)
const y int = 42
var i int = y      // OK
// var f float64 = y  // Error: cannot use y (int) as float64
```

### Iota (enumeration)

```go
const (
    Sunday = iota    // 0
    Monday           // 1
    Tuesday          // 2
    Wednesday        // 3
)

const (
    Kilobyte = 1 << (10 * iota)  // 1024
    Megabyte                      // 1048576
    Gigabyte                      // 1073741824
)
```

---

## 🎯 BÀI TẬP

### Bài 1: Khai báo biến

Tạo các biến sau:

- Tên (string)
- Tuổi (int)
- Chiều cao (float64)
- Học sinh hay không (bool)

```go
package main
import "fmt"

func main() {
    // TODO: Khai báo biến ở đây

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
}
```

### Bài 2: Tính BMI

- Input: cân nặng (kg), chiều cao (m)
- Formula: BMI = cân nặng / (chiều cao²)
- Output: In ra BMI

```go
weight := 70.0
height := 1.75
bmi := weight / (height * height)
fmt.Printf("BMI: %.2f\n", bmi)
```

### Bài 3: Type conversion

```go
s := "100"
// Convert string "100" to int 100
// Add 50 to it
// Convert back to string
// Print result
```

### Bài 4: Constants

```go
const (
    Red = iota
    Green
    Blue
)

fmt.Println(Red, Green, Blue)  // Output: 0 1 2
```

---

## 💡 LƯU Ý QUAN TRỌNG

1. **Type matters**: Go là strongly typed - phải match type
2. **Short declaration** `:=` chỉ dùng trong hàm
3. **Conversion explicit**: Phải viết `float64(i)`, không tự động
4. **String indexing**: `s[i]` trả về `byte`, không rune
5. **Blank identifier**: `_` để bỏ qua giá trị

---

## 📊 FORMAT STRING (fmt.Printf)

```go
fmt.Printf("%d\n", 42)          // 42 (int)
fmt.Printf("%s\n", "Hello")     // Hello (string)
fmt.Printf("%f\n", 3.14)        // 3.140000 (float)
fmt.Printf("%.2f\n", 3.14)      // 3.14 (2 decimal places)
fmt.Printf("%t\n", true)        // true (bool)
fmt.Printf("%v\n", 42)          // 42 (any - auto format)
fmt.Printf("%T\n", 42)          // int (type)
```

---

## ✅ CHECKLIST

- [ ] Hiểu 5 cách khai báo biến
- [ ] Biết các kiểu dữ liệu cơ bản
- [ ] Biết zero values
- [ ] Có thể convert giữa các type
- [ ] Hiểu constants và iota

---

**Tiếp theo:** Lesson 4 - Collections (Arrays, Slices, Maps)
