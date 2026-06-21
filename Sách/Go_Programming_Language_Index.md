# The Go Programming Language - Comprehensive Index & Summary

**Book:** The Go Programming Language by Alan A. A. Donovan and Brian W. Kernighan

---

## Table of Contents - Quick Navigation

| Chapter | Topic                 | Key Concepts                                            |
| ------- | --------------------- | ------------------------------------------------------- |
| 1       | Getting Started       | Tour of Go, Hello World, Command Line                   |
| 2       | Program Structure     | Declarations, Variables, Constants                      |
| 3       | Data Types            | Numbers, Strings, Booleans, Composite Types             |
| 4       | Composite Types       | Arrays, Slices, Maps, Structs, JSON                     |
| 5       | Functions             | Declarations, Parameters, Return Values, Variadic       |
| 6       | Methods               | Method Declarations, Pointer Receivers, Value Receivers |
| 7       | Interfaces            | Interface Types, Assertion, Type Switches               |
| 8       | Goroutines & Channels | Concurrency, Channels, Buffered Channels, Close, Select |
| 9       | Concurrency           | Shared Memory, Sync Package, Race Detector              |
| 10      | Packages & Go Tool    | Import, Init Function, Blank Import                     |
| 11      | Testing               | Testing Package, Benchmark, Examples                    |
| 12      | Reflection            | Type & Value, Unsafe Package                            |
| 13      | Low-Level Programming | Unsafe, Cgo                                             |

---

## 1. GETTING STARTED

### 1.1 Introduction

- Go is a modern language designed for systems programming
- Compiled, concurrent, garbage-collected
- Simple syntax inspired by C

### 1.2 Hello, World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### 1.3 Command Line Arguments

- `os.Args` - access command-line arguments
- First argument is program name

### 1.4 Finding Duplicate Lines

- Basic file I/O example
- Using maps to count occurrences

### 1.5 The Go Tool

- `go run` - compile and run
- `go build` - compile to executable
- `go test` - run tests
- `go fmt` - format code
- `go vet` - analyze code for errors

---

## 2. PROGRAM STRUCTURE

### 2.1 Names

- Case-sensitive identifiers
- Exported (capitalized) vs unexported (lowercase)
- Blank identifier `_`

### 2.2 Declarations

Four major kinds:

- **const** - constants
- **type** - type declarations
- **var** - variables
- **func** - functions

### 2.3 Variables

```go
var x int
var y = 10              // type inference
z := 20                 // short declaration
```

- Zero values: 0 for numbers, "" for strings, false for bools, nil for pointers
- Multiple variable declaration
- Blank identifier for unused variables

### 2.4 Assignments

- Tuple assignment: `x, y = y, x`
- Multiple assignment from function returns

### 2.5 Type Declarations

```go
type Celsius float64
```

- Define new types based on existing types
- Named types have distinct identity

### 2.6 Packages and Files

- Each file must start with `package` declaration
- Typically one package per directory
- `main` package is special - contains entry point

### 2.7 Scope

- File scope
- Package scope
- Universe scope
- Functions create new scope

---

## 3. DATA TYPES

### 3.1 Integers

- Signed: `int8`, `int16`, `int32`, `int64`, `int`
- Unsigned: `uint8`, `uint16`, `uint32`, `uint64`, `uint`, `uintptr`
- `byte` is alias for `uint8`
- `rune` is alias for `int32`

### 3.2 Floating-Point Numbers

- `float32`, `float64`
- IEEE 754 representation
- Special values: `NaN`, `Inf`

### 3.3 Complex Numbers

- `complex64`, `complex128`
- Complex literals: `3+4i`
- `real()`, `imag()` functions

### 3.4 Booleans

- `true`, `false`
- Operators: `&&`, `||`, `!`

### 3.5 Strings

- Immutable sequence of bytes
- String literals: double quotes or backticks
- Escape sequences: `\n`, `\t`, `\\`
- Raw strings with backticks: `\n` is literal
- `len()` - number of bytes
- `s[i]` - i-th byte
- String concatenation with `+`

