# Project 3: Event-Driven Order System

## 🎯 Mục tiêu

Xây dựng một **order management system** theo kiến trúc event-driven với:

- ✅ Domain-Driven Design: Aggregates, Value Objects, Domain Events
- ✅ CQRS: tách Command và Query handlers
- ✅ In-process Event Bus (pub/sub)
- ✅ Saga pattern (đặt hàng → giảm kho → thanh toán)
- ✅ Event Sourcing cho Order aggregate
- ✅ Observability: structured logs + Prometheus metrics
- ✅ RESTful HTTP API
- ✅ Comprehensive test suite

---

## 📋 Yêu cầu

### Domain Model

```go
// Aggregates
type Order struct {
    id         OrderID
    customerID CustomerID
    items      []OrderItem
    status     OrderStatus
    total      Money
    events     []DomainEvent
}

type Inventory struct {
    productID ProductID
    qty       int
    reserved  int
}

// Value Objects (immutable)
type OrderID    string
type CustomerID string
type ProductID  string
type Money      struct { Amount float64; Currency string }

// Domain Events
type OrderCreated    struct { OrderID, CustomerID, Total, OccurredAt }
type ItemAdded       struct { OrderID, ProductID, Qty, Price, OccurredAt }
type OrderPlaced     struct { OrderID, Total, OccurredAt }
type StockReserved   struct { OrderID, ProductID, Qty, OccurredAt }
type StockFailed     struct { OrderID, ProductID, Reason, OccurredAt }
type PaymentCharged  struct { OrderID, Amount, OccurredAt }
type PaymentFailed   struct { OrderID, Reason, OccurredAt }
type OrderCompleted  struct { OrderID, OccurredAt }
type OrderCancelled  struct { OrderID, Reason, OccurredAt }
```

### Commands & Queries (CQRS)

```go
// Commands (change state)
type CreateOrderCommand struct { CustomerID string; Items []ItemInput }
type AddItemCommand     struct { OrderID, ProductID string; Qty int }
type PlaceOrderCommand  struct { OrderID string }
type CancelOrderCommand struct { OrderID, Reason string }

// Queries (read state - no side effects)
type GetOrderQuery    struct { OrderID string }
type ListOrdersQuery  struct { CustomerID string; Status string }

// Read models (can be denormalized)
type OrderView struct {
    ID           string
    CustomerID   string
    Items        []ItemView
    TotalAmount  float64
    Status       string
    StatusLabel  string  // "Pending", "Processing", "Completed"
    CanCancel    bool
    CreatedAt    time.Time
}
```

### HTTP API Endpoints

| Method | Path                      | Description           |
| ------ | ------------------------- | --------------------- |
| POST   | `/orders`                 | Create order          |
| GET    | `/orders/:id`             | Get order by ID       |
| GET    | `/orders?customer_id=xxx` | List customer orders  |
| POST   | `/orders/:id/items`       | Add item to order     |
| POST   | `/orders/:id/place`       | Place (confirm) order |
| POST   | `/orders/:id/cancel`      | Cancel order          |
| GET    | `/metrics`                | Prometheus metrics    |
| GET    | `/health`                 | Health check          |

---

## 🏗️ Architecture

```
event_order_system/
├── domain/
│   ├── order/
│   │   ├── order.go           # Order aggregate
│   │   ├── order_events.go    # Domain events
│   │   └── order_repo.go      # Repository interface
│   ├── inventory/
│   │   ├── inventory.go       # Inventory aggregate
│   │   └── inventory_repo.go
│   └── shared/
│       └── value_objects.go   # Money, OrderID, etc.
│
├── application/
│   ├── commands/
│   │   ├── create_order.go
│   │   ├── place_order.go
│   │   └── cancel_order.go
│   ├── queries/
│   │   ├── get_order.go
│   │   └── list_orders.go
│   └── sagas/
│       └── place_order_saga.go
│
├── infrastructure/
│   ├── eventbus/
│   │   └── eventbus.go        # In-process pub/sub
│   ├── repository/
│   │   ├── memory_order_repo.go
│   │   └── memory_inventory_repo.go
│   └── metrics/
│       └── metrics.go
│
├── adapters/
│   └── http/
│       ├── order_handler.go
│       └── middleware.go
│
├── go.mod
└── main.go
```

