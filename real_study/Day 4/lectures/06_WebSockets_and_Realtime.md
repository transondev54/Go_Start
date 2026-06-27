# Lesson 6: WebSockets & Real-time Communication

## 📖 Nội dung bài học

1. WebSocket protocol là gì?
2. WebSocket server với `gorilla/websocket`
3. Hub pattern - quản lý connections
4. Message broadcasting & rooms
5. Presence tracking
6. Server-Sent Events (SSE) - alternative

---

## 1️⃣ WEBSOCKET LÀ GÌ?

**WebSocket** là protocol cho phép bi-directional, full-duplex communication qua một TCP connection duy nhất.

### HTTP vs WebSocket

```
HTTP (Request-Response):
  Client → Server: "GET /data"
  Server → Client: "200 OK {...}"
  (Connection đóng)

WebSocket (Persistent):
  Client → Server: Handshake (HTTP Upgrade)
  Server → Client: 101 Switching Protocols
  ─────────────────────────────────────────
  Client ↔ Server: messages bất kỳ lúc nào
  Client ↔ Server: messages bất kỳ lúc nào
  (Connection mở cho đến khi ai đó đóng)
```

### Cài đặt

```bash
go get github.com/gorilla/websocket
```

---

## 2️⃣ WEBSOCKET SERVER CƠ BẢN

```go
package main

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    // Kiểm tra origin (security!) - trong production phải validate origin
    CheckOrigin: func(r *http.Request) bool {
        // TODO: check r.Header.Get("Origin") against allowed origins
        return true // tạm thời allow all
    },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
    // Upgrade HTTP → WebSocket
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("upgrade error: %v", err)
        return
    }
    defer conn.Close()

    // Echo server: đọc message và gửi lại
    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err,
                websocket.CloseGoingAway,
                websocket.CloseAbnormalClosure,
            ) {
                log.Printf("read error: %v", err)
            }
            break
        }
        log.Printf("received: %s", message)

        // Echo back
        if err := conn.WriteMessage(messageType, message); err != nil {
            log.Printf("write error: %v", err)
            break
        }
    }
}

func main() {
    http.HandleFunc("/ws", wsHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

## 3️⃣ HUB PATTERN - QUẢN LÝ NHIỀU CONNECTIONS

```go
package hub

import (
    "encoding/json"
    "log"
    "sync"

    "github.com/gorilla/websocket"
)

// Message gửi qua WebSocket
type Message struct {
    Type    string          `json:"type"`
    Room    string          `json:"room"`
    Sender  string          `json:"sender"`
    Content string          `json:"content"`
}

// Client đại diện cho một WebSocket connection
type Client struct {
    id   string
    hub  *Hub
    conn *websocket.Conn
    send chan []byte
    room string
}

// Hub quản lý tất cả clients
type Hub struct {
    mu         sync.RWMutex
    clients    map[string]*Client       // clientID → client
    rooms      map[string]map[string]*Client // roomID → {clientID → client}
    register   chan *Client
    unregister chan *Client
    broadcast  chan broadcastMsg
}

type broadcastMsg struct {
    room    string
    message []byte
    exclude string // exclude this client ID
}

func New() *Hub {
    return &Hub{
        clients:    make(map[string]*Client),
        rooms:      make(map[string]map[string]*Client),
        register:   make(chan *Client, 256),
        unregister: make(chan *Client, 256),
        broadcast:  make(chan broadcastMsg, 256),
    }
}

// Run xử lý events của hub (chạy trong goroutine riêng)
func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            h.mu.Lock()
            h.clients[client.id] = client
            if _, ok := h.rooms[client.room]; !ok {
                h.rooms[client.room] = make(map[string]*Client)
            }
            h.rooms[client.room][client.id] = client
            h.mu.Unlock()
            log.Printf("client %s joined room %s", client.id, client.room)

        case client := <-h.unregister:
            h.mu.Lock()
            if _, ok := h.clients[client.id]; ok {
                delete(h.clients, client.id)
                delete(h.rooms[client.room], client.id)
                close(client.send)
            }
            h.mu.Unlock()
            log.Printf("client %s left room %s", client.id, client.room)

        case msg := <-h.broadcast:
            h.mu.RLock()
            roomClients := h.rooms[msg.room]
            h.mu.RUnlock()

            for id, client := range roomClients {
                if id == msg.exclude {
                    continue
                }
                select {
                case client.send <- msg.message:
                default:
                    // Client send buffer full → disconnect
                    close(client.send)
                    h.mu.Lock()
                    delete(h.clients, id)
                    delete(h.rooms[msg.room], id)
                    h.mu.Unlock()
                }
            }
        }
    }
}

