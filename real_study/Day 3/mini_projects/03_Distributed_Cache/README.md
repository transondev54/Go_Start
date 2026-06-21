# Project 3: Distributed Cache System

## 🎯 Mục tiêu

Xây dựng một **distributed in-memory cache** với:

- ✅ Thread-safe concurrent access
- ✅ TTL (Time To Live) support
- ✅ Eviction policies (LRU, LFU)
- ✅ Statistics & monitoring
- ✅ Network protocol (optional)
- ✅ Advanced patterns

---

## 📋 Yêu cầu

### Core Features

1. **Basic Cache Operations**
   - `Set(key, value, ttl)` - Store value with TTL
   - `Get(key)` - Retrieve value
   - `Delete(key)` - Remove value
   - `Clear()` - Clear all entries
   - `Exists(key)` - Check existence

2. **Concurrency Safety**
   - Thread-safe concurrent reads/writes
   - No data races or deadlocks
   - Efficient locking (minimize contention)
   - Support for high throughput

3. **TTL Management**
   - Automatic expiration
   - Lazy deletion (on access)
   - Periodic cleanup goroutine
   - TTL refresh capability

4. **Eviction Policies**
   - **LRU (Least Recently Used)**: Evict least recently accessed
   - **LFU (Least Frequently Used)**: Evict least frequently used
   - **FIFO (First In First Out)**: Evict oldest entries
   - Configurable max size

5. **Statistics**
   - Hit/miss ratio
   - Eviction count
   - Memory usage
   - Operations per second

6. **Data Models**
   ```go
   type CacheEntry struct {
       Value    interface{}
       TTL      time.Duration
       AddedAt  time.Time
       LastUsed time.Time
       UseCount int64
   }
   ```

---

## 🏗️ Architecture

```
cache/
├── main.go              # Entry point
├── cache.go             # Core cache implementation
├── entry.go             # Cache entry definition
├── eviction.go          # Eviction policies
├── stats.go             # Statistics tracking
├── cleanup.go           # TTL cleanup
├── sync.go              # Synchronization primitives
└── network.go           # Network protocol (bonus)
```

---

## 📝 Implementation Steps

### Step 1: Define Cache Entry

```go
type CacheEntry struct {
    Value      interface{}
    TTL        time.Duration
    AddedAt    time.Time
    LastUsed   time.Time
    UseCount   int64
    Size       int64
}

func (e *CacheEntry) IsExpired() bool {
    if e.TTL == 0 {
        return false // No expiration
    }
    return time.Now().After(e.AddedAt.Add(e.TTL))
}
```

### Step 2: Create Core Cache

```go
type Cache struct {
    data       map[string]*CacheEntry
    mu         sync.RWMutex
    maxSize    int64
    currentSize int64

    stats      *Stats
    eviction   EvictionPolicy
}

func NewCache(maxSize int64, policy EvictionPolicy) *Cache {
    return &Cache{
        data:       make(map[string]*CacheEntry),
        maxSize:    maxSize,
        eviction:   policy,
        stats:      NewStats(),
    }
}
```

### Step 3: Implement Get/Set

```go
func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()

    entry, ok := c.data[key]
    if !ok {
        c.stats.RecordMiss()
        return nil, false
    }

    if entry.IsExpired() {
        c.stats.RecordMiss()
        return nil, false
    }

    entry.LastUsed = time.Now()
    entry.UseCount++
    c.stats.RecordHit()

    return entry.Value, true
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
    c.mu.Lock()
    defer c.mu.Unlock()

    size := int64(len(key) + estimateSize(value))

    // Make space if needed
    for c.currentSize+size > c.maxSize && len(c.data) > 0 {
        evictedKey := c.eviction.SelectVictim(c.data)
        c.evict(evictedKey)
    }

    entry := &CacheEntry{
        Value:    value,
        TTL:      ttl,
        AddedAt:  time.Now(),
        LastUsed: time.Now(),
        UseCount: 1,
        Size:     size,
    }

    c.data[key] = entry
    c.currentSize += size
    c.stats.RecordSet()

    return nil
}

func (c *Cache) evict(key string) {
    if entry, ok := c.data[key]; ok {
        delete(c.data, key)
        c.currentSize -= entry.Size
        c.stats.RecordEviction()
    }
}
```

### Step 4: Implement Eviction Policies

