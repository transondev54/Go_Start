# Lesson 4: Collections (Arrays, Slices, Maps)

## 📖 Nội dung bài học

1. Arrays - Mảng cố định
2. Slices - Mảng động
3. Maps - Từ điển/Hash table
4. Iterating collections

---

## 1️⃣ ARRAYS - MẢNG CỐ ĐỊNH

### Khai báo array

```go
// Cách 1: Chỉ định length
var numbers [5]int

// Cách 2: Khởi tạo với giá trị
var fruits [3]string = [3]string{"Apple", "Banana", "Orange"}

// Cách 3: Length inference (tự đếm)
names := [...]string{"Alice", "Bob", "Charlie"}  // length = 3

// Cách 4: Partial initialization
scores := [5]int{10, 20}  // [10, 20, 0, 0, 0]
```

### Truy cập & sửa

```go
arr := [5]int{1, 2, 3, 4, 5}

// Truy cập
fmt.Println(arr[0])      // 1
fmt.Println(arr[4])      // 5

// Sửa giá trị
arr[2] = 100
fmt.Println(arr[2])      // 100

// Độ dài
fmt.Println(len(arr))    // 5

// Range loop
for i, v := range arr {
    fmt.Printf("Index %d: %v\n", i, v)
}
```

### Đặc điểm array

- **Fixed length**: Khi khai báo, length không thay đổi
- **Same type**: Tất cả phần tử cùng kiểu
- **Index from 0**: Bắt đầu từ 0
- **Zero values**: Uninitialized elements = 0/""/false

---

## 2️⃣ SLICES - MẢNG ĐỘNG

### Khai báo slice

```go
// Cách 1: Tương tự array nhưng không chỉ length
var numbers []int

// Cách 2: Khởi tạo
fruits := []string{"Apple", "Banana"}

// Cách 3: Make function
numbers := make([]int, 5)           // length=5, capacity=5
numbers := make([]int, 5, 10)       // length=5, capacity=10

// Cách 4: Slice từ array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // [2, 3, 4] (index 1 đến 3)
```

### Slice operations

```go
s := []int{1, 2, 3, 4, 5}

// Độ dài & capacity
len(s)    // 5 (phần tử hiện có)
cap(s)    // 5 (có thể chứa được bao nhiêu)

// Append (thêm phần tử)
s = append(s, 6)        // [1, 2, 3, 4, 5, 6]
s = append(s, 7, 8, 9)  // [1, 2, 3, 4, 5, 6, 7, 8, 9]

// Copy
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)  // dst = [1, 2, 3]

// Slicing
s := []int{1, 2, 3, 4, 5}
s[1:3]    // [2, 3]
s[:3]     // [1, 2, 3]
s[2:]     // [3, 4, 5]
s[:]      // [1, 2, 3, 4, 5]
```

### Nil slice vs Empty slice

```go
// Nil slice
var nilSlice []int
nilSlice == nil           // true
len(nilSlice)             // 0
cap(nilSlice)             // 0

// Empty slice
emptySlice := []int{}
emptySlice == nil         // false
len(emptySlice)           // 0

// Append to nil slice
nilSlice = append(nilSlice, 1)  // Works fine
```

### Slice internals

Slice chứa 3 thành phần:

- **Pointer**: Trỏ tới array bên dưới
- **Length**: Số phần tử hiện có
- **Capacity**: Kích thước array bên dưới

```go
s := []int{1, 2, 3, 4, 5}
s = s[1:3]      // Points to [2,3], len=2, cap=4
```

---

## 3️⃣ MAPS - TỪ ĐIỂN/HASH TABLE

### Khai báo map

```go
// Cách 1
var ages map[string]int

// Cách 2: Make (cần phải make)
ages := make(map[string]int)

// Cách 3: Khởi tạo với giá trị
person := map[string]string{
    "name": "Alice",
    "city": "Hanoi",
}
```

### Map operations

