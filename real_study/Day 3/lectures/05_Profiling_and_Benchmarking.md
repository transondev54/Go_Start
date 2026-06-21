# Lesson 5: Profiling & Benchmarking

## 📖 Nội dung bài học

1. Profiling là gì?
2. Benchmarking Go code
3. CPU profiling
4. Memory profiling
5. Analyzing profiles
6. Best practices

---

## 1️⃣ PROFILING LÀ GÌ?

### Định nghĩa

**Profiling** là quá trình đo lường performance của chương trình.

### Các loại profiling:

- **CPU Profiling**: Xem function nào tốn CPU nhiều nhất
- **Memory Profiling**: Xem chương trình dùng bao nhiêu memory
- **Goroutine Profiling**: Xem có bao nhiêu goroutines chạy
- **Trace**: Timeline chi tiết của execution

### Khi nào profile?

✅ Chương trình chạy chậm
✅ Dùng quá nhiều memory
✅ Optimize code
✅ Performance regression detection

---

## 2️⃣ BENCHMARKING GO CODE

### Viết Benchmarks

```go
// benchmark_test.go
func BenchmarkSum(b *testing.B) {
    nums := []int{1, 2, 3, 4, 5}

    // b.N được Go tự động điều chỉnh
    for i := 0; i < b.N; i++ {
        sum := 0
        for _, n := range nums {
            sum += n
        }
    }
}
```

### Chạy Benchmarks

```bash
# Chạy benchmark
go test -bench=. -benchmem

# Output
# BenchmarkSum-8    100000000    10.25 ns/op    0 B/op    0 allocs/op
#            ^        ^            ^              ^         ^
#          name      N          ns/op          B/op     allocs/op
```

### Benchmark Results

```
100000000        →  Số lần function chạy
10.25 ns/op      →  Time per operation (nanoseconds)
0 B/op           →  Bytes allocated per operation
0 allocs/op      →  Number of allocations per operation
```

### Example Benchmarks

```go
func BenchmarkSliceAppend(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var s []int
        for j := 0; j < 100; j++ {
            s = append(s, j)
        }
    }
}

func BenchmarkSlicePreAllocate(b *testing.B) {
    for i := 0; i < b.N; i++ {
        s := make([]int, 0, 100)
        for j := 0; j < 100; j++ {
            s = append(s, j)
        }
    }
}
```

### Benchmark Tips

```go
// Reset timer nếu setup tốn time
func BenchmarkComplexSetup(b *testing.B) {
    setup() // Preparation
    b.ResetTimer() // Reset timer

    for i := 0; i < b.N; i++ {
        // Actual benchmark
    }
}

// Stop timer tạm thời
func BenchmarkWithCleanup(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Benchmark code
        b.StopTimer()
        cleanup()
        b.StartTimer()
    }
}

// Sub-benchmarks
func BenchmarkAlgorithms(b *testing.B) {
    b.Run("Linear", BenchmarkLinearSearch)
    b.Run("Binary", BenchmarkBinarySearch)
}
```

---

## 3️⃣ CPU PROFILING

### Profile via Testing

```bash
# Generate CPU profile
go test -cpuprofile=cpu.prof -bench=.

# Analyze profile
go tool pprof cpu.prof

# Commands in pprof
(pprof) top         # Top functions by CPU time
(pprof) list main   # Show code with CPU info
(pprof) web         # Open in browser (need graphviz)
```

### CPU Profile Example

```
(pprof) top
Showing nodes accounting for 2.5s, 92.6% of 2.7s total
      flat  flat%   sum%        cum   cum%
     1.5s 55.6% 55.6%      1.8s 66.7%  main.expensiveFunction
     0.8s 29.6% 85.2%      0.9s 33.3%  main.slowHelper
     0.2s  7.4% 92.6%      0.2s  7.4%  runtime.memclr
```

### CPU Profile dari HTTP Server

```go
import _ "net/http/pprof"

func main() {
    go func() {
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    // Your server code
}
```

