# Lesson 7: Structs & Methods

## 📖 Nội dung bài học

1. Struct declaration
2. Struct fields
3. Nested structs
4. Methods
5. Pointer receivers

---

## 1️⃣ STRUCT DECLARATION

### Cơ bản

```go
type Person struct {
    Name string
    Age  int
    City string
}

// Tạo instance
p := Person{
    Name: "Alice",
    Age:  25,
    City: "Hanoi",
}

fmt.Println(p.Name)  // Alice
fmt.Println(p.Age)   // 25
```

### Positional initialization

```go
// Đúng thứ tự
p := Person{"Alice", 25, "Hanoi"}

// Or named fields
p := Person{
    Name: "Alice",
    Age:  25,
}
// City = "" (zero value)
```

### Anonymous fields (embedding)

```go
type Coordinate struct {
    X, Y float64
}

type Point struct {
    Coordinate  // Anonymous field
    Label       string
}

p := Point{
    Coordinate: Coordinate{X: 3.5, Y: 4.5},
    Label:      "A",
}

fmt.Println(p.X)     // 3.5 (promoted field)
fmt.Println(p.Label) // A
```

---

## 2️⃣ STRUCT FIELDS

### Field types

```go
type Employee struct {
    Name      string
    Age       int
    Salary    float64
    Active    bool
    Tags      []string
    Metadata  map[string]string
}

emp := Employee{
    Name:   "Bob",
    Age:    30,
    Salary: 50000.0,
    Active: true,
    Tags:   []string{"developer", "lead"},
    Metadata: map[string]string{
        "level": "senior",
    },
}
```

### Exported vs Unexported

```go
type Account struct {
    Username  string      // Exported (visible outside package)
    password  string      // Unexported (package-private)
    Balance   float64
    privateId int
}

// Outside package, can only access:
// - account.Username
// - account.Balance
// Cannot access password, privateId
```

### Default values (zero values)

```go
var p Person
fmt.Println(p.Name)  // "" (empty string)
fmt.Println(p.Age)   // 0
fmt.Println(p.City)  // "" (empty string)
```

---

## 3️⃣ NESTED STRUCTS

### Struct of structs

```go
type Address struct {
    Street string
    City   string
    Postal string
}

type Employee struct {
    Name    string
    Address Address
}

emp := Employee{
    Name: "Alice",
    Address: Address{
        Street: "123 Main St",
        City:   "Hanoi",
        Postal: "10000",
    },
}

fmt.Println(emp.Address.City)  // Hanoi
```

### Embedding (composition)

```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person
    EmployeeID string
}

emp := Employee{
    Person: Person{
        Name: "Bob",
        Age:  30,
    },
    EmployeeID: "E001",
}

fmt.Println(emp.Name)       // Bob (promoted)
fmt.Println(emp.Age)        // 30 (promoted)
fmt.Println(emp.EmployeeID) // E001
```

---

## 4️⃣ METHODS

### Method declaration

```go
type Circle struct {
    Radius float64
}

// Method with value receiver
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Method with value receiver
func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Usage
circle := Circle{Radius: 5}
fmt.Println(circle.Area())       // 78.54
fmt.Println(circle.Perimeter())  // 31.42
```

---

## 5️⃣ POINTER RECEIVERS

### Value receiver vs Pointer receiver

```go
type Account struct {
    Balance float64
}

// Value receiver (không thay đổi original)
func (a Account) GetBalance() float64 {
    return a.Balance
}

// Pointer receiver (thay đổi original)
func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}

func (a *Account) Withdraw(amount float64) error {
    if amount > a.Balance {
        return fmt.Errorf("insufficient funds")
    }
    a.Balance -= amount
    return nil
}

// Usage
acc := Account{Balance: 1000}
fmt.Println(acc.GetBalance())  // 1000

acc.Deposit(500)
fmt.Println(acc.GetBalance())  // 1500

acc.Withdraw(200)
fmt.Println(acc.GetBalance())  // 1300
```

