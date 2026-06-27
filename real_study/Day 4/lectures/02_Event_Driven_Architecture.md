# Lesson 2: Event-Driven Architecture

## 📖 Nội dung bài học

1. Event-driven architecture là gì?
2. Pub/Sub pattern trong Go
3. Message Queue cơ bản
4. CQRS (Command Query Responsibility Segregation)
5. Outbox Pattern
6. Practical examples

---

## 1️⃣ EVENT-DRIVEN ARCHITECTURE LÀ GÌ?

**Event-Driven Architecture (EDA)** là kiến trúc trong đó các components giao tiếp qua **events** thay vì gọi trực tiếp lẫn nhau.

### So sánh kiến trúc

```
// ❌ Tight coupling - gọi trực tiếp
func (o *OrderService) PlaceOrder(order Order) {
    o.repo.Save(order)
    o.emailService.SendConfirmation(order)    // direct call
    o.inventoryService.ReduceStock(order)     // direct call
    o.analyticsService.TrackOrder(order)      // direct call
}

// ✅ Loose coupling - event-driven
func (o *OrderService) PlaceOrder(order Order) {
    o.repo.Save(order)
    o.eventBus.Publish(OrderPlacedEvent{Order: order}) // fire event
    // Email, inventory, analytics đều lắng nghe event này
}
```

### Ưu điểm

- **Decoupling**: Các services không cần biết nhau
- **Scalability**: Xử lý async, dễ scale
- **Flexibility**: Thêm subscriber mới không cần sửa publisher
- **Resilience**: Failure ở một subscriber không ảnh hưởng service khác

---

## 2️⃣ PUB/SUB PATTERN TRONG GO

### In-process Event Bus

```go
package eventbus

import (
    "fmt"
    "sync"
)

// Event là interface mọi event phải implement
type Event interface {
    EventName() string
}

// Handler xử lý một event
type Handler func(event Event) error

// EventBus quản lý publish/subscribe
type EventBus struct {
    mu       sync.RWMutex
    handlers map[string][]Handler
}

func New() *EventBus {
    return &EventBus{
        handlers: make(map[string][]Handler),
    }
}

// Subscribe đăng ký handler cho event
func (eb *EventBus) Subscribe(eventName string, handler Handler) {
    eb.mu.Lock()
    defer eb.mu.Unlock()
    eb.handlers[eventName] = append(eb.handlers[eventName], handler)
}

// Publish phát event đến tất cả subscribers
func (eb *EventBus) Publish(event Event) error {
    eb.mu.RLock()
    handlers := eb.handlers[event.EventName()]
    eb.mu.RUnlock()

    for _, h := range handlers {
        if err := h(event); err != nil {
            return fmt.Errorf("handler error for %s: %w", event.EventName(), err)
        }
    }
    return nil
}

// PublishAsync phát event async (không block)
func (eb *EventBus) PublishAsync(event Event) {
    eb.mu.RLock()
    handlers := eb.handlers[event.EventName()]
    eb.mu.RUnlock()

    for _, h := range handlers {
        h := h // capture loop variable
        go func() {
            if err := h(event); err != nil {
                fmt.Printf("async handler error: %v\n", err)
            }
        }()
    }
}
```

### Định nghĩa Events

```go
// events/order_events.go
package events

import "time"

type OrderPlaced struct {
    OrderID    string
    CustomerID string
    Amount     float64
    Items      []OrderItem
    PlacedAt   time.Time
}

func (e OrderPlaced) EventName() string { return "order.placed" }

type OrderCancelled struct {
    OrderID  string
    Reason   string
    CancelAt time.Time
}

func (e OrderCancelled) EventName() string { return "order.cancelled" }

type OrderShipped struct {
    OrderID     string
    TrackingNum string
    ShippedAt   time.Time
}

func (e OrderShipped) EventName() string { return "order.shipped" }
```

### Sử dụng Event Bus

```go
func main() {
    bus := eventbus.New()

    // Email service subscribes
    bus.Subscribe("order.placed", func(event eventbus.Event) error {
        e := event.(events.OrderPlaced)
        fmt.Printf("📧 Sending confirmation email to customer %s\n", e.CustomerID)
        return nil
    })

    // Inventory service subscribes
    bus.Subscribe("order.placed", func(event eventbus.Event) error {
        e := event.(events.OrderPlaced)
        fmt.Printf("📦 Reducing stock for order %s\n", e.OrderID)
        return nil
    })

    // Analytics subscribes
    bus.Subscribe("order.placed", func(event eventbus.Event) error {
        e := event.(events.OrderPlaced)
        fmt.Printf("📊 Tracking order %.2f for analytics\n", e.Amount)
        return nil
    })

    // Publish event
    bus.Publish(events.OrderPlaced{
        OrderID:    "ORD-001",
        CustomerID: "CUST-42",
        Amount:     150.00,
        PlacedAt:   time.Now(),
    })
}
```

---

## 3️⃣ MESSAGE QUEUE CƠ BẢN

Channel-based queue trong Go:

