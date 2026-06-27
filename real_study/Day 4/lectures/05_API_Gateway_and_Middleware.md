# Lesson 5: API Gateway & Middleware

## 📖 Nội dung bài học

1. API Gateway pattern
2. Middleware chain pattern
3. Rate Limiting
4. JWT Authentication Middleware
5. Request validation & sanitization
6. CORS handling

---

## 1️⃣ API GATEWAY PATTERN

API Gateway là entry point duy nhất cho tất cả client requests, xử lý cross-cutting concerns.

```
Client ──→ API Gateway ──→ OrderService
                      ──→ UserService
                      ──→ InventoryService
                      ──→ NotificationService

Gateway xử lý:
  - Authentication / Authorization
  - Rate limiting
  - Request logging
  - CORS
  - Request/Response transformation
  - Circuit breaking
  - Load balancing
```

---

## 2️⃣ MIDDLEWARE CHAIN PATTERN

```go
// Middleware type
type Middleware func(http.Handler) http.Handler

// Chain kết hợp nhiều middlewares
func Chain(middlewares ...Middleware) Middleware {
    return func(final http.Handler) http.Handler {
        for i := len(middlewares) - 1; i >= 0; i-- {
            final = middlewares[i](final)
        }
        return final
    }
}

// Sử dụng
chain := Chain(
    LoggingMiddleware,
    RecoveryMiddleware,
    CORSMiddleware,
    AuthMiddleware,
    RateLimitMiddleware,
)

mux := http.NewServeMux()
mux.Handle("/api/orders", chain(ordersHandler))
```

### Recovery Middleware

```go
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if rec := recover(); rec != nil {
                log.Error().
                    Interface("panic", rec).
                    Str("stack", string(debug.Stack())).
                    Msg("panic recovered")

                w.Header().Set("Content-Type", "application/json")
                w.WriteHeader(http.StatusInternalServerError)
                json.NewEncoder(w).Encode(map[string]string{
                    "error": "internal server error",
                })
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

---

## 3️⃣ RATE LIMITING

### Token Bucket algorithm

```go
package ratelimit

import (
    "net/http"
    "sync"
    "time"
)

type TokenBucket struct {
    mu         sync.Mutex
    tokens     float64
    maxTokens  float64
    refillRate float64 // tokens per second
    lastRefill time.Time
}

func NewTokenBucket(maxTokens, refillRate float64) *TokenBucket {
    return &TokenBucket{
        tokens:     maxTokens,
        maxTokens:  maxTokens,
        refillRate: refillRate,
        lastRefill: time.Now(),
    }
}

func (tb *TokenBucket) Allow() bool {
    tb.mu.Lock()
    defer tb.mu.Unlock()

    // Refill tokens based on elapsed time
    now := time.Now()
    elapsed := now.Sub(tb.lastRefill).Seconds()
    tb.tokens = min(tb.maxTokens, tb.tokens+elapsed*tb.refillRate)
    tb.lastRefill = now

    if tb.tokens >= 1 {
        tb.tokens--
        return true
    }
    return false
}

func min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}

// Per-IP rate limiter
type IPRateLimiter struct {
    mu       sync.RWMutex
    limiters map[string]*TokenBucket
    rate     float64
    capacity float64
}

func NewIPRateLimiter(rate, capacity float64) *IPRateLimiter {
    rl := &IPRateLimiter{
        limiters: make(map[string]*TokenBucket),
        rate:     rate,
        capacity: capacity,
    }
    // Cleanup goroutine
    go rl.cleanup()
    return rl
}

func (rl *IPRateLimiter) GetLimiter(ip string) *TokenBucket {
    rl.mu.RLock()
    limiter, ok := rl.limiters[ip]
    rl.mu.RUnlock()

    if !ok {
        rl.mu.Lock()
        limiter = NewTokenBucket(rl.capacity, rl.rate)
        rl.limiters[ip] = limiter
        rl.mu.Unlock()
    }
    return limiter
}

func (rl *IPRateLimiter) cleanup() {
    ticker := time.NewTicker(5 * time.Minute)
    for range ticker.C {
        rl.mu.Lock()
        // Simple cleanup: clear all (in production, use TTL-based cleanup)
        rl.limiters = make(map[string]*TokenBucket)
        rl.mu.Unlock()
    }
}

