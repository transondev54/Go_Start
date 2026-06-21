# Lesson 8: Review, Q&A & Final Assessment

## 📖 Nội dung bài học

Ôn tập lại toàn bộ kiến thức ngày 1 và đánh giá mức độ hiểu biết.

---

## 🎯 RECAP - TÓM TẮT KIẾN THỨC

### Lesson 1: Introduction

- Go là ngôn ngữ compiled, concurrent, garbage-collected
- Tính năng: nhanh, đơn giản, perfect cho backend
- Cài đặt & cấu trúc Hello World

### Lesson 2: Hello World & Setup

- Tạo project với `go mod init`
- Cấu trúc thư mục (cmd, pkg, internal)
- `go run` vs `go build`

### Lesson 3: Variables & Types

- 5 cách khai báo biến
- 7 kiểu dữ liệu cơ bản (int, float, string, bool, complex, rune, byte)
- Type conversion explicit
- Constants & iota

### Lesson 4: Collections

- **Arrays**: Fixed length, indexed
- **Slices**: Dynamic, linh hoạt, `append()`, `copy()`
- **Maps**: Key-value, `delete()`, safe read

### Lesson 5: Control Flow

- **If/Else**: Conditions không cần parentheses
- **For**: 3 dạng (C-style, range, while-like)
- **Switch**: Multiple cases, fallthrough
- **Break/Continue**: Loop control

### Lesson 6: Functions

- Khai báo: `func name(params) returnType`
- Multiple returns & error handling
- Named returns & bare return
- Variadic functions: `func f(nums ...int)`
- Anonymous functions & closures

### Lesson 7: Structs & Methods

- Struct: Group related data
- Methods: Functions với receiver
- Value receiver vs Pointer receiver
- Embedding: Composition > Inheritance

---

## 🧪 MINI QUIZ

### Level 1: Easy (5 câu)

**Q1:** Hàm entry point trong Go là gì?

<details>
<summary>Answer</summary>
`func main()` trong package `main`
</details>

**Q2:** Khác nhau giữa `var x int` và `x := 5`?

<details>
<summary>Answer</summary>
- `var x int`: Package level hoặc trong hàm, zero value = 0
- `x := 5`: Chỉ trong hàm, type inference
</details>

**Q3:** Phần tử zero value của slice là gì?

<details>
<summary>Answer</summary>
`nil`, có len=0, cap=0
</details>

**Q4:** Go có `while` loop không?

<details>
<summary>Answer</summary>
Không, dùng `for i < 10 { }` để thay thế
</details>

**Q5:** Struct method receiver viết ở đâu?

<details>
<summary>Answer</summary>
Giữa `func` và method name: `func (s Struct) Method()`
</details>

---

### Level 2: Medium (5 câu)

**Q1:** Viết code in bảng cửu chương 7 (1-10)

**Q2:** Khác biệt giữa value receiver và pointer receiver?

**Q3:** Code tính average của slice numbers?

**Q4:** Cách xóa phần tử tại index 2 trong slice?

**Q5:** Map được khởi tạo như thế nào?

---

### Level 3: Hard (3 câu)

**Q1:** Closure là gì? Cho ví dụ?

**Q2:** Làm sao để capture variable trong closure?

**Q3:** Khác biệt giữa array, slice, map khi iterate?

---

## 📚 COMMON MISTAKES

### ❌ Mistake 1: Quên import package

```go
// WRONG
func main() {
    Println("Hello")  // Undefined
}

// RIGHT
import "fmt"
func main() {
    fmt.Println("Hello")
}
```

### ❌ Mistake 2: Nil map assignment

```go
// WRONG
var m map[string]int
m["key"] = 1  // panic

// RIGHT
m := make(map[string]int)
m["key"] = 1
```

### ❌ Mistake 3: For range không ghi index/value

```go
// Ghi cả hai
for i, v := range arr { }

// Chỉ ghi value
for _, v := range arr { }

// Chỉ ghi index
for i := range arr { }
```

### ❌ Mistake 4: String indexing trả về byte

```go
s := "Hello"
fmt.Println(s[0])  // 72 (byte value, không 'H')

// Để get rune
for i, r := range s {
    fmt.Printf("%d: %c\n", i, r)  // %c format char
}
```

### ❌ Mistake 5: Pointer receiver vs Value receiver

```go
// Value receiver (không modify original)
func (p Person) SetAge(age int) {
    p.Age = age  // Chỉ modify copy
}

// Pointer receiver (modify original)
func (p *Person) SetAge(age int) {
    p.Age = age  // Modify original
}
```

