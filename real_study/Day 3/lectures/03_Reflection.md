# Lesson 3: Reflection in Go

## 📖 Nội dung bài học

1. Reflection là gì?
2. reflect package
3. Type inspection (TypeOf, ValueOf)
4. Dynamic access to struct fields
5. Dynamic method calling
6. Practical examples & best practices

---

## 1️⃣ REFLECTION LÀ GÌ?

### Định nghĩa

**Reflection** là khả năng để một chương trình **inspect và manipulate** các objects tại runtime.

### Khi nào dùng?

✅ ORM/Database mapping
✅ JSON marshaling/unmarshaling
✅ RPC frameworks
✅ Testing frameworks
✅ Configuration parsing

### Cảnh báo

⚠️ Reflection là **slow** - tránh dùng trong hot paths
⚠️ Tạo ra code khó debug
⚠️ Mất type safety

---

## 2️⃣ REFLECT PACKAGE

### Import

```go
import "reflect"
```

### Core Functions

```go
// Trả về Type của value
t := reflect.TypeOf(value)

// Trả về Value wrapper
v := reflect.ValueOf(value)

// Trả về zero value của type
zero := reflect.Zero(t)

// Tạo value từ type
v := reflect.New(t)
```

---

## 3️⃣ TYPE INSPECTION

### TypeOf - Lấy type information

```go
// Primitive types
t := reflect.TypeOf(42)
fmt.Println(t)           // int
fmt.Println(t.Name())    // int
fmt.Println(t.Kind())    // int (Kind là category)

// Structs
type Person struct {
    Name string
    Age  int
}

p := Person{"Alice", 30}
t := reflect.TypeOf(p)
fmt.Println(t.Name())    // Person
fmt.Println(t.NumField())// 2
```

### ValueOf - Lấy value information

```go
v := reflect.ValueOf(42)
fmt.Println(v.Kind())    // int
fmt.Println(v.Interface())// 42 (returns interface{})

// Chuyển đổi
i := v.Interface().(int) // Type assertion
fmt.Println(i)
```

### Struct Field Inspection

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

t := reflect.TypeOf(Person{})

// Lặp qua fields
for i := 0; i < t.NumField(); i++ {
    field := t.Field(i)
    fmt.Printf("Field %d: %s (%s) = %s\n",
        i,
        field.Name,
        field.Type,
        field.Tag.Get("json"),
    )
}
// Output:
// Field 0: Name (string) = name
// Field 1: Age (int) = age
```

---

## 4️⃣ DYNAMIC ACCESS TO STRUCT FIELDS

### Đọc Field Values

```go
p := Person{"Alice", 30}
v := reflect.ValueOf(p)

// Lấy field theo index
nameField := v.Field(0)
fmt.Println(nameField.String()) // Alice

ageField := v.Field(1)
fmt.Println(ageField.Int())     // 30

// Lấy field theo tên
nameField = v.FieldByName("Name")
fmt.Println(nameField.String()) // Alice
```

### Ghi Field Values

```go
p := Person{"Alice", 30}
v := reflect.ValueOf(&p).Elem() // Phải là pointer

// Set field
v.FieldByName("Name").SetString("Bob")
v.FieldByName("Age").SetInt(25)

fmt.Println(p) // {Bob 25}
```

### Example: Generic JSON Unmarshaling

```go
func unmarshal(data []byte, v interface{}) error {
    m := make(map[string]interface{})
    if err := json.Unmarshal(data, &m); err != nil {
        return err
    }

    reflectValue := reflect.ValueOf(v).Elem()
    reflectType := reflectValue.Type()

    for i := 0; i < reflectType.NumField(); i++ {
        field := reflectType.Field(i)
        jsonKey := field.Tag.Get("json")

        if val, ok := m[jsonKey]; ok {
            reflectValue.Field(i).Set(reflect.ValueOf(val))
        }
    }

    return nil
}
```

---

## 5️⃣ DYNAMIC METHOD CALLING

### Method Lookup

```go
type Calculator struct{}

func (c Calculator) Add(a, b int) int {
    return a + b
}

calc := Calculator{}
v := reflect.ValueOf(calc)

// Lấy method
method := v.MethodByName("Add")

// Gọi method
args := []reflect.Value{reflect.ValueOf(5), reflect.ValueOf(3)}
result := method.Call(args)

fmt.Println(result[0].Int()) // 8
```

### Example: RPC Framework

```go
type Service struct{}

func (s Service) Sum(a, b int) int {
    return a + b
}

func call(service interface{}, methodName string, args ...interface{}) interface{} {
    v := reflect.ValueOf(service)
    method := v.MethodByName(methodName)

    // Convert args
    reflectArgs := make([]reflect.Value, len(args))
    for i, arg := range args {
        reflectArgs[i] = reflect.ValueOf(arg)
    }

    // Call method
    result := method.Call(reflectArgs)

    return result[0].Interface()
}

// Usage
service := Service{}
result := call(service, "Sum", 5, 3)
fmt.Println(result) // 8
```

---

## 6️⃣ BEST PRACTICES

### 1. Avoid Reflection in Hot Paths

```go
// ❌ Sai: Slow
for _, item := range items {
    v := reflect.ValueOf(item)
    processViaReflection(v) // Slow!
}

// ✅ Đúng: Cache reflection results
t := reflect.TypeOf(items[0])
for _, item := range items {
    v := reflect.ValueOf(item)
    processWithCachedType(v, t)
}
```

### 2. Always Check CanSet & CanAddr

```go
v := reflect.ValueOf(myValue)

// Kiểm tra trước khi set
if v.CanSet() {
    v.SetString("new value")
} else {
    fmt.Println("Cannot set value")
}
```

### 3. Handle Pointer/Non-Pointer

```go
func getStructValue(v interface{}) reflect.Value {
    val := reflect.ValueOf(v)

    // Nếu là pointer, deref
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }

    return val
}
```

### 4. Validate Before Type Assertion

```go
// ❌ Sai: Panic nếu type sai
i := v.Interface().(int)

// ✅ Đúng: Check type trước
if v.Kind() == reflect.Int {
    i := v.Int()
}
```

---

## 💡 EXAMPLES

### Example 1: Generic Struct Printer

```go
func printStruct(v interface{}) {
    val := reflect.ValueOf(v)
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }

    typ := val.Type()

    fmt.Printf("%s{\n", typ.Name())
    for i := 0; i < typ.NumField(); i++ {
        field := typ.Field(i)
        value := val.Field(i)
        fmt.Printf("  %s: %v\n", field.Name, value.Interface())
    }
    fmt.Println("}")
}
```

### Example 2: Generic Deep Copy

```go
func deepCopy(src interface{}) interface{} {
    srcVal := reflect.ValueOf(src)
    dstVal := reflect.New(srcVal.Type()).Elem()

    for i := 0; i < srcVal.NumField(); i++ {
        dstVal.Field(i).Set(srcVal.Field(i))
    }

    return dstVal.Interface()
}
```

---

## 📝 EXERCISES

1. **Field Inspection**: Lấy tất cả field names của struct
2. **Dynamic Setting**: Set struct fields từ map
3. **Method Call**: Gọi method động

---

## 📚 RESOURCES

- [Reflect Package](https://pkg.go.dev/reflect)
- [Go Reflection Laws](https://go.dev/blog/laws-of-reflection)