```go
package queue

import (
    "context"
    "sync"
)

type Message struct {
    Topic   string
    Payload []byte
}

type Queue struct {
    ch     chan Message
    wg     sync.WaitGroup
    cancel context.CancelFunc
}

func New(bufferSize int) *Queue {
    return &Queue{
        ch: make(chan Message, bufferSize),
    }
}

// Publish gửi message vào queue (non-blocking nếu có buffer)
func (q *Queue) Publish(topic string, payload []byte) {
    q.ch <- Message{Topic: topic, Payload: payload}
}

// Consume xử lý messages với worker pool
func (q *Queue) Consume(ctx context.Context, numWorkers int, handler func(Message) error) {
    for i := 0; i < numWorkers; i++ {
        q.wg.Add(1)
        go func() {
            defer q.wg.Done()
            for {
                select {
                case msg, ok := <-q.ch:
                    if !ok {
                        return
                    }
                    if err := handler(msg); err != nil {
                        // log error, could retry or DLQ
                        fmt.Printf("handler error: %v\n", err)
                    }
                case <-ctx.Done():
                    return
                }
            }
        }()
    }
}

// Close đóng queue và chờ workers xong
func (q *Queue) Close() {
    close(q.ch)
    q.wg.Wait()
}
```

---

## 4️⃣ CQRS PATTERN

**CQRS** tách biệt **Commands** (thay đổi state) và **Queries** (đọc state).

```
Traditional:
  OrderService.CreateOrder() → Read + Write same model
  OrderService.GetOrders()   → Read + Write same model

CQRS:
  Commands: CreateOrder, UpdateOrder, CancelOrder → Write model
  Queries:  GetOrder, ListOrders, SearchOrders    → Read model (có thể denormalized)
```

### Command side

```go
// commands/create_order.go
package commands

type CreateOrderCommand struct {
    CustomerID string
    Items      []OrderItem
}

type CreateOrderHandler struct {
    repo     OrderRepository
    eventBus EventBus
}

func (h *CreateOrderHandler) Handle(ctx context.Context, cmd CreateOrderCommand) (string, error) {
    // Validate
    if len(cmd.Items) == 0 {
        return "", errors.New("order must have at least one item")
    }

    // Create domain object
    order := NewOrder(cmd.CustomerID, cmd.Items)

    // Save to write store
    if err := h.repo.Save(ctx, order); err != nil {
        return "", fmt.Errorf("save order: %w", err)
    }

    // Publish event
    h.eventBus.Publish(OrderPlacedEvent{Order: order})

    return order.ID, nil
}
```

### Query side

```go
// queries/get_order.go
package queries

type GetOrderQuery struct {
    OrderID string
}

// OrderView là denormalized read model (có thể khác write model)
type OrderView struct {
    ID           string
    CustomerName string // denormalized từ customer
    Items        []ItemView
    TotalAmount  float64
    Status       string
    StatusLabel  string // "Đang xử lý", "Đã giao"
}

type GetOrderHandler struct {
    readDB ReadDB // có thể là read replica hoặc materialized view
}

func (h *GetOrderHandler) Handle(ctx context.Context, q GetOrderQuery) (*OrderView, error) {
    return h.readDB.FindOrderView(ctx, q.OrderID)
}
```

---

## 5️⃣ OUTBOX PATTERN

Giải quyết vấn đề: "Làm sao đảm bảo DB save và event publish cùng xảy ra?"

```go
// Thay vì: save → publish (không atomic)
// Dùng: save + outbox record → background poller publishes

type OutboxEvent struct {
    ID        string
    EventType string
    Payload   []byte
    CreatedAt time.Time
    Published bool
}

// Trong transaction: lưu order + outbox event
func (s *OrderService) PlaceOrder(ctx context.Context, order Order) error {
    return s.db.Transaction(func(tx *sql.Tx) error {
        // 1. Save order
        if err := saveOrderTx(tx, order); err != nil {
            return err
        }

        // 2. Save outbox event (same transaction!)
        payload, _ := json.Marshal(OrderPlacedEvent{Order: order})
        return saveOutboxEventTx(tx, OutboxEvent{
            EventType: "order.placed",
            Payload:   payload,
        })
    })
}

// Background goroutine poll và publish outbox events
func (s *OrderService) StartOutboxPoller(ctx context.Context) {
    ticker := time.NewTicker(5 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            s.processOutboxEvents(ctx)
        case <-ctx.Done():
            return
        }
    }
}
```

---

## 🧠 QUIZ - 5 CÂU HỎI

1. Sự khác nhau giữa sync và async event publishing là gì? Khi nào dùng cái nào?
2. CQRS giải quyết vấn đề gì mà traditional architecture không giải quyết được?
3. Outbox pattern giải quyết vấn đề gì?
4. Nếu một subscriber fail, event-driven architecture xử lý như thế nào?
5. Khi nào nên dùng in-process event bus vs external message broker (Kafka, RabbitMQ)?

---

## 📌 KEY TAKEAWAYS

- Event-driven → loose coupling, async processing, flexibility
- Pub/Sub pattern dễ implement với Go channels và goroutines
- CQRS tách read/write model → optimize từng side độc lập
- Outbox pattern đảm bảo at-least-once delivery
- In-process bus cho single service; external broker cho multi-service