```go
ages := make(map[string]int)

// Add/Update
ages["Alice"] = 25
ages["Bob"] = 30

// Read
fmt.Println(ages["Alice"])  // 25

// Delete
delete(ages, "Bob")

// Check existence
age, ok := ages["Alice"]
if ok {
    fmt.Println("Age:", age)  // Age: 25
}

// Iteration
for name, age := range ages {
    fmt.Printf("%s: %d\n", name, age)
}

// Get all keys
for name := range ages {
    fmt.Println(name)
}
```

### Lưu ý về map

- **Nil map**: Không thể thêm giá trị - phải dùng `make()`
- **Read missing key**: Trả về zero value, không error
- **Key types**: Phải comparable (int, string, bool, v.v)
- **Value types**: Có thể là bất kỳ type

```go
// WRONG
var m map[string]int
m["key"] = 1  // panic: assignment to entry in nil map

// RIGHT
m := make(map[string]int)
m["key"] = 1  // OK
```

---

## 4️⃣ ITERATING COLLECTIONS

### For loop với range

```go
// Array
arr := [3]string{"a", "b", "c"}
for i, v := range arr {
    fmt.Printf("%d: %s\n", i, v)
}

// Slice
s := []int{10, 20, 30}
for i, v := range s {
    fmt.Printf("Index %d: %d\n", i, v)
}

// Map
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

// String
str := "Hello"
for i, char := range str {
    fmt.Printf("%d: %c\n", i, char)  // %c = character
}

// Ignoring values
for i := range arr {
    fmt.Println(i)  // Chỉ in index
}

for _, v := range arr {
    fmt.Println(v)  // Chỉ in value
}
```

---

## 📝 COMPLETE EXAMPLE

```go
package main

import "fmt"

func main() {
    // ===== ARRAY =====
    fmt.Println("=== ARRAYS ===")
    scores := [5]int{10, 20, 30, 40, 50}
    for i, score := range scores {
        fmt.Printf("Score %d: %d\n", i, score)
    }

    // ===== SLICE =====
    fmt.Println("\n=== SLICES ===")
    fruits := []string{"Apple", "Banana"}
    fruits = append(fruits, "Orange")
    for _, fruit := range fruits {
        fmt.Println("-", fruit)
    }

    // ===== MAP =====
    fmt.Println("\n=== MAPS ===")
    ages := map[string]int{
        "Alice": 25,
        "Bob":   30,
        "Charlie": 28,
    }
    for name, age := range ages {
        fmt.Printf("%s: %d years old\n", name, age)
    }

    // Check & Delete
    if age, ok := ages["Alice"]; ok {
        fmt.Printf("Alice is %d\n", age)
    }
    delete(ages, "Bob")
    fmt.Println("After delete:", ages)
}
```

---

## 🎯 BÀI TẬP

### Bài 1: Array & Slice

```go
// Tạo array 5 số
// Convert thành slice
// Append 2 số mới
// In tất cả
```

### Bài 2: Map điểm học sinh

```go
students := map[string]float64{
    "Alice": 8.5,
    "Bob": 7.0,
    "Charlie": 9.0,
}

// In tất cả học sinh và điểm
// Tính điểm trung bình
// Tìm học sinh có điểm cao nhất
```

### Bài 3: Slice manipulation

```go
nums := []int{1, 2, 3, 4, 5}

// Tạo slice chỉ chứa phần tử từ index 1-3
// Append 100
// In len và cap
```

---

## 💡 LƯU Ý

1. **Arrays**: Fixed length, cần chỉ định size
2. **Slices**: Dynamic, linh hoạt hơn
3. **Maps**: Khóa-giá trị, không có order
4. **Nil**: Check trước khi dùng map
5. **Range**: Không có order với maps

---

## ✅ CHECKLIST

- [ ] Hiểu arrays (fixed length)
- [ ] Hiểu slices (dynamic)
- [ ] Biết append, copy, slicing
- [ ] Sử dụng maps
- [ ] Iterate collections với range

---

**Tiếp theo:** Lesson 5 - Control Flow (If, For, Switch)