// Middleware
func RateLimitMiddleware(rl *IPRateLimiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ip := extractIP(r)
            if !rl.GetLimiter(ip).Allow() {
                w.Header().Set("Retry-After", "1")
                http.Error(w, `{"error":"rate limit exceeded"}`, http.StatusTooManyRequests)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}

func extractIP(r *http.Request) string {
    // Check X-Forwarded-For for proxy setups
    if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
        // Take first IP (closest to client)
        if idx := strings.Index(xff, ","); idx != -1 {
            return strings.TrimSpace(xff[:idx])
        }
        return xff
    }
    // Fall back to RemoteAddr
    host, _, _ := net.SplitHostPort(r.RemoteAddr)
    return host
}
```

---

## 4️⃣ JWT AUTHENTICATION MIDDLEWARE

```go
package auth

import (
    "context"
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID string `json:"user_id"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

type contextKey string

const ClaimsKey contextKey = "claims"

// GenerateToken tạo JWT token
func GenerateToken(userID, role, secret string, ttl time.Duration) (string, error) {
    claims := Claims{
        UserID: userID,
        Role:   role,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}

// ParseToken validate và parse JWT
func ParseToken(tokenStr, secret string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(
        tokenStr,
        &Claims{},
        func(t *jwt.Token) (interface{}, error) {
            if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
            }
            return []byte(secret), nil
        },
    )
    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }
    return claims, nil
}

// AuthMiddleware kiểm tra Bearer token
func AuthMiddleware(secret string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            header := r.Header.Get("Authorization")
            if !strings.HasPrefix(header, "Bearer ") {
                http.Error(w, `{"error":"missing authorization"}`, http.StatusUnauthorized)
                return
            }

            tokenStr := strings.TrimPrefix(header, "Bearer ")
            claims, err := ParseToken(tokenStr, secret)
            if err != nil {
                http.Error(w, `{"error":"invalid token"}`, http.StatusUnauthorized)
                return
            }

            // Inject claims vào context
            ctx := context.WithValue(r.Context(), ClaimsKey, claims)
            next.ServeHTTP(w, r.WithContext(ctx))
        })
    }
}

// RequireRole kiểm tra role
func RequireRole(role string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            claims, ok := r.Context().Value(ClaimsKey).(*Claims)
            if !ok || claims.Role != role {
                http.Error(w, `{"error":"forbidden"}`, http.StatusForbidden)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
```

---

## 5️⃣ CORS MIDDLEWARE

```go
func CORSMiddleware(allowedOrigins []string) func(http.Handler) http.Handler {
    allowed := make(map[string]bool)
    for _, o := range allowedOrigins {
        allowed[o] = true
    }

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            origin := r.Header.Get("Origin")

            if allowed[origin] || allowed["*"] {
                w.Header().Set("Access-Control-Allow-Origin", origin)
                w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
                w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
                w.Header().Set("Access-Control-Max-Age", "3600")
            }

            // Handle preflight
            if r.Method == http.MethodOptions {
                w.WriteHeader(http.StatusNoContent)
                return
            }

            next.ServeHTTP(w, r)
        })
    }
}
```

---

## 🧠 QUIZ - 5 CÂU HỎI

1. Token Bucket rate limiting khác Fixed Window rate limiting như thế nào?
2. Tại sao nên dùng `X-Forwarded-For` header thay vì `RemoteAddr` để lấy client IP?
3. JWT `Claims` có thể bị giả mạo không? Tại sao có/không?
4. Middleware chain xử lý theo thứ tự nào (request và response)?
5. CORS preflight request là gì và tại sao browser gửi nó?

---

## 📌 KEY TAKEAWAYS

- Middleware chain = composable cross-cutting concerns
- Token Bucket cho phép burst requests, Fixed Window thì không
- JWT: stateless auth, dùng secret key để sign và verify
- Luôn validate JWT signing method để tránh "none" algorithm attack
- CORS cần xử lý OPTIONS preflight requests riêng
