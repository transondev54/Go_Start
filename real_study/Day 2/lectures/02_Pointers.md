# Lesson 2: Pointers & Reference Semantics

## 📖 Nội dung bài học

1. Pointers là gì?
2. Syntax con trỏ
3. Value vs Pointer receivers
4. Khi nào dùng pointers?
5. Pointer mistakes

---

## 1️⃣ POINTERS LÀ GÌ?

### Định nghĩa

**Pointer** là biến lưu **địa chỉ bộ nhớ** của một biến khác.

```go
var x int = 42
var p *int = &x  // p là pointer đến x
```

### Hai operator chính

| Operator | Tên                  | Ý nghĩa                 |
| -------- | -------------------- | ----------------------- |
| `&`      | Address operator     | Lấy địa chỉ của biến    |
| `*`      | Dereference operator | Lấy giá trị tại địa chỉ |

---

## 2️⃣ SYNTAX CON TRỎ

### Khai báo pointers

```go
// Khai báo
var ptr *int
var p *string
var p2 **int  // pointer to pointer

// Khởi tạo
var x int = 10
p := &x
```

### Sử dụng

```go
x := 10
p := &x

fmt.Println(x)   // 10 (value)
fmt.Println(p)   // 0xc0000120a8 (address)
fmt.Println(*p)  // 10 (dereference - lấy value)

// Thay đổi qua pointer
*p = 20
fmt.Println(x)   // 20 (x thay đổi)
```

---

## 3️⃣ VALUE VS POINTER RECEIVERS

### Value Receiver - Copy

```go
type Counter struct {
    count int
}

// Value receiver - nhận copy
func (c Counter) Increment() {
    c.count++  // Chỉ thay đổi copy, không affect original
}

func main() {
    counter := Counter{count: 0}
    counter.Increment()
    fmt.Println(counter.count)  // 0 (không thay đổi!)
}
```

### Pointer Receiver - Reference

```go
type Counter struct {
    count int
}

// Pointer receiver - nhận reference
func (c *Counter) Increment() {
    c.count++  // Thay đổi original
}

func main() {
    counter := Counter{count: 0}
    counter.Increment()
    fmt.Println(counter.count)  // 1 (thay đổi!)
}
```

---

## 4️⃣ VÍ DỤ: UPDATE CONTACT

### Sai (Value receiver)

```go
type Contact struct {
    ID    int
    Name  string
    Phone string
}

func (c Contact) UpdateName(name string) {
    c.Name = name  // ❌ Không update được original
}

func main() {
    contact := Contact{ID: 1, Name: "John"}
    contact.UpdateName("Jane")
    fmt.Println(contact.Name)  // John (không đổi)
}
```

### Đúng (Pointer receiver)

```go
func (c *Contact) UpdateName(name string) {
    c.Name = name  // ✅ Update được original
}

func main() {
    contact := Contact{ID: 1, Name: "John"}
    contact.UpdateName("Jane")
    fmt.Println(contact.Name)  // Jane (đổi rồi)
}
```

---

## 5️⃣ KHI NÀO DÙNG POINTERS?

### ✅ Dùng Pointer Receiver khi

- **Thay đổi struct** - Update fields
- **Struct lớn** - Tránh copy toàn bộ struct
- **Implement interface** - Consistent với methods khác
- **Methods cần access fields** - Chỉnh sửa state

### ✅ Dùng Value Receiver khi

- **Immutable** - Không thay đổi struct
- **Struct nhỏ** - (int, string, nhỏ struct)
- **Thread-safe** - Copy ensures safety

---

## 6️⃣ NIL POINTERS

### Cẩn thận với nil

```go
var p *int
fmt.Println(p)      // <nil>
fmt.Println(*p)     // ❌ PANIC! nil pointer dereference

// Kiểm tra nil
if p != nil {
    fmt.Println(*p)  // ✅ Safe
}
```

---

## 📝 TÓM TẮT POINTERS

```go
// Khai báo & khởi tạo
var ptr *int
ptr = &x

// Thao tác
value := *ptr        // Dereference
address := &variable // Address

// Receivers
func (s *Struct) Method() {}  // Pointer receiver
func (s Struct) Method() {}   // Value receiver
```

---

## 💡 BEST PRACTICES

1. **Mặc định dùng pointer receiver** - cho structs
2. **Kiểm tra nil** - trước khi dereference
3. **Tránh pointer to pointer** - nếu không cần thiết
4. **Be consistent** - tất cả methods của type dùng cùng loại receiver

---

## 🎯 BƯỚC TIẾP THEO

- Đọc **Lesson 3** về Error Handling
- Fix Contact Manager để dùng pointer receivers
- Thực hành pointers trong mini project
