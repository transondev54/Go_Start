# Lesson 4: Generics & Type Parameters

## 📖 Nội dung bài học

1. Generics là gì? (Go 1.18+)
2. Type parameters cơ bản
3. Constraints
4. Generic functions
5. Generic types (structs, interfaces)
6. Practical examples

---

## 1️⃣ GENERICS LÀ GÌ?

### Định nghĩa

**Generics** cho phép bạn viết code mà hoạt động với nhiều types mà vẫn giữ type safety.

### Trước Generics (Go < 1.18)

```go
// ❌ Sai: Type-specific code
func IntSliceSum(nums []int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return sum
}

func FloatSliceSum(nums []float64) float64 {
    sum := 0.0
    for _, n := range nums {
        sum += n
    }
    return sum
}

// Code duplication!
```

### Với Generics (Go 1.18+)

```go
// ✅ Đúng: Generic code
func SliceSum[T constraints.Number](nums []T) T {
    var sum T
    for _, n := range nums {
        sum += n
    }
    return sum
}

// Hoạt động với mọi numeric types!
intSum := SliceSum[int]([]int{1, 2, 3})
floatSum := SliceSum[float64]([]float64{1.1, 2.2, 3.3})
```

---

## 2️⃣ TYPE PARAMETERS CƠ BẢN

### Cú pháp

```go
// Generic function với 1 type parameter
func Print[T any](v T) {
    fmt.Println(v)
}

// Sử dụng
Print[int](42)
Print[string]("hello")
```

### Type Inference

```go
// Có thể omit type parameter
Print(42)      // T inferred as int
Print("hello") // T inferred as string
```

### Multiple Type Parameters

```go
// Generic function với 2 type parameters
func Pair[T, U any](first T, second U) {
    fmt.Println(first, second)
}

// Sử dụng
Pair(1, "one")       // T=int, U=string
Pair("a", 3.14)      // T=string, U=float64
```

---

## 3️⃣ CONSTRAINTS

### Định nghĩa

**Constraint** định nghĩa những types mà type parameter có thể nhận.

### any Constraint

```go
// T có thể là bất kỳ type nào
func Process[T any](v T) {
    // ...
}
```

### Comparable Constraint

```go
// T phải có thể so sánh (==, !=)
func Equals[T comparable](a, b T) bool {
    return a == b
}

// Hoạt động: int, string, arrays, structs
Equals(1, 1)              // true
Equals("a", "b")          // false
Equals([2]int{1, 2}, [2]int{1, 2}) // true

// ❌ Sai: slice không comparable
Equals([]int{1}, []int{1}) // compile error
```

### Interface Constraints

```go
// Constraint dựa trên interface
type Reader interface {
    Read(p []byte) (n int, err error)
}

func ReadData[T Reader](r T) {
    // T phải implement Reader interface
}
```

### Composite Constraints (Union)

```go
// Constraint với multiple types
type Number interface {
    int | int64 | float64 | float32
}

func Sum[T Number](nums []T) T {
    var sum T
    for _, n := range nums {
        sum += n
    }
    return sum
}
```

### Predefined Constraints

```go
import "golang.org/x/exp/constraints"

// Numeric types
func Sum[T constraints.Number](nums []T) T { ... }

// Integer types
func Max[T constraints.Integer](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Ordered types (int, uint, float, string)
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```

---

## 4️⃣ GENERIC FUNCTIONS

### Example 1: Generic Min/Max

```go
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Usage
fmt.Println(Min(1, 2))           // 1
fmt.Println(Min("a", "b"))       // "a"
fmt.Println(Max(3.14, 2.71))     // 3.14
```

### Example 2: Generic Filter

```go
func Filter[T any](items []T, predicate func(T) bool) []T {
    result := []T{}
    for _, item := range items {
        if predicate(item) {
            result = append(result, item)
        }
    }
    return result
}

// Usage
nums := []int{1, 2, 3, 4, 5}
evens := Filter(nums, func(n int) bool {
    return n%2 == 0
})
fmt.Println(evens) // [2, 4]
```

---

## 5️⃣ GENERIC TYPES

### Generic Structs

```go
// Generic Stack
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    if len(s.items) == 0 {
        var zero T
        return zero
    }

    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

// Usage
stack := Stack[int]{}
stack.Push(1)
stack.Push(2)
fmt.Println(stack.Pop()) // 2
```

### Generic Interfaces

```go
type Container[T any] interface {
    Add(T)
    Remove(T)
    Contains(T) bool
}

type List[T comparable] struct {
    items []T
}

func (l *List[T]) Add(item T) {
    l.items = append(l.items, item)
}

func (l *List[T]) Remove(item T) {
    for i, v := range l.items {
        if v == item {
            l.items = append(l.items[:i], l.items[i+1:]...)
            break
        }
    }
}

func (l *List[T]) Contains(item T) bool {
    for _, v := range l.items {
        if v == item {
            return true
        }
    }
    return false
}
```

---

## 6️⃣ PRACTICAL EXAMPLES

### Example 1: Generic Map Function

```go
func Map[T, U any](items []T, transform func(T) U) []U {
    result := make([]U, len(items))
    for i, item := range items {
        result[i] = transform(item)
    }
    return result
}

// Usage
nums := []int{1, 2, 3}
squared := Map(nums, func(n int) int {
    return n * n
})
fmt.Println(squared) // [1, 4, 9]

strs := []string{"a", "b", "c"}
lengths := Map(strs, func(s string) int {
    return len(s)
})
fmt.Println(lengths) // [1, 1, 1]
```

### Example 2: Generic Cache

```go
type Cache[K comparable, V any] struct {
    data map[K]V
    mu   sync.RWMutex
}

func NewCache[K comparable, V any]() *Cache[K, V] {
    return &Cache[K, V]{
        data: make(map[K]V),
    }
}

func (c *Cache[K, V]) Set(key K, value V) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

// Usage
cache := NewCache[string, int]()
cache.Set("count", 42)
if val, ok := cache.Get("count"); ok {
    fmt.Println(val) // 42
}
```

---

## 7️⃣ BEST PRACTICES

### 1. Đừng Over-Use Generics

```go
// ❌ Sai: Generics không cần thiết
func PrintValue[T any](v T) {
    fmt.Println(v)
}

// ✅ Đúng: Interface{} enough
func PrintValue(v interface{}) {
    fmt.Println(v)
}
```

### 2. Be Specific with Constraints

```go
// ❌ Sai: Quá generic
func Process[T any](items []T) {
    // ...
}

// ✅ Đúng: Specify constraint
func Process[T comparable](items []T) {
    // ...
}
```

### 3. Avoid Complex Constraints

```go
// ❌ Sai: Quá phức tạp
type Complex interface {
    int | int64 | float64 | string |
    []int | map[string]int | CustomType
}

// ✅ Đúng: Simple and clear
type Number interface {
    int | int64 | float64
}
```

---

## 📝 EXERCISES

1. **Generic Min/Max**: Tạo function generic tìm min/max
2. **Generic Filter**: Implement generic filter function
3. **Generic Stack**: Implement generic stack type

---

## 📚 RESOURCES

- [Go Generics Proposal](https://go.dev/blog/generics)
- [Generics Tutorial](https://go.dev/doc/tutorial/generics)
- [Constraints Package](https://pkg.go.dev/golang.org/x/exp/constraints)
