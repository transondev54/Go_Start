# Lesson 7: Security Best Practices

## 📖 Nội dung bài học

1. Input validation
2. SQL injection prevention
3. Secure password handling
4. HTTPS & TLS
5. Encryption
6. Common vulnerabilities

---

## 1️⃣ INPUT VALIDATION

### Whitelist Validation

```go
// ✅ Đúng: Whitelist pattern
func ValidateUsername(username string) error {
    if !regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`).MatchString(username) {
        return fmt.Errorf("invalid username")
    }
    return nil
}

// Usage
err := ValidateUsername(username)
if err != nil {
    http.Error(w, "Invalid input", http.StatusBadRequest)
    return
}
```

### Length Limits

```go
// ❌ Sai: No length check
func ProcessInput(input string) {
    // Attacker can send huge input
}

// ✅ Đúng: Check length
func ProcessInput(input string) error {
    const maxLen = 1000
    if len(input) > maxLen {
        return fmt.Errorf("input too long")
    }
    // Process input
    return nil
}
```

### Type Validation

```go
// ✅ Đúng: Validate type
func ValidateAge(ageStr string) (int, error) {
    age, err := strconv.Atoi(ageStr)
    if err != nil {
        return 0, fmt.Errorf("invalid age")
    }

    if age < 0 || age > 150 {
        return 0, fmt.Errorf("age out of range")
    }

    return age, nil
}
```

---

## 2️⃣ SQL INJECTION PREVENTION

### Parameterized Queries

```go
import "database/sql"

// ❌ Sai: SQL injection vulnerable
func GetUser(db *sql.DB, userID string) {
    query := fmt.Sprintf("SELECT * FROM users WHERE id = %s", userID)
    // Attacker: userID = "1 OR 1=1"
}

// ✅ Đúng: Use parameterized queries
func GetUser(db *sql.DB, userID string) (User, error) {
    var user User
    query := "SELECT id, name, email FROM users WHERE id = ?"
    err := db.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email)
    return user, err
}
```

### Prepared Statements

```go
// ✅ Đúng: Prepared statement
func InsertUser(db *sql.DB, name, email string) error {
    stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(name, email)
    return err
}
```

---

## 3️⃣ SECURE PASSWORD HANDLING

### Password Hashing

```go
import "golang.org/x/crypto/bcrypt"

// ✅ Đúng: Hash password
func HashPassword(password string) (string, error) {
    // Cost indicates how many times hash is computed
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash), err
}

// ✅ Đúng: Verify password
func VerifyPassword(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword(
        []byte(hashedPassword),
        []byte(password),
    )
    return err == nil
}
```

### Secure Random Tokens

```go
import "crypto/rand"

func GenerateToken(length int) (string, error) {
    bytes := make([]byte, length)
    _, err := rand.Read(bytes)
    if err != nil {
        return "", err
    }

    // Convert to hex
    return hex.EncodeToString(bytes), nil
}
```

---

## 4️⃣ HTTPS & TLS

### HTTPS Server

```go
func StartServer() error {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handleRequest)

    // Use HTTPS
    return http.ListenAndServeTLS(
        ":443",
        "cert.pem",
        "key.pem",
        mux,
    )
}
```

### TLS Configuration

```go
func StartSecureServer() error {
    tlsConfig := &tls.Config{
        // Require TLS 1.2 or higher
        MinVersion: tls.VersionTLS12,

        // Use strong cipher suites
        CipherSuites: []uint16{
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
        },
    }

    server := &http.Server{
        Addr:      ":443",
        Handler:   http.DefaultServeMux,
        TLSConfig: tlsConfig,
    }

    return server.ListenAndServeTLS("cert.pem", "key.pem")
}
```

---

## 5️⃣ ENCRYPTION

### Symmetric Encryption (AES)

```go
import "crypto/aes"
import "crypto/cipher"
import "crypto/rand"

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    // Generate random IV
    iv := make([]byte, aes.BlockSize)
    if _, err := rand.Read(iv); err != nil {
        return nil, err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    ciphertext := make([]byte, len(plaintext))
    stream.XORKeyStream(ciphertext, plaintext)

    // Prepend IV to ciphertext
    return append(iv, ciphertext...), nil
}

func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    // Extract IV
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    plaintext := make([]byte, len(ciphertext))
    stream.XORKeyStream(plaintext, ciphertext)

    return plaintext, nil
}
```

---

## 6️⃣ COMMON VULNERABILITIES

### Cross-Site Scripting (XSS)

```go
import "html"

// ❌ Sai: Vulnerable to XSS
func OutputUser(w http.ResponseWriter, userName string) {
    fmt.Fprintf(w, "<div>User: %s</div>", userName)
    // If userName = "<script>alert('hacked')</script>"
}

// ✅ Đúng: Escape HTML
func OutputUser(w http.ResponseWriter, userName string) {
    fmt.Fprintf(w, "<div>User: %s</div>", html.EscapeString(userName))
}
```

### Cross-Site Request Forgery (CSRF)

```go
import "gorilla/csrf"

func ProtectForms(next http.Handler) http.Handler {
    return csrf.Protect([]byte("auth-key"))(next)
}

// In template
// {{ csrf.TemplateField }}
```

### Rate Limiting

```go
import "golang.org/x/time/rate"

func RateLimitMiddleware(limiter *rate.Limiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if !limiter.Allow() {
                http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
```

---

## 7️⃣ SECURITY CHECKLIST

- ✅ Validate all input
- ✅ Use parameterized queries
- ✅ Hash passwords with bcrypt
- ✅ Use HTTPS/TLS
- ✅ Escape output (prevent XSS)
- ✅ Implement CSRF protection
- ✅ Use secure random for tokens
- ✅ Implement rate limiting
- ✅ Sanitize file uploads
- ✅ Keep dependencies updated

---

## 📝 EXERCISES

1. **Input Validation**: Implement regex-based validation
2. **Password Hashing**: Implement login system with bcrypt
3. **HTTPS Server**: Create HTTPS server with proper TLS config

---

## 📚 RESOURCES

- [Go Security](https://golang.org/wiki/SecurityPolicy)
- [Crypto Package](https://pkg.go.dev/crypto)
- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
