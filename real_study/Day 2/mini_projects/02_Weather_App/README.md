# Mini Project 2: Weather App with Goroutines

## 📝 Mô tả

Xây dựng một weather app sử dụng:

- OpenWeatherMap API (hoặc mock API)
- Goroutines để fetch dữ liệu từ multiple cities
- Channels để communicate giữa goroutines
- Caching để tránh duplicate API calls
- JSON parsing

---

## 📋 Yêu cầu

### Tính năng bắt buộc

1. **Weather struct**:

   ```go
   type Weather struct {
       City       string
       Temp       float64
       Humidity   int
       Description string
       WindSpeed  float64
   }
   ```

2. **API Integration**:
   - Fetch từ OpenWeatherMap (hoặc mock)
   - Parse JSON response
   - Handle API errors

3. **Concurrency**:
   - Fetch weather cho multiple cities **concurrently** (goroutines)
   - Sử dụng channels để collect results
   - Timeout để tránh hang

4. **Caching**:
   - Lưu weather data in-memory
   - Return cached data nếu fetch gần đây
   - Có TTL (time-to-live) cho cache

5. **Display**:
   - Format output đẹp
   - Show tất cả 4-5 cities
   - Error messages rõ ràng

### Ví dụ output

```
╔═══════════════════════════════════════╗
║        Weather App v1.0               ║
╚═══════════════════════════════════════╝

Fetching weather for: Hanoi, Bangkok, Tokyo, Seoul, Singapore...

🌍 Weather Report
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🌏 Hanoi, Vietnam
   Temperature: 28°C
   Humidity: 72%
   Condition: Partly Cloudy
   Wind Speed: 3.2 m/s

🌏 Bangkok, Thailand
   Temperature: 32°C
   Humidity: 65%
   Condition: Sunny
   Wind Speed: 2.1 m/s

... (3 cities more)
```

---

## 🎯 Learning Objectives

- ✅ Goroutines & concurrency
- ✅ Channels & communication patterns
- ✅ HTTP requests (GET)
- ✅ JSON unmarshaling
- ✅ Error handling với goroutines
- ✅ Caching pattern

---

## 📚 Bước thực hiện

### Bước 1: Setup

```bash
mkdir weather_app
cd weather_app
go mod init weather_app
code main.go
```

### Bước 2: Define structs

```go
type Weather struct {
    City        string
    Temp        float64
    Humidity    int
    Description string
    WindSpeed   float64
}

type WeatherService struct {
    cache map[string]*CachedWeather
    mu    sync.Mutex
}

type CachedWeather struct {
    data      *Weather
    timestamp time.Time
}
```

### Bước 3: Mock API (nếu không muốn real API)

```go
func FetchWeatherMock(city string) (*Weather, error) {
    // Return mock data tùy vào city
    weatherData := map[string]Weather{
        "Hanoi": {City: "Hanoi", Temp: 28, Humidity: 72, ...},
        "Tokyo": {City: "Tokyo", Temp: 15, Humidity: 55, ...},
    }
    // ...
}
```

### Bước 4: Fetch with Goroutines

```go
func (ws *WeatherService) FetchMultipleCities(cities []string) []*Weather {
    results := make(chan *Weather, len(cities))

    for _, city := range cities {
        go func(c string) {
            weather, _ := FetchWeather(c)
            results <- weather
        }(city)
    }

    var weathers []*Weather
    for i := 0; i < len(cities); i++ {
        weathers = append(weathers, <-results)
    }

    return weathers
}
```

### Bước 5: Caching

```go
func (ws *WeatherService) GetWithCache(city string) (*Weather, error) {
    ws.mu.Lock()
    if cached, ok := ws.cache[city]; ok {
        if time.Since(cached.timestamp) < 10*time.Minute {
            ws.mu.Unlock()
            return cached.data, nil
        }
    }
    ws.mu.Unlock()

    // Fetch from API
    weather, err := FetchWeather(city)
    if err != nil {
        return nil, err
    }

    ws.mu.Lock()
    ws.cache[city] = &CachedWeather{
        data:      weather,
        timestamp: time.Now(),
    }
    ws.mu.Unlock()

    return weather, nil
}
```

### Bước 6: Display results

```go
func DisplayWeather(weathers []*Weather) {
    for _, w := range weathers {
        fmt.Printf("🌏 %s\n", w.City)
        fmt.Printf("   Temperature: %.1f°C\n", w.Temp)
        fmt.Printf("   Humidity: %d%%\n", w.Humidity)
        // ...
    }
}
```

---

## 📦 API Options

### Option 1: OpenWeatherMap (Real API)

```
https://api.openweathermap.org/data/2.5/weather?q=Hanoi&appid=YOUR_API_KEY
```

Sign up: https://openweathermap.org/api

### Option 2: Mock Data (No API key needed)

```go
func FetchWeatherMock(city string) (*Weather, error) {
    // Return hardcoded data
}
```

---

## 📦 Bonus Features

- [ ] Search by coordinates (latitude/longitude)
- [ ] Forecast (5-day, hourly)
- [ ] Save favorites
- [ ] UV index
- [ ] Air quality
- [ ] Historical data
- [ ] Temperature unit conversion (C/F)
- [ ] Auto-refresh every N minutes

---

## ✅ Checklist

- [ ] Project setup
- [ ] Structs defined
- [ ] Mock/Real API working
- [ ] Single city fetch works
- [ ] Multiple cities with goroutines
- [ ] Results collected via channels
- [ ] Caching implemented
- [ ] Error handling
- [ ] Display formatted
- [ ] Tested (goroutine safety)

---

## 🔍 Test Scenarios

```
1. Fetch single city → Verify data
2. Fetch multiple cities concurrently → Verify all loaded
3. Cache hit → Verify uses cached data
4. Cache miss → Verify fetches fresh data
5. API timeout → Graceful error
6. Invalid city → Error message
```

---

## 📊 Scoring Rubric (0-100)

- **Functionality (40%)**: Concurrent fetching, caching
- **Code Quality (30%)**: Goroutine safety, error handling
- **Concurrency (20%)**: Proper use of channels, sync
- **Testing (10%)**: Verify concurrent behavior

---

## 🔗 Resources

- Goroutines: [sync package](https://pkg.go.dev/sync)
- Channels: [Effective Go](https://golang.org/doc/effective_go#concurrency)
- HTTP Client: [net/http](https://pkg.go.dev/net/http)