### Khi nào dùng pointer receiver

- **Pointer**: Muốn modify struct
- **Value**: Chỉ cần read data

```go
// Value receiver (safe, no side effects)
func (c Circle) Area() float64 { ... }

// Pointer receiver (modify state)
func (p *Person) HaveBirthday() {
    p.Age++
}
```

---

## 📝 COMPLETE EXAMPLE

```go
package main

import (
    "fmt"
    "math"
)

// 1. Struct definition
type Student struct {
    Name    string
    Age     int
    Grade   float64
    Active  bool
}

// 2. Value receiver method
func (s Student) GetInfo() string {
    return fmt.Sprintf("Name: %s, Age: %d, Grade: %.2f",
        s.Name, s.Age, s.Grade)
}

// 3. Pointer receiver method
func (s *Student) UpdateGrade(newGrade float64) {
    if newGrade >= 0 && newGrade <= 10 {
        s.Grade = newGrade
    }
}

// 4. Method returning boolean
func (s Student) IsPassing() bool {
    return s.Grade >= 5.0
}

// 5. Struct with composition
type University struct {
    Name     string
    Students []Student
}

func (u *University) AddStudent(s Student) {
    u.Students = append(u.Students, s)
}

func (u University) GetStudentCount() int {
    return len(u.Students)
}

func main() {
    // Create students
    student1 := Student{
        Name:   "Alice",
        Age:    20,
        Grade:  8.5,
        Active: true,
    }

    fmt.Println(student1.GetInfo())
    fmt.Println("Passing:", student1.IsPassing())

    // Update grade
    student1.UpdateGrade(9.0)
    fmt.Println("Updated Grade:", student1.Grade)

    // Create university
    uni := University{
        Name:     "Hanoi University",
        Students: []Student{},
    }

    uni.AddStudent(student1)
    uni.AddStudent(Student{Name: "Bob", Age: 21, Grade: 7.0})

    fmt.Printf("\n%s has %d students\n", uni.Name, uni.GetStudentCount())
}
```

---

## 🎯 BÀI TẬP

### Bài 1: Rectangle struct

```go
type Rectangle struct {
    Width, Height float64
}

// Methods:
// - Area() float64
// - Perimeter() float64
// - IsSquare() bool

r := Rectangle{Width: 5, Height: 5}
fmt.Println("Area:", r.Area())
fmt.Println("Perimeter:", r.Perimeter())
fmt.Println("Is Square:", r.IsSquare())
```

### Bài 2: Bank account

```go
type BankAccount struct {
    Owner    string
    Balance  float64
    Number   string
}

// Methods:
// - Deposit(amount float64)
// - Withdraw(amount float64) error
// - GetBalance() float64
// - Display()

account := BankAccount{Owner: "Alice", Balance: 1000, Number: "001"}
account.Deposit(500)
account.Withdraw(200)
account.Display()
```

### Bài 3: Todo list

```go
type Todo struct {
    Title       string
    Description string
    Done        bool
}

type TodoList struct {
    Items []Todo
}

// Methods:
// - AddTodo(title, desc string)
// - CompleteTodo(index int)
// - ListAll()
// - ListPending()
```

---

## 💡 LƯU Ý

1. **Methods**: Chỉ định được cho named types (không primitive)
2. **Receiver**: Viết giữa `func` và method name
3. **Pointer receiver**: Cần khi muốn modify struct
4. **Embedding**: Giống inheritance nhưng là composition
5. **Exported**: Capitalize first letter (e.g., `GetBalance()`)

---

## ✅ CHECKLIST

- [ ] Tạo được structs
- [ ] Hiểu exported vs unexported fields
- [ ] Viết được methods với value receiver
- [ ] Viết được methods với pointer receiver
- [ ] Dùng được struct embedding

---

**Tiếp theo:** Lesson 8 - Review & Q&A