### 3.6 Constants

```go
const (
    KB = 1000
    MB = 1000 * KB
)
```

- Untyped or typed constants
- `iota` - enumeration constant generator
- Const expressions

### 3.7 Rune, Byte, String

- Rune: Unicode code point
- Byte: single byte
- String: sequence of bytes
- Type conversions between them

---

## 4. COMPOSITE TYPES

### 4.1 Arrays

```go
var a [3]int
b := [...]int{1, 2, 3}  // length inferred
```

- Fixed length
- Index from 0
- Initialization: `a := [3]int{1, 2, 3}`

### 4.2 Slices

```go
s := []int{1, 2, 3}
s = append(s, 4)
```

- Dynamic array
- Properties: pointer, length, capacity
- `s[i:j]` - slice from i to j-1
- `len(s)` - length
- `cap(s)` - capacity
- `append()` - add elements
- `copy()` - copy elements

### 4.3 Maps

```go
ages := make(map[string]int)
ages["alice"] = 31
delete(ages, "alice")
```

- Hash table/dictionary
- Key types must be comparable
- Safe to read missing key (returns zero value)
- Multiple return for existence check: `v, ok := ages["alice"]`

### 4.4 Structs

```go
type Employee struct {
    Name string
    ID   int
}
```

- Group related data
- Named fields
- Field initialization: `Employee{Name: "Alice", ID: 1}`
- Embedding: anonymous fields for composition

### 4.5 JSON

- `json.Marshal()` - Go value → JSON
- `json.Unmarshal()` - JSON → Go value
- Struct tags for marshaling: `` `json:"name"` ``
- Only exported fields are marshaled

### 4.6 Text and HTML Templates

- `text/template` package
- `html/template` package (auto-escapes)
- Template syntax: `{{.Field}}`

---

## 5. FUNCTIONS

### 5.1 Function Declarations

```go
func add(x, y int) int {
    return x + y
}
```

- Parameters and return types required
- Multiple return values: `func div(x, y int) (int, error)`
- Named return values

### 5.2 Multiple Return Values

```go
x, y := swap(1, 2)
q, r := div(10, 3)
```

### 5.3 Bare Return

```go
func split(sum int) (x, y int) {
    x = sum / 2
    y = sum - x
    return  // returns x, y
}
```

### 5.4 Variadic Functions

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

### 5.5 Deferred Function Calls

```go
defer fmt.Println("world")
fmt.Println("hello")
// Output: hello, world
```

- `defer` delays execution until function returns
- Multiple defers execute in LIFO order
- Common use: cleanup, closing files

### 5.6 Panic and Recover

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("recovered from:", r)
    }
}()
panic("something went wrong")
```

- `panic()` - stops normal execution
- `recover()` - regain control of panicked goroutine

### 5.7 Function Values

```go
var f func(int) int
f = square
fmt.Println(f(3))
```

### 5.8 Anonymous Functions

```go
func(x int) int {
    return x * x
}(3)
```

- Closures capture surrounding scope
- Common in concurrent programming

---

## 6. METHODS

### 6.1 Method Declarations

```go
type Point struct{ X, Y float64 }