```bash
# Access profile từ browser hoặc CLI
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

---

## 4️⃣ MEMORY PROFILING

### Benchmark Memory

```bash
# Memory allocation info
go test -bench=. -benchmem

# Generates heap profile
go test -memprofile=mem.prof -bench=.
```

### Memory Profile Analysis

```bash
go tool pprof mem.prof

(pprof) top -cum
Showing nodes accounting for 512.5MB, 95.2% of 536.3MB total
      flat  flat%   sum%        cum   cum%
 256.0MB 47.8% 47.8%    384.0MB 71.6%  main.processData
 128.0MB 23.8% 71.6%    256.0MB 47.8%  main.buildCache
  64.0MB 11.9% 83.5%    128.0MB 23.8%  strings.Join
  ...
```

### Memory Optimization Example

```go
// ❌ Sai: Nhiều allocations
func ProcessStrings(strs []string) string {
    var result string
    for _, s := range strs {
        result = result + s + ", " // Allocates every time!
    }
    return result
}

// ✅ Đúng: Một allocation
func ProcessStrings(strs []string) string {
    var buf strings.Builder
    for i, s := range strs {
        if i > 0 {
            buf.WriteString(", ")
        }
        buf.WriteString(s)
    }
    return buf.String()
}
```

---

## 5️⃣ ANALYZING PROFILES

### PProf Commands

```bash
# top: Show top functions
(pprof) top10        # Top 10 functions

# list: Show source code with metrics
(pprof) list functionName

# web: Generate graph (need graphviz)
(pprof) web

# pdf: Generate PDF report
(pprof) pdf > profile.pdf

# disasm: Disassemble function
(pprof) disasm functionName
```

### Example Analysis

```bash
# Collect profile
go test -cpuprofile=cpu.prof -bench=BenchmarkSort

# Analyze
go tool pprof cpu.prof

# Show top 20 functions
(pprof) top20

# Show functions in specific package
(pprof) top20 -base=profile1.prof profile2.prof  # Compare profiles

# Focus on specific function
(pprof) list slowFunction

# Generate graph
(pprof) web processData
```

---

## 6️⃣ TRACE ANALYSIS

### Generate Trace

```bash
# Generate trace file
go test -trace=trace.out -bench=.

# Analyze trace
go tool trace trace.out
```

### Trace Information

- Goroutine execution timeline
- System call details
- Memory allocations
- Processor utilization

---

## 💡 OPTIMIZATION WORKFLOW

### 1. Benchmark Your Code

```bash
go test -bench=. -benchmem | tee baseline.txt
```

### 2. Profile to Find Bottlenecks

```bash
go test -cpuprofile=cpu.prof -bench=.
go tool pprof cpu.prof
```

### 3. Make Optimization

```go
// Optimize code
```

### 4. Benchmark Again

```bash
go test -bench=. -benchmem | tee optimized.txt
benchstat baseline.txt optimized.txt
```

### 5. Compare Results

```bash
benchstat baseline.txt optimized.txt
```

---

## 🔍 BENCHSTAT USAGE

### Comparing Benchmarks

```bash
# Run before optimization
go test -bench=. -benchmem > before.txt

# Run after optimization
go test -bench=. -benchmem > after.txt

# Compare
benchstat before.txt after.txt
```

### Example Output

```
name          old time/op    new time/op    delta
Sum-8         10.25ns ± 2%    5.10ns ± 3%  -50.24%  (p=0.000 n=10+10)

name          old alloc/op   new alloc/op   delta
Sum-8         0.00B ± 0%     0.00B ± 0%     (all equal)
```

---

## 📝 EXERCISES

1. **Benchmark**: Viết benchmark cho function của bạn
2. **CPU Profile**: Profile code và tìm bottleneck
3. **Memory Optimization**: Optimize memory usage
4. **Comparison**: Compare optimization results

---

## 📚 RESOURCES

- [Testing Package](https://pkg.go.dev/testing)
- [Pprof Documentation](https://github.com/google/pprof/tree/master/doc)
- [Go Diagnostics](https://go.dev/doc/diagnostics)