---

## 🚀 BEST PRACTICES

### ✅ Best Practice 1: Error handling

```go
result, err := someFunction()
if err != nil {
    // Handle error
    return err
}
// Use result
```

### ✅ Best Practice 2: Naming conventions

- Packages: `lowercase`, no underscores
- Functions: `PascalCase` (exported), `camelCase` (unexported)
- Variables: `camelCase`
- Constants: `UPPERCASE`

### ✅ Best Practice 3: Keep functions small

- Một function = một trách nhiệm
- Dễ test, dễ maintain

### ✅ Best Practice 4: Use meaningful names

```go
// BAD
func f(x int) int { return x * 2 }

// GOOD
func doubleValue(value int) int { return value * 2 }
```

### ✅ Best Practice 5: Validate input

```go
func divide(x, y float64) (float64, error) {
    if y == 0 {
        return 0, errors.New("division by zero")
    }
    return x / y, nil
}
```

---

## 📋 FINAL ASSESSMENT CHECKLIST

### Kiến thức cơ bản

- [ ] Hiểu Go là gì & tại sao học
- [ ] Setup project với `go mod init`
- [ ] Khai báo & dùng biến
- [ ] Hiểu các kiểu dữ liệu
- [ ] Type conversion explicit

### Collections & Iteration

- [ ] Phân biệt array vs slice
- [ ] Dùng được slice operations (append, copy)
- [ ] Dùng được map
- [ ] Iterate với range

### Control Flow

- [ ] Viết if/else statements
- [ ] Viết for loops (3 dạng)
- [ ] Viết switch/case
- [ ] Dùng break & continue

### Functions

- [ ] Khai báo function
- [ ] Return multiple values
- [ ] Error handling
- [ ] Variadic functions
- [ ] Closures

### Structs & Methods

- [ ] Define structs
- [ ] Khai báo methods
- [ ] Value vs Pointer receivers
- [ ] Dùng embedded structs

---

## 💬 Q&A SECTION

### Q: Khi nào dùng pointer receiver?

**A:** Khi method cần modify struct. Ngược lại dùng value receiver.

### Q: Array vs Slice?

**A:** Array có fixed length, slice động. Slice linh hoạt hơn.

### Q: Sao Go không có inheritance?

**A:** Go dùng composition (embedding) thay vì inheritance - đơn giản hơn.

### Q: Error handling tại sao không dùng exception?

**A:** Go ưa explicit error handling hơn implicit exceptions - rõ ràng hơn.

### Q: Go có class không?

**A:** Không, dùng structs + methods thay vì classes.

### Q: Package level variable được khởi tạo khi nào?

**A:** Trước hàm `main()`, nhưng sau imports.

---

## 🎯 OBJECTIVES ACHIEVED?

Sau ngày học này, bạn nên có thể:

- ✅ Cài đặt Go environment
- ✅ Viết basic Go program
- ✅ Hiểu & dùng variables, types
- ✅ Làm việc với collections
- ✅ Viết control flow
- ✅ Tạo & gọi functions
- ✅ Tạo structs & methods
- ✅ Xử lý errors
- ✅ Viết mini projects

---

## 📊 FINAL SCORING CRITERIA

### Mini Projects (60%)

- **BMI Calculator** (20%): 0-20 points
- **Number Guessing Game** (20%): 0-20 points
- **Contact Manager** (20%): 0-20 points

### Quizzes (30%)

- **Mini Quiz** (10%): 0-10 points
- **Final Assessment Checklist** (20%): 0-20 points

### Bonus (10%)

- Extra features trong projects
- Code quality & best practices
- Additional features

### Total: 0-100 points

---

## 🎓 NEXT STEPS

### Ngày 2+:

- Interfaces
- Goroutines & Concurrency
- Error handling (custom errors)
- Testing
- Standard library
- Package design

### Để tiếp tục học:

1. Hoàn thành tất cả mini projects
2. Làm lại các bài tập nếu cần
3. Khám phá standard library
4. Tìm dự án thực tế để apply

---

## 💪 MOTIVATION

Bạn vừa hoàn thành **ngày đầu tiên** học Go!

Điều này không phải là kết thúc, mà là **khởi đầu** của hành trình lập trình với Go.

Những dự án lớn hơn đang chờ bạn phía trước. 🚀

---

**Hãy submit mini projects để nhận feedback!**