func (p Point) Distance() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}
```

- Methods are functions with receiver
- Receiver comes between `func` and method name

### 6.2 Pointer Receivers

```go
func (p *Point) Move(dx, dy float64) {
    p.X += dx
    p.Y += dy
}
```

- Modifies original struct
- Go auto-converts: `p.Move(1, 1)` works even if p is `Point`

### 6.3 Composing Types by Struct Embedding

```go
type ColoredPoint struct {
    Point
    Color string
}
```

- Embedded type methods promoted to outer type
- Inheritance-like behavior

### 6.4 Method Value and Method Expression

```go
p := Point{1, 2}
f := p.Distance    // method value
g := (*Point).Move  // method expression
```

---

## 7. INTERFACES

### 7.1 Interface Types

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

- Contract: what a type must do
- No explicit implementation declaration
- Structural subtyping (duck typing)

### 7.2 Interface Satisfaction

- Type satisfies interface if it has all methods
- `io.Reader`, `io.Writer` - standard interfaces

### 7.3 Type Assertions

```go
r := reader.(io.Writer)
w, ok := reader.(io.Writer)
```

- Extract concrete value from interface
- Panic if wrong type (without `ok` check)

### 7.4 Type Switches

```go
switch x := v.(type) {
case int:
    fmt.Println("int:", x)
case string:
    fmt.Println("string:", x)
}
```

### 7.5 Interface{} - Empty Interface

```go
var v interface{}
```

- Empty interface satisfied by all types
- Used for generic functions
- Type assertion needed to use

### 7.6 Error Interface

```go
type error interface {
    Error() string
}
```

- Convention: last return value is error or nil
- Check: `if err != nil { /* handle */ }`

---

## 8. GOROUTINES & CHANNELS

### 8.1 Goroutines

```go
go func() {
    fmt.Println("concurrent")
}()
```

- Lightweight concurrency
- Prefix function call with `go`
- Main function waits for goroutines

### 8.2 Channels

```go
ch := make(chan int)
ch <- x  // send
x := <- ch  // receive
```

- Communicate between goroutines
- Unbuffered: send blocks until receive
- Buffered: `ch := make(chan int, 10)`

### 8.3 Unidirectional Channels

```go
func send(ch chan<- int)       // send-only
func receive(ch <-chan int)    // receive-only
```

### 8.4 Closing Channels

```go
close(ch)
x, ok := <-ch  // ok = false if closed
```

### 8.5 Range over Channels

```go
for x := range ch {
    fmt.Println(x)
}
```

### 8.6 Select Statement

```go
select {
case x := <-ch1:
    fmt.Println(x)
case y := <-ch2:
    fmt.Println(y)
default:
    fmt.Println("no data")
}
```

- Wait on multiple channel operations
- Executes first case that's ready
- `default` if none ready

---

## 9. CONCURRENCY WITH SHARED MEMORY

### 9.1 Mutex

```go
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()
```

- Protect shared data
- Mutual exclusion lock

### 9.2 Sync Package

- `Mutex` - exclusive lock
- `RWMutex` - read-write lock
- `WaitGroup` - wait for goroutines
- `Once` - execute function once
- `Cond` - condition variable

### 9.3 Race Detector

```bash
go run -race main.go
```

- Detects concurrent memory access
- First argument to go tool

### 9.4 Atomics

- `sync/atomic` package
- Atomic operations on shared variables

---

## 10. PACKAGES AND GO TOOL

### 10.1 Package Declaration and Imports

```go
package main

