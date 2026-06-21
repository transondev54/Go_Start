# Lesson 6: JSON Marshaling & API Calls

## 📖 Nội dung bài học

1. JSON encoding & decoding
2. Struct tags
3. HTTP requests
4. Parsing API responses
5. Error handling

---

## 1️⃣ JSON MARSHALING

### Encode (Struct → JSON)

```go
import \"encoding/json\"

type Person struct {
    Name  string
    Age   int
    Email string
}

func main() {
    person := Person{
        Name:  \"John Doe\",
        Age:   30,
        Email: \"john@example.com\",
    }

    // Marshal to JSON
    jsonData, err := json.Marshal(person)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(jsonData))
    // Output: {\"Name\":\"John Doe\",\"Age\":30,\"Email\":\"john@example.com\"}
}
```

### Decode (JSON → Struct)

```go
func main() {
    jsonString := `{\"Name\":\"Jane\",\"Age\":25,\"Email\":\"jane@example.com\"}`

    var person Person
    err := json.Unmarshal([]byte(jsonString), &person)
    if err != nil {
        panic(err)
    }

    fmt.Printf(\"%+v\\n\", person)
    // Output: {Name:Jane Age:25 Email:jane@example.com}
}
```

---

## 2️⃣ STRUCT TAGS

### JSON tags

```go
type User struct {
    // Mapping JSON field name
    ID        int    `json:\"id\"`
    Name      string `json:\"name\"`
    Email     string `json:\"email\"`
    Password  string `json:\"-\"`  // Ignore này khi marshal
}

// Với tags:
// Input JSON: {\"id\":1,\"name\":\"John\",\"email\":\"john@example.com\"}
// Output: User{ID:1, Name:\"John\", Email:\"john@example.com\", Password:\"\"}
```

### Advanced tags

```go
type Product struct {
    ID       int     `json:\"id\"`
    Name     string  `json:\"name\"`
    Price    float64 `json:\"price,string\"`  // Convert string to float
    InStock  bool    `json:\"in_stock,omitempty\"`  // Omit if empty
    Category string  `json:\"category\"`
}
```

---

## 3️⃣ HTTP REQUESTS

### GET request

```go
import \"net/http\"

func main() {
    // GET request
    resp, err := http.Get(\"https://api.example.com/users/1\")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    // Đọc response body
    body, _ := io.ReadAll(resp.Body)

    // Parse JSON
    var user User
    json.Unmarshal(body, &user)

    fmt.Printf(\"%+v\\n\", user)
}
```

### POST request

```go
func main() {
    user := User{
        Name:  \"John\",
        Email: \"john@example.com\",
    }

    // Convert struct to JSON
    jsonData, _ := json.Marshal(user)

    // POST request
    resp, err := http.Post(
        \"https://api.example.com/users\",
        \"application/json\",
        bytes.NewBuffer(jsonData),
    )
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println(resp.Status)
}
```

### Custom requests (PUT, DELETE)

```go
func main() {
    client := &http.Client{}

    // PUT request
    req, _ := http.NewRequest(\"PUT\", \"https://api.example.com/users/1\", nil)
    req.Header.Set(\"Authorization\", \"Bearer token123\")

    resp, _ := client.Do(req)
    defer resp.Body.Close()
}
```

---

## 4️⃣ PARSING API RESPONSES

### Simple API response

```go
type GitHubUser struct {
    Login string `json:\"login\"`
    ID    int    `json:\"id\"`
    Repos int    `json:\"public_repos\"`
}

func GetGitHubUser(username string) (*GitHubUser, error) {
    url := fmt.Sprintf(\"https://api.github.com/users/%s\", username)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Check status
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf(\"API error: %d\", resp.StatusCode)
    }

    body, _ := io.ReadAll(resp.Body)

    var user GitHubUser
    if err := json.Unmarshal(body, &user); err != nil {
        return nil, err
    }

    return &user, nil
}
```

### Nested structures

```go
type WeatherResponse struct {
    City string `json:\"name\"`
    Main struct {
        Temp      float64 `json:\"temp\"`
        Humidity  int     `json:\"humidity\"`
    } `json:\"main\"`
    Weather []struct {
        Description string `json:\"description\"`
    } `json:\"weather\"`
}

func GetWeather(city string) (*WeatherResponse, error) {
    // API call...
    var weather WeatherResponse
    json.Unmarshal(body, &weather)
    return &weather, nil
}
```

---

## 5️⃣ ERROR HANDLING

### Status codes

```go
func MakeRequest(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf(\"request failed: %w\", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusNotFound {
        return nil, fmt.Errorf(\"resource not found\")
    }

    if resp.StatusCode >= 400 {
        return nil, fmt.Errorf(\"API error: %d\", resp.StatusCode)
    }

    return io.ReadAll(resp.Body)
}
```

---

## 📝 TÓM TẮT JSON & HTTP

```go
// Marshal (Struct → JSON)
jsonData, _ := json.Marshal(obj)

// Unmarshal (JSON → Struct)
json.Unmarshal([]byte(jsonString), &obj)

// HTTP GET
resp, _ := http.Get(url)

// HTTP POST
http.Post(url, \"application/json\", buffer)
```

---

## 💡 BEST PRACTICES

1. **Luôn kiểm tra error** - requests có thể fail
2. **Kiểm tra status code** - không phải tất cả 200 là OK
3. **Defer resp.Body.Close()** - tránh leak
4. **Dùng custom HTTP client** - timeout, headers, v.v.
5. **Validate JSON structure** - trước khi sử dụng

---

## 🎯 BƯỚC TIẾP THEO

- Thực hành JSON marshaling
- Thử gọi public APIs
- Sử dụng trong mini project
