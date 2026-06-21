# Lesson 6: Optimization Techniques

## 📖 Nội dung bài học

1. Common performance bottlenecks
2. Memory pooling
3. Caching strategies
4. String optimization
5. Slice & array optimization
6. Concurrency optimization

---

## 1️⃣ COMMON PERFORMANCE BOTTLENECKS

### Memory Allocations

```go
// ❌ Sai: Allocates memory in loop
func ProcessItems(items []Item) []Result {
    var results []Result
    for _, item := range items {
        results = append(results, processItem(item))
    }
    return results
}

// ✅ Đúng: Pre-allocate
func ProcessItems(items []Item) []Result {
    results := make([]Result, len(items))
    for i, item := range items {
        results[i] = processItem(item)
    }
    return results
}
```

### String Concatenation

```go
// ❌ Sai: Creates multiple strings
func BuildString(parts []string) string {
    result := ""
    for _, part := range parts {
        result = result + part // Allocates every time!
    }
    return result
}

// ✅ Đúng: Use strings.Builder
func BuildString(parts []string) string {
    var buf strings.Builder
    for _, part := range parts {
        buf.WriteString(part)
    }
    return buf.String()
}
```

### Reflection Overhead

```go
// ❌ Sai: Reflection in hot path
for _, item := range items {
    v := reflect.ValueOf(item)
    processViaReflection(v) // Slow!
}

// ✅ Đúng: Process directly
for _, item := range items {
    process(item)
}
```

---

## 2️⃣ MEMORY POOLING

### sync.Pool

**sync.Pool** reuses allocated objects untuk mengurangi garbage collection.

```go
// Object pool
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

// Get dari pool
buf := bufferPool.Get().(*bytes.Buffer)
defer bufferPool.Put(buf)

// Use buffer
buf.WriteString("Hello")
fmt.Println(buf.String())
```

### Example: Connection Pool

```go
type ConnectionPool struct {
    pool sync.Pool
    max  int
}

func (p *ConnectionPool) Get() (*Connection, error) {
    val := p.pool.Get()
    if val == nil {
        return newConnection()
    }
    return val.(*Connection), nil
}

func (p *ConnectionPool) Put(conn *Connection) {
    conn.Reset()
    p.pool.Put(conn)
}
```

### Buffer Pool for HTTP

```go
import "sync"

type BufferPool struct {
    pool sync.Pool
    size int
}

func (p *BufferPool) Get() []byte {
    val := p.pool.Get()
    if val == nil {
        return make([]byte, 0, p.size)
    }
    b := val.([]byte)
    return b[:0]
}

func (p *BufferPool) Put(b []byte) {
    if cap(b) <= p.size {
        p.pool.Put(b)
    }
}
```

---

## 3️⃣ CACHING STRATEGIES

### Simple In-Memory Cache

```go
type Cache struct {
    data map[string]interface{}
    mu   sync.RWMutex
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    val, ok := c.data[key]
    return val, ok
}

func (c *Cache) Set(key string, val interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.data[key] = val
}
```

### TTL Cache

```go
type CacheItem struct {
    Value      interface{}
    ExpiresAt  time.Time
}

type TTLCache struct {
    data map[string]CacheItem
    mu   sync.RWMutex
}

func (c *TTLCache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()

    item, ok := c.data[key]
    if !ok {
        return nil, false
    }

    if time.Now().After(item.ExpiresAt) {
        return nil, false
    }

    return item.Value, true
}

func (c *TTLCache) Set(key string, val interface{}, ttl time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()

    c.data[key] = CacheItem{
        Value:     val,
        ExpiresAt: time.Now().Add(ttl),
    }
}
```

### LRU Cache

```go
type LRUCache struct {
    maxSize int
    cache   map[string]*Node
    list    *list.List
    mu      sync.Mutex
}

type Node struct {
    Key   string
    Value interface{}
}

func (c *LRUCache) Get(key string) (interface{}, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()

    node, ok := c.cache[key]
    if !ok {
        return nil, false
    }

    // Move to front (most recently used)
    c.list.MoveToFront(node)
    return node.Value, true
}

func (c *LRUCache) Set(key string, value interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()

    if node, ok := c.cache[key]; ok {
        node.Value = value
        c.list.MoveToFront(node)
        return
    }

    node := &Node{Key: key, Value: value}
    c.cache[key] = node
    c.list.PushFront(node)

    if len(c.cache) > c.maxSize {
        last := c.list.Back()
        c.list.Remove(last)
        delete(c.cache, last.Value.(*Node).Key)
    }
}
```

