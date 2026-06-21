# Lesson 6: Functions

## 📖 Nội dung bài học

1. Function declaration
2. Parameters & return values
3. Multiple returns
4. Named return values
5. Variadic functions
6. Anonymous functions
7. Closures

---

## 1️⃣ FUNCTION DECLARATION

### Cú pháp cơ bản

```go
func greet(name string) {
    fmt.Println("Hello,", name)
}

// Gọi hàm
greet("Alice")
```

### Hàm có return value

```go
func add(x int, y int) int {
    return x + y
}

// Gọi
result := add(5, 3)
fmt.Println(result)  // 8
```

### Hàm với cùng type parameters

```go
// Cách 1 (verbose)
func sum(x int, y int, z int) int {
    return x + y + z
}

// Cách 2 (compact)
func sum(x, y, z int) int {
    return x + y + z
}
```

---

## 2️⃣ PARAMETERS & RETURN VALUES

### Single return

```go
func getName() string {
    return "Alice"
}

name := getName()
```

### Parameter passing

```go
// Pass by value (copy)
func modifyValue(x int) {
    x = 100  // Không ảnh hưởng biến ngoài
}

// Pass by reference (pointer)
func modifyPointer(x *int) {
    *x = 100  // Ảnh hưởng biến ngoài
}

// Sử dụng
value := 5
modifyValue(value)
fmt.Println(value)  // 5 (không thay đổi)

modifyPointer(&value)
fmt.Println(value)  // 100 (thay đổi)
```

---

## 3️⃣ MULTIPLE RETURN VALUES

### Return 2+ values

```go
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
}

// Gọi
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Result:", result)
}
```

### Ignore return values

```go
result, _ := divide(10, 2)
fmt.Println(result)
```

### Common pattern: (value, error)

```go
func getValue() (string, error) {
    // ... logic
    if /* error */ {
        return "", errors.New("failed")
    }
    return "value", nil
}

// Usage
value, err := getValue()
if err != nil {
    // Handle error
}
```

---

## 4️⃣ NAMED RETURN VALUES

### Named returns

```go
// Không dùng named
func getCoordinates() (float64, float64) {
    return 3.5, 4.5
}

// Dùng named returns
func getCoordinates() (x, y float64) {
    x = 3.5
    y = 4.5
    return
}

// Gọi
x, y := getCoordinates()
fmt.Printf("X: %.1f, Y: %.1f\n", x, y)
```

### Bare return

```go
func calculate(x, y int) (sum, product int) {
    sum = x + y
    product = x * y
    return  // Trả về sum & product
}

s, p := calculate(5, 3)
fmt.Println(s, p)  // 8 15
```

---

## 5️⃣ VARIADIC FUNCTIONS

### Hàm với số lượng tham số thay đổi

```go
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Gọi
fmt.Println(sum(1, 2, 3))        // 6
fmt.Println(sum(1, 2, 3, 4, 5))  // 15
fmt.Println(sum())               // 0
```

### Unpack slice

```go
numbers := []int{1, 2, 3, 4}
fmt.Println(sum(numbers...))  // 10
```

### Variadic with fixed parameters

```go
func format(prefix string, values ...string) string {
    result := prefix + ": "
    for i, v := range values {
        if i > 0 {
            result += ", "
        }
        result += v
    }
    return result
}

fmt.Println(format("Names", "Alice", "Bob", "Charlie"))
```

---

## 6️⃣ ANONYMOUS FUNCTIONS

### Function value

```go
// Assign function to variable
add := func(x, y int) int {
    return x + y
}

result := add(5, 3)
fmt.Println(result)  // 8
```

### Call immediately

```go
result := func(x, y int) int {
    return x + y
}(5, 3)

fmt.Println(result)  // 8
```

---

## 7️⃣ CLOSURES

### Function returning function

```go
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

// Gọi
double := makeMultiplier(2)
triple := makeMultiplier(3)

fmt.Println(double(5))  // 10
fmt.Println(triple(5))  // 15
```

### Closure capturing variables

```go
counter := 0

increment := func() {
    counter++
}

increment()
fmt.Println(counter)  // 1

increment()
fmt.Println(counter)  // 2
```

---

## 📝 COMPLETE EXAMPLE

```go
package main

import (
    "fmt"
    "math"
)

// 1. Hàm cơ bản
func greet(name string) {
    fmt.Println("Hello,", name)
}

// 2. Hàm có return
func add(x, y int) int {
    return x + y
}

// 3. Multiple returns
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return x / y, nil
}

// 4. Named returns
func getCircleStats(radius float64) (area, circumference float64) {
    area = math.Pi * radius * radius
    circumference = 2 * math.Pi * radius
    return
}

// 5. Variadic function
func average(numbers ...float64) float64 {
    if len(numbers) == 0 {
        return 0
    }
    sum := 0.0
    for _, n := range numbers {
        sum += n
    }
    return sum / float64(len(numbers))
}

func main() {
    greet("Alice")
    fmt.Println("Sum:", add(5, 3))

    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division:", result)
    }

    area, circ := getCircleStats(5)
    fmt.Printf("Circle (r=5): Area=%.2f, Circumference=%.2f\n", area, circ)

    fmt.Println("Average:", average(1, 2, 3, 4, 5))
}
```

---

## 🎯 BÀI TẬP

### Bài 1: Hàm tính BMI

```go
func calculateBMI(weight, height float64) string {
    bmi := weight / (height * height)
    if bmi < 18.5 {
        return "Underweight"
    } else if bmi < 25 {
        return "Normal"
    } else {
        return "Overweight"
    }
}

fmt.Println(calculateBMI(70, 1.75))
```

### Bài 2: Hàm tìm min/max

```go
func findMinMax(numbers ...int) (min, max int) {
    if len(numbers) == 0 {
        return 0, 0
    }
    min, max = numbers[0], numbers[0]
    for _, n := range numbers {
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }
    return
}

min, max := findMinMax(3, 1, 4, 1, 5, 9, 2, 6)
fmt.Printf("Min: %d, Max: %d\n", min, max)
```

### Bài 3: Hàm factorial

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n - 1)
}

fmt.Println("5! =", factorial(5))
```

### Bài 4: Hàm tạo multiplier

```go
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

double := makeMultiplier(2)
triple := makeMultiplier(3)

fmt.Println("2*5 =", double(5))
fmt.Println("3*5 =", triple(5))
```

---

## 💡 LƯU Ý

1. **Named parameters**: Ngôn ngữ Go không có default parameters
2. **Error handling**: Go convention - return error as last value
3. **Bare return**: Dùng khi return values đã named
4. **Variadic**: `...Type` phải là tham số cuối cùng
5. **Closures**: Cẩn thận với variable capture

---

## ✅ CHECKLIST

- [ ] Viết được hàm cơ bản
- [ ] Dùng được multiple returns
- [ ] Dùng được named returns
- [ ] Viết được variadic functions
- [ ] Hiểu closures

---

**Tiếp theo:** Lesson 7 - Structs & Methods