import (
    "fmt"
    "math"
)
```

- First statement in file: package declaration
- Blank lines separate import groups

### 10.2 Exported & Unexported Names

- Capitalized: exported (visible outside package)
- Lowercase: unexported (package-private)

### 10.3 Package Initialization

- `init` function runs before `main`
- Useful for package setup

### 10.4 Blank Import

```go
import _ "database/sql/driver"
```

- Import for side effects only
- Executes `init` functions

### 10.5 Go Tool

- `go get` - download packages
- `go list` - list packages
- `go doc` - documentation
- `go mod init` - initialize module
- `go.mod` - module definition

---

## 11. TESTING

### 11.1 The Test Function

```go
func TestSqrt(t *testing.T) {
    v := math.Sqrt(2.0)
    if v != 1.4142... {
        t.Errorf("expected 1.414..., got %v", v)
    }
}
```

- File: `*_test.go`
- Function: `TestXxx` (Xxx starts with capital)
- Parameter: `*testing.T`

### 11.2 Table-Driven Tests

```go
var tests = []struct {
    input  int
    output int
}{
    {2, 4},
    {3, 9},
}
```

- Test multiple inputs efficiently

### 11.3 Benchmarks

```go
func BenchmarkSqrt(b *testing.B) {
    for i := 0; i < b.N; i++ {
        math.Sqrt(2)
    }
}
```

- Measure performance
- Run: `go test -bench=.`

### 11.4 Examples

```go
func ExampleSqrt() {
    fmt.Println(math.Sqrt(4))
    // Output: 2
}
```

- Runnable examples
- Serves as documentation

---

## 12. REFLECTION

### 12.1 The Laws of Reflection

1. Reflection goes from interface value to reflection object
2. Reflection goes from reflection object to interface value
3. To modify a value, it must be addressable

### 12.2 Reflect Package

```go
t := reflect.TypeOf(v)
f := reflect.ValueOf(v)
```

- `TypeOf()` - get type
- `ValueOf()` - get value
- `Kind()` - category of type

### 12.3 Display Recursive Values

- Inspect arbitrary values
- Handle cycles

### 12.4 Example: Encoding S-Expressions

- Use reflection for serialization

---

## 13. LOW-LEVEL PROGRAMMING

### 13.1 Unsafe Package

```go
type Pointer *ArbitraryType
```

- Bypass type system
- For C interoperability
- Use rarely and carefully

### 13.2 Cgo

```go
/*
#include <stdio.h>
*/
import "C"
```

- Call C code from Go
- Go code can be called from C

---

## KEY PATTERNS & BEST PRACTICES

### Error Handling

```go
if err != nil {
    return err
}
```

### Defer Pattern

```go
f, _ := os.Open(file)
defer f.Close()
```

### Interface Design

- Keep interfaces small
- Export only what's needed
- Use io.Reader, io.Writer

### Concurrency Patterns

- Use channels for communication
- Goroutines are cheap
- Select for multiplexing

### Naming Conventions

- Package names: short, lowercase
- Function names: mixed case, exported capitalized
- Interfaces: -er suffix (Reader, Writer, Closer)

---

## STANDARD LIBRARY HIGHLIGHTS

| Package         | Purpose                    |
| --------------- | -------------------------- |
| `fmt`           | Formatting and printing    |
| `io`            | Basic I/O interfaces       |
| `os`            | Operating system           |
| `strings`       | String manipulation        |
| `math`          | Mathematical functions     |
| `time`          | Time and duration          |
| `json`          | JSON encoding/decoding     |
| `encoding/json` | JSON support               |
| `net/http`      | HTTP client and server     |
| `database/sql`  | Database access            |
| `sync`          | Synchronization primitives |
| `context`       | Cancellation and timeouts  |
| `sort`          | Sorting                    |
| `regexp`        | Regular expressions        |

---

## COMMON GOTCHAS & TIPS

1. **Nil Slices vs Empty Slices**
   - `var s []int` is nil, `s = []int{}` is empty
   - Both have len=0, but nil has cap=0

2. **Map Zero Value**
   - `var m map[string]int` is nil
   - Must use `make()` to initialize

3. **String Indexing**
   - `s[i]` returns byte, not rune
   - Use `for range s` for Unicode

4. **Goroutine Lifecycle**
   - Main function doesn't wait for goroutines
   - Program exits when main returns

5. **Channel Blocking**
   - Send on full buffered channel blocks
   - Receive on empty channel blocks

6. **Panic vs Error**
   - `panic` for programmer errors
   - `error` return for runtime issues

---

## QUICK REFERENCE - OPERATORS & SYNTAX

### Operators

- Arithmetic: `+`, `-`, `*`, `/`, `%`
- Comparison: `==`, `!=`, `<`, `>`, `<=`, `>=`
- Logical: `&&`, `||`, `!`
- Bitwise: `&`, `|`, `^`, `<<`, `>>`
- Assignment: `=`, `+=`, `-=`, `*=`, `/=`, `%=`, etc.

### Control Flow

```go
if condition { } else if { } else { }
for i := 0; i < n; i++ { }
for range slice { }
switch x { case 1: case 2: default: }
```

### Function Definition

```go
func name(param type) returnType { }
func name() (type1, type2) { }
```

---

**Note**: This index is designed to help you quickly find topics without reading the entire PDF. Use this as a quick reference, then refer back to the book for detailed explanations and code examples.
