# Project 2: Real-time Chat Server

## 🎯 Mục tiêu

Xây dựng một **WebSocket chat server** với:

- ✅ Multiple chat rooms
- ✅ Message broadcasting đến tất cả users trong room
- ✅ Presence tracking (join/leave notifications)
- ✅ Message history (in-memory, last 50 messages per room)
- ✅ Username support
- ✅ Graceful shutdown
- ✅ Simple HTML client để test

---

## 📋 Yêu cầu

### Message Types

```go
// Client → Server
type ClientMessage struct {
    Type    string `json:"type"`    // "chat", "join", "leave", "ping"
    Room    string `json:"room"`
    Content string `json:"content"`
}

// Server → Client
type ServerMessage struct {
    Type      string    `json:"type"`    // "chat", "system", "history", "presence", "error"
    Room      string    `json:"room"`
    Sender    string    `json:"sender"`
    Content   string    `json:"content"`
    Timestamp time.Time `json:"timestamp"`
    UserCount int       `json:"user_count,omitempty"`
}
```

### Features cần implement

1. **Rooms**
   - Tạo room tự động khi user đầu tiên join
   - Xoá room khi user cuối cùng rời
   - Có thể join nhiều rooms

2. **Broadcasting**
   - Khi user gửi chat → broadcast đến tất cả trong room
   - Khi user join/leave → thông báo presence cho room

3. **Message History**
   - Lưu 50 messages gần nhất cho mỗi room
   - Khi user join room → gửi history cho user đó

4. **Presence**
   - Track danh sách user đang online trong mỗi room
   - Endpoint `GET /rooms` trả danh sách rooms + user count

5. **Connection Management**
   - Ping/pong để detect dead connections
   - Read/write timeout
   - Graceful disconnect

6. **API Endpoints**
   - `GET /` - Serve HTML client
   - `GET /ws?room=general&username=Alice` - WebSocket endpoint
   - `GET /rooms` - List active rooms (JSON)

---

## 🏗️ Architecture

```
realtime_chat/
├── hub/
│   ├── hub.go         # Hub: manages all connections & rooms
│   ├── client.go      # Client: represents one WS connection
│   └── room.go        # Room: manages messages & members
├── handler/
│   └── ws_handler.go  # HTTP/WS handlers
├── static/
│   └── index.html     # Simple chat UI
├── go.mod
└── main.go
```

### Hub Design

```go
type Hub struct {
    mu         sync.RWMutex
    rooms      map[string]*Room     // roomName → Room
    register   chan *Client
    unregister chan *Client
    broadcast  chan BroadcastReq
}

type Room struct {
    name     string
    clients  map[*Client]struct{}
    history  []ServerMessage        // circular buffer, max 50
    mu       sync.RWMutex
}

type Client struct {
    id       string
    username string
    room     string
    hub      *Hub
    conn     *websocket.Conn
    send     chan ServerMessage
}
```

---

## 🚀 Implementation Steps

### Step 1: Setup

```bash
mkdir realtime_chat ; cd realtime_chat
go mod init realtime_chat
go get github.com/gorilla/websocket
```

### Step 2: Implement Hub

- `Hub.Run()` goroutine: xử lý register/unregister/broadcast channels
- `Room.AddMessage()`: thêm vào circular buffer (max 50)
- `Room.GetHistory()`: trả messages gần nhất

### Step 3: Implement Client

- `client.readPump()`: đọc messages, gửi vào Hub's broadcast channel
- `client.writePump()`: ghi messages từ `send` channel vào WebSocket
- Ping/pong heartbeat

### Step 4: HTTP Handlers

```go
func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    room := r.URL.Query().Get("room")
    // ... upgrade và tạo client
}

func ListRooms(hub *Hub, w http.ResponseWriter, r *http.Request) {
    // trả JSON list of rooms với user count
}
```

### Step 5: HTML Client

Tạo `static/index.html` với WebSocket client JavaScript:

```html
<!-- Kết nối WS, hiển thị messages, input để gửi -->
```

### Step 6: Test

```bash
go run main.go
# Mở browser: http://localhost:8080
# Mở nhiều tabs → kiểm tra broadcast, presence
```

---

## 📊 Expected Behavior

```
# User Alice joins room "general"
Server → Alice: [HISTORY] Last 3 messages...
Server → room: [SYSTEM] Alice joined the room (5 users)

# Alice sends a message
Server → room: [CHAT] Alice: Hello everyone!

# User Bob leaves
Server → room: [SYSTEM] Bob left the room (4 users)
```

### API: `GET /rooms`

```json
[
  { "name": "general", "userCount": 5, "messageCount": 142 },
  { "name": "golang", "userCount": 2, "messageCount": 38 }
]
```

---

## 🌟 Bonus Features

- [ ] Private messages (whisper): `@username message`
- [ ] Message reactions (emoji)
- [ ] Rate limiting per connection (max 10 msg/sec)
- [ ] Persistent storage: save messages to file hoặc SQLite
- [ ] SSE fallback endpoint cho clients không hỗ trợ WebSocket
- [ ] Room admin: kick user, ban, set topic

---

## ✅ Tiêu chí hoàn thành

- [ ] Multiple users có thể chat trong cùng room
- [ ] Presence notifications (join/leave) hoạt động
- [ ] Message history gửi khi user join
- [ ] Dead connections được cleanup (ping/pong)
- [ ] `/rooms` endpoint trả đúng thông tin
- [ ] Graceful shutdown: đóng tất cả connections sạch sẽ
- [ ] `go test -race ./...` không có data races
