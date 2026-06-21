# Lesson 5: Control Flow (If, For, Switch)

## 📖 Nội dung bài học

1. If/Else statements
2. For loops
3. Switch statements
4. Break & Continue
5. Nested loops & conditions

---

## 1️⃣ IF/ELSE STATEMENTS

### If cơ bản

```go
if age >= 18 {
    fmt.Println("Adult")
}
```

### If/Else

```go
if age >= 18 {
    fmt.Println("Adult")
} else {
    fmt.Println("Minor")
}
```

### If/Else If/Else

```go
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else if score >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("F")
}
```

### If với initialization

```go
if x := 5; x > 3 {
    fmt.Println("x > 3")
}
// x không dùng được ngoài block này
```

### Comparison operators

```go
x == y      // Equal
x != y      // Not equal
x < y       // Less than
x > y       // Greater than
x <= y      // Less or equal
x >= y      // Greater or equal
```

### Logical operators

```go
x > 0 && x < 10     // AND
x < 0 || x > 10     // OR
!valid              // NOT
```

---

## 2️⃣ FOR LOOPS

### For cơ bản (C-style)

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)  // 0, 1, 2, 3, 4
}
```

**Phần tử:**

- **i := 0**: Khởi tạo
- **i < 5**: Điều kiện
- **i++**: Increment

### For với range

```go
arr := []int{10, 20, 30}

// Cả index và value
for i, v := range arr {
    fmt.Printf("Index %d: %v\n", i, v)
}

// Chỉ value
for _, v := range arr {
    fmt.Println(v)
}

// Chỉ index
for i := range arr {
    fmt.Println(i)
}
```

### While loop (dùng for)

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

### Infinite loop

```go
for {
    fmt.Println("Forever!")
    // Cần break để thoát
}
```

### Break & Continue

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break       // Thoát vòng lặp
    }
    fmt.Println(i)  // 0, 1, 2, 3, 4
}

for i := 0; i < 10; i++ {
    if i == 5 {
        continue    // Bỏ qua iteration này
    }
    fmt.Println(i)  // 0, 1, 2, 3, 4, 6, 7, 8, 9
}
```

---

## 3️⃣ SWITCH STATEMENTS

### Switch cơ bản

```go
day := 3

switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
default:
    fmt.Println("Unknown day")
}
// Output: Wednesday
```

### Switch với multiple cases

```go
switch day {
case 1, 2, 3, 4, 5:
    fmt.Println("Weekday")
case 6, 7:
    fmt.Println("Weekend")
}
```

### Switch với fallthrough

```go
x := 2

switch x {
case 1:
    fmt.Println("One")
    fallthrough  // Chạy tiếp case tiếp theo
case 2:
    fmt.Println("One or Two")
    fallthrough
case 3:
    fmt.Println("One, Two, or Three")
default:
    fmt.Println("Other")
}
// Output:
// One or Two
// One, Two, or Three
```

### Switch với expression

```go
score := 85

switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 70:
    fmt.Println("C")
default:
    fmt.Println("F")
}
// Output: B
```

### Switch với initialization

```go
switch x := getValue(); x {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
}
```

---

## 📝 COMPLETE EXAMPLE

```go
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // ===== IF/ELSE =====
    fmt.Println("=== IF/ELSE ===")
    age := 20
    if age >= 18 {
        fmt.Println("Bạn là người lớn")
    } else {
        fmt.Println("Bạn còn nhỏ")
    }

    // ===== FOR LOOP =====
    fmt.Println("\n=== FOR LOOP ===")
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d ", i)
    }
    fmt.Println()

    // ===== RANGE =====
    fmt.Println("\n=== RANGE ===")
    fruits := []string{"Apple", "Banana", "Orange"}
    for i, fruit := range fruits {
        fmt.Printf("%d: %s\n", i, fruit)
    }

    // ===== SWITCH =====
    fmt.Println("\n=== SWITCH ===")
    day := 3
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Unknown")
    }

    // ===== BREAK & CONTINUE =====
    fmt.Println("\n=== BREAK & CONTINUE ===")
    for i := 0; i < 10; i++ {
        if i == 5 {
            continue
        }
        if i == 8 {
            break
        }
        fmt.Printf("%d ", i)
    }
    fmt.Println()
}
```

---

## 🎯 BÀI TẬP

### Bài 1: Kiểm tra số

```go
num := 15

if num < 0 {
    fmt.Println("Negative")
} else if num == 0 {
    fmt.Println("Zero")
} else {
    fmt.Println("Positive")
}
```

### Bài 2: In bảng cửu chương

```go
// In bảng 5
for i := 1; i <= 10; i++ {
    fmt.Printf("5 x %d = %d\n", i, 5*i)
}
```

### Bài 3: Hệ thống rating

```go
rating := 4.5

switch {
case rating >= 4.5:
    fmt.Println("Excellent")
case rating >= 4.0:
    fmt.Println("Very Good")
case rating >= 3.0:
    fmt.Println("Good")
default:
    fmt.Println("Poor")
}
```

### Bài 4: Tìm số trong slice

```go
numbers := []int{1, 5, 3, 8, 2, 9}
target := 8

found := false
for _, num := range numbers {
    if num == target {
        found = true
        break
    }
}

if found {
    fmt.Println("Found!")
} else {
    fmt.Println("Not found!")
}
```

---

## 💡 LƯU Ý

1. **Braces required**: Không tùy chọn `{}`
2. **No parentheses**: Không cần `if (x > 5)` như C
3. **Default execution**: Switch không cần `default`
4. **For is everything**: Go không có `while`, chỉ có `for`
5. **Range loop**: Không có order với maps

---

## ✅ CHECKLIST

- [ ] Viết được if/else
- [ ] Viết được for loops
- [ ] Dùng được range loop
- [ ] Viết được switch statements
- [ ] Dùng được break & continue

---

**Tiếp theo:** Lesson 6 - Functions