---

## 4️⃣ STRING OPTIMIZATION

### String Builder

```go
// ❌ Sai: O(n²) complexity
func ConcatStrings(strs []string) string {
    result := ""
    for _, s := range strs {
        result += s
    }
    return result
}

// ✅ Đúng: O(n) complexity
func ConcatStrings(strs []string) string {
    var buf strings.Builder
    for _, s := range strs {
        buf.WriteString(s)
    }
    return buf.String()
}
```

### Preallocate String Builder

```go
func ConcatStrings(strs []string) string {
    var buf strings.Builder

    // Calculate total size
    totalSize := 0
    for _, s := range strs {
        totalSize += len(s)
    }

    // Preallocate
    buf.Grow(totalSize)

    for _, s := range strs {
        buf.WriteString(s)
    }
    return buf.String()
}
```

---

## 5️⃣ SLICE & ARRAY OPTIMIZATION

### Pre-allocate Slices

```go
// ❌ Sai: Multiple allocations
var result []int
for i := 0; i < 1000; i++ {
    result = append(result, i)
}

// ✅ Đúng: Single allocation
result := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    result = append(result, i)
}
```

### Use Arrays When Size is Fixed

```go
// ❌ Sai: Slice overhead
func ProcessFixedSize() {
    items := make([]int, 5)
    // ...
}

// ✅ Đúng: Array (stack allocated)
func ProcessFixedSize() {
    var items [5]int
    // ...
}
```

### Minimize Slice Copies

```go
// ❌ Sai: Copies slice
func Process(items []int) {
    copy := items
    // Work with copy
}

// ✅ Đúng: Pass by reference
func Process(items []int) {
    // Work with items directly
}

// ✅ Đúng: Pass slice by value (header only copied)
func Process(items []int) {
    for i := range items {
        items[i] = process(items[i])
    }
}
```

---

## 6️⃣ CONCURRENCY OPTIMIZATION

### Reduce Lock Contention

```go
// ❌ Sai: Lock per operation
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    c.count++
    c.mu.Unlock()
}

// ✅ Đúng: Batch operations
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) IncBatch(amount int) {
    c.mu.Lock()
    c.count += amount
    c.mu.Unlock()
}
```

### Use sync.Map for Concurrent Access

```go
// Fast path for reads
var cache sync.Map

func Get(key string) (interface{}, bool) {
    return cache.Load(key)
}

func Set(key string, value interface{}) {
    cache.Store(key, value)
}

// Much faster than RWMutex for mostly read workloads
```

### Partition Data

```go
// ❌ Sai: Single lock for all data
type DataStore struct {
    mu   sync.RWMutex
    data map[string]interface{}
}

// ✅ Đúng: Multiple locks (sharding)
type ShardedStore struct {
    shards []*Shard
}

type Shard struct {
    mu   sync.RWMutex
    data map[string]interface{}
}

func (s *ShardedStore) Get(key string) (interface{}, bool) {
    shard := s.getShard(key)
    shard.mu.RLock()
    defer shard.mu.RUnlock()
    val, ok := shard.data[key]
    return val, ok
}
```

---

## 💡 OPTIMIZATION CHECKLIST

- ✅ Profile code to find bottlenecks
- ✅ Pre-allocate slices & strings.Builder
- ✅ Minimize memory allocations
- ✅ Use sync.Pool for frequent allocations
- ✅ Implement caching where appropriate
- ✅ Batch operations to reduce lock contention
- ✅ Use sync.Map for read-heavy workloads
- ✅ Partition data to reduce lock contention

---

## 📝 EXERCISES

1. **String Optimization**: Optimize string concatenation
2. **Cache Implementation**: Implement LRU cache
3. **Pool Usage**: Use sync.Pool for buffer pooling

---

## 📚 RESOURCES

- [Go Performance Blog](https://go.dev/blog/pprof)
- [Sync Package](https://pkg.go.dev/sync)
- [Strings Package](https://pkg.go.dev/strings)