// Broadcast gửi message đến tất cả trong room
func (h *Hub) Broadcast(room string, msg []byte, excludeID string) {
    h.broadcast <- broadcastMsg{room: room, message: msg, exclude: excludeID}
}

// RoomCount số clients trong room
func (h *Hub) RoomCount(room string) int {
    h.mu.RLock()
    defer h.mu.RUnlock()
    return len(h.rooms[room])
}
```

### Client goroutines (read/write pumps)

```go
// readPump đọc messages từ WebSocket
func (c *Client) readPump() {
    defer func() {
        c.hub.unregister <- c
        c.conn.Close()
    }()

    c.conn.SetReadLimit(512 * 1024) // 512KB max message
    c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
    c.conn.SetPongHandler(func(string) error {
        c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
        return nil
    })

    for {
        _, rawMsg, err := c.conn.ReadMessage()
        if err != nil {
            break
        }

        var msg Message
        if err := json.Unmarshal(rawMsg, &msg); err != nil {
            continue
        }
        msg.Sender = c.id

        data, _ := json.Marshal(msg)
        c.hub.Broadcast(c.room, data, "") // broadcast to all including sender
    }
}

// writePump ghi messages vào WebSocket
func (c *Client) writePump() {
    ticker := time.NewTicker(30 * time.Second) // ping interval
    defer func() {
        ticker.Stop()
        c.conn.Close()
    }()

    for {
        select {
        case message, ok := <-c.send:
            c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if !ok {
                c.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
                return
            }

        case <-ticker.C:
            c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                return
            }
        }
    }
}
```

---

## 4️⃣ WEBSOCKET HANDLER

```go
func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
    room := r.URL.Query().Get("room")
    if room == "" {
        room = "general"
    }

    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        return
    }

    client := &Client{
        id:   generateID(),
        hub:  hub,
        conn: conn,
        send: make(chan []byte, 256),
        room: room,
    }

    hub.register <- client

    // Start pumps in goroutines
    go client.writePump()
    go client.readPump()

    // Send welcome message
    welcome, _ := json.Marshal(Message{
        Type:    "system",
        Room:    room,
        Content: fmt.Sprintf("Welcome! %d users in room %s", hub.RoomCount(room), room),
    })
    client.send <- welcome
}
```

---

## 5️⃣ SERVER-SENT EVENTS (SSE)

SSE là alternative một chiều (server → client only):

```go
// SSE handler - server push updates
func SSEHandler(w http.ResponseWriter, r *http.Request) {
    // Set SSE headers
    w.Header().Set("Content-Type", "text/event-stream")
    w.Header().Set("Cache-Control", "no-cache")
    w.Header().Set("Connection", "keep-alive")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "streaming not supported", http.StatusInternalServerError)
        return
    }

    // Send events until client disconnects
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-r.Context().Done():
            return
        case t := <-ticker.C:
            // SSE format: "data: <content>\n\n"
            fmt.Fprintf(w, "data: %s\n\n", t.Format(time.RFC3339))
            flusher.Flush()
        }
    }
}

// SSE client (JavaScript side):
// const evtSource = new EventSource("/events");
// evtSource.onmessage = (e) => console.log(e.data);
```

---

## 🧠 QUIZ - 5 CÂU HỎI

1. Tại sao WebSocket cần HTTP Upgrade handshake?
2. Hub pattern trong WebSocket giải quyết vấn đề gì?
3. Read pump và write pump tại sao phải chạy trong goroutines riêng biệt?
4. Khi nào dùng SSE thay vì WebSocket?
5. `SetReadDeadline` và ping/pong có mục đích gì?

---

## 📌 KEY TAKEAWAYS

- WebSocket = persistent bi-directional connection qua HTTP Upgrade
- Hub pattern: centralized manager, clients giao tiếp qua channels
- Read/write pumps riêng: không block lẫn nhau
- Luôn set read limit và deadlines để tránh resource exhaustion
- SSE: simpler, server-only push, HTTP cơ bản, không cần special protocol