```go
type EvictionPolicy interface {
    SelectVictim(data map[string]*CacheEntry) string
}

// LRU Implementation
type LRUPolicy struct{}

func (p *LRUPolicy) SelectVictim(data map[string]*CacheEntry) string {
    var lruKey string
    var lruTime time.Time

    for key, entry := range data {
        if lruTime.IsZero() || entry.LastUsed.Before(lruTime) {
            lruKey = key
            lruTime = entry.LastUsed
        }
    }

    return lruKey
}

// LFU Implementation
type LFUPolicy struct{}

func (p *LFUPolicy) SelectVictim(data map[string]*CacheEntry) string {
    var lfuKey string
    var minCount int64 = math.MaxInt64

    for key, entry := range data {
        if entry.UseCount < minCount {
            lfuKey = key
            minCount = entry.UseCount
        }
    }

    return lfuKey
}
```

### Step 5: Add Statistics

```go
type Stats struct {
    hits       int64
    misses     int64
    sets       int64
    evictions  int64
    mu         sync.Mutex
}

func (s *Stats) RecordHit() {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.hits++
}

func (s *Stats) RecordMiss() {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.misses++
}

func (s *Stats) GetStats() map[string]interface{} {
    s.mu.Lock()
    defer s.mu.Unlock()

    total := s.hits + s.misses
    hitRate := 0.0
    if total > 0 {
        hitRate = float64(s.hits) / float64(total)
    }

    return map[string]interface{}{
        "hits":       s.hits,
        "misses":     s.misses,
        "hit_rate":   hitRate,
        "sets":       s.sets,
        "evictions":  s.evictions,
    }
}
```

### Step 6: Implement TTL Cleanup

```go
func (c *Cache) StartCleanup(interval time.Duration) {
    go func() {
        ticker := time.NewTicker(interval)
        defer ticker.Stop()

        for range ticker.C {
            c.cleanup()
        }
    }()
}

func (c *Cache) cleanup() {
    c.mu.Lock()
    defer c.mu.Unlock()

    expired := []string{}

    for key, entry := range c.data {
        if entry.IsExpired() {
            expired = append(expired, key)
        }
    }

    for _, key := range expired {
        c.evict(key)
    }
}
```

---

## ✅ Test Cases

```go
// Benchmark
func BenchmarkCacheGet(b *testing.B) {
    cache := NewCache(1000000, &LRUPolicy{})
    cache.Set("key", "value", 0)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        cache.Get("key")
    }
}

// Concurrency Test
func TestConcurrentAccess(t *testing.T) {
    cache := NewCache(1000000, &LRUPolicy{})

    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for j := 0; j < 1000; j++ {
                key := fmt.Sprintf("key-%d-%d", id, j)
                cache.Set(key, j, 0)
                cache.Get(key)
            }
        }(i)
    }

    wg.Wait()
}

// TTL Test
func TestTTLExpiration(t *testing.T) {
    cache := NewCache(1000, &LRUPolicy{})
    cache.Set("key", "value", 100*time.Millisecond)

    if _, ok := cache.Get("key"); !ok {
        t.Fatal("Key should exist")
    }

    time.Sleep(150 * time.Millisecond)

    if _, ok := cache.Get("key"); ok {
        t.Fatal("Key should be expired")
    }
}
```

---

## 📊 Performance Targets

- **Get latency**: < 1ms
- **Set latency**: < 1ms
- **Throughput**: > 100k ops/sec
- **Memory efficiency**: High (configurable eviction)
- **Hit rate**: > 80% (typical workload)

---

## 🌟 Bonus Features

- [ ] Network protocol (Redis-like)
- [ ] Persistence (RDB/AOF)
- [ ] Clustering & replication
- [ ] Pub/Sub support
- [ ] Lua scripting
- [ ] Stream support
- [ ] Transactions
- [ ] Bloom filters
- [ ] Consistent hashing
- [ ] Sharding

---

## 📚 Resources

- [Sync Package](https://pkg.go.dev/sync)
- [Container List](https://pkg.go.dev/container/list)
- [Cache Patterns](https://www.usenix.org/system/files/conference/nsdi13/nsdi13-final170.pdf)

---

## 💡 Learning Goals

✨ Advanced synchronization primitives (sync.Map, RWMutex)
✨ Reflection & dynamic programming
✨ Memory management & optimization
✨ Concurrent data structures
✨ TTL and expiration handling
✨ Eviction policies & algorithms
✨ Statistics collection & monitoring
✨ Production-grade system design