---

## 🚀 Implementation Steps

### Step 1: Domain Layer

Bắt đầu từ `domain/order/order.go`:

```go
// Không import gì từ infrastructure!
type Order struct { ... }

func NewOrder(customerID CustomerID) *Order { ... }
func (o *Order) AddItem(product ProductID, qty int, price Money) error { ... }
func (o *Order) Place() error { ... }
func (o *Order) Cancel(reason string) error { ... }
func (o *Order) PullEvents() []DomainEvent { ... }
```

### Step 2: Event Bus

```go
// infrastructure/eventbus/eventbus.go
type EventBus struct { ... }
func (eb *EventBus) Subscribe(eventName string, handler Handler) { ... }
func (eb *EventBus) Publish(event DomainEvent) { ... }
```

### Step 3: Place Order Saga

```go
// application/sagas/place_order_saga.go
// Subscribe to events và orchestrate:
// OrderPlaced → reserveStock → StockReserved → chargePayment → PaymentCharged → completeOrder
// StockFailed / PaymentFailed → cancelOrder (compensation)
```

### Step 4: Command/Query Handlers

```go
// Mỗi handler nhận repository và eventbus
// Command handlers: modify state, publish events
// Query handlers: read-only, no side effects
```

### Step 5: HTTP Layer

```go
// adapters/http/order_handler.go
// Decode JSON → call command handler → return response
// Implement tất cả 6 endpoints
```

### Step 6: Metrics

```go
var (
    ordersCreated   = prometheus.NewCounter(...)
    ordersCompleted = prometheus.NewCounter(...)
    ordersCancelled = prometheus.NewCounter(...)
    orderDuration   = prometheus.NewHistogram(...)
)
```

### Step 7: Tests

```go
// domain/order/order_test.go - unit tests
// application/commands/create_order_test.go - with mock repos
// Test saga happy path and compensation
```

---

## 📊 Expected Flow

```
POST /orders {"customer_id": "cust-1", "items": [...]}
→ CreateOrderCommand
→ Order created, OrderCreated event published
← 201 {"order_id": "ord-123"}

POST /orders/ord-123/place
→ PlaceOrderCommand
→ OrderPlaced event published
→ Saga: reserveStock → StockReserved → chargePayment → PaymentCharged → OrderCompleted
← 200 {"status": "completed"}

# Nếu stock không đủ:
→ StockFailed event
→ Saga compensation: cancelOrder
← GET /orders/ord-123 → status: "cancelled"
```

---

## 🌟 Bonus Features

- [ ] Event Sourcing: lưu events vào append-only log, rebuild từ events
- [ ] Snapshotting: lưu snapshot của aggregate để tránh replay quá nhiều events
- [ ] Dead Letter Queue: events không xử lý được → DLQ để retry sau
- [ ] Distributed tracing: trace ID qua toàn bộ saga
- [ ] Outbox pattern: đảm bảo DB save và event publish atomic
- [ ] Read model projection: cập nhật read model khi nhận events

---

## ✅ Tiêu chí hoàn thành

- [ ] Domain layer hoàn toàn không import infrastructure
- [ ] Happy path: tạo order → đặt hàng → hoàn thành
- [ ] Compensation: stock fail → order cancelled
- [ ] Tất cả 6 HTTP endpoints hoạt động đúng
- [ ] Prometheus metrics expose tại `/metrics`
- [ ] `go test ./...` pass (bao gồm saga tests)
- [ ] `go test -race ./...` không có data races
- [ ] Structured logs hiển thị event flow rõ ràng
