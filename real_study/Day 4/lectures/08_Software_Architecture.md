# Lesson 8: Software Architecture Patterns

## 📖 Nội dung bài học

1. Domain-Driven Design (DDD) cơ bản
2. Hexagonal Architecture (Ports & Adapters)
3. Clean Architecture
4. Event Sourcing
5. Áp dụng trong Go project
6. So sánh các kiến trúc

---

## 1️⃣ DOMAIN-DRIVEN DESIGN (DDD)

DDD tập trung vào **domain** (business logic) thay vì technical concerns.

### Các khái niệm cốt lõi

```
┌─────────────────────────────────────────┐
│              DOMAIN MODEL               │
│                                         │
│  Entity      Value Object    Aggregate  │
│  (có ID)     (immutable)     (root)     │
│                                         │
│  Domain Service    Domain Event         │
│  (business logic)  (điều gì đó đã xảy) │
│                                         │
│  Repository   Factory                   │
│  (interface)  (tạo aggregates)          │
└─────────────────────────────────────────┘
```

### Entity vs Value Object

```go
// Entity: có identity (ID), mutable
type Order struct {
    id         OrderID  // identity
    customerID CustomerID
    items      []OrderItem
    status     OrderStatus
    total      Money
}

func (o *Order) ID() OrderID { return o.id }

// Value Object: không có identity, immutable, comparable by value
type Money struct {
    Amount   decimal.Decimal
    Currency string
}

func (m Money) Add(other Money) (Money, error) {
    if m.Currency != other.Currency {
        return Money{}, errors.New("currency mismatch")
    }
    return Money{Amount: m.Amount.Add(other.Amount), Currency: m.Currency}, nil
}

func (m Money) Equal(other Money) bool {
    return m.Currency == other.Currency && m.Amount.Equal(other.Amount)
}

// Aggregate: nhóm các entities/value objects, 1 root
type Order struct { // Order là Aggregate Root
    id    OrderID
    items []OrderItem // entities bên trong aggregate
    // ...
}

// Chỉ tương tác với Order qua Order (root), không access OrderItem trực tiếp
func (o *Order) AddItem(product ProductID, qty int, price Money) error {
    if o.status != StatusPending {
        return ErrOrderNotEditable
    }
    // business rule: max 50 items per order
    if len(o.items) >= 50 {
        return ErrTooManyItems
    }
    o.items = append(o.items, newOrderItem(product, qty, price))
    o.recalculateTotal()
    return nil
}
```

### Domain Events

```go
// Domain events mô tả điều gì đó đã xảy ra trong domain
type OrderPlaced struct {
    OrderID    OrderID
    CustomerID CustomerID
    Total      Money
    OccurredAt time.Time
}

// Aggregate tích lũy events
type Order struct {
    // ...
    events []DomainEvent
}

func (o *Order) Place() error {
    if len(o.items) == 0 {
        return ErrEmptyOrder
    }
    o.status = StatusPlaced
    o.addEvent(OrderPlaced{
        OrderID:    o.id,
        CustomerID: o.customerID,
        Total:      o.total,
        OccurredAt: time.Now(),
    })
    return nil
}

func (o *Order) PullEvents() []DomainEvent {
    events := o.events
    o.events = nil
    return events
}
```

---

## 2️⃣ HEXAGONAL ARCHITECTURE (PORTS & ADAPTERS)

```
                    ┌─────────────────────────────┐
                    │       APPLICATION CORE       │
                    │                              │
  HTTP Adapter  →   │  ┌──────────────────────┐   │   → DB Adapter
  gRPC Adapter  →   │  │   Domain / Use Cases  │   │   → Cache Adapter
  CLI Adapter   →   │  └──────────────────────┘   │   → Message Broker
                    │                              │
                    │  Inbound Ports  Outbound Ports│
                    └─────────────────────────────┘
```

### Project structure

```
order-service/
├── domain/                    # Không phụ thuộc gì bên ngoài
│   ├── order.go               # Aggregate, entities
│   ├── order_events.go        # Domain events
│   ├── order_repository.go    # Repository interface (outbound port)
│   └── order_service.go       # Domain service
│
├── application/               # Use cases (orchestrates domain)
│   ├── commands/
│   │   ├── place_order.go
│   │   └── cancel_order.go
│   └── queries/
│       └── get_order.go
│
├── adapters/
│   ├── inbound/               # HTTP, gRPC, CLI (inbound adapters)
│   │   ├── http/
│   │   │   └── order_handler.go
│   │   └── grpc/
│   │       └── order_server.go
│   └── outbound/              # DB, Cache, external APIs (outbound adapters)
│       ├── postgres/
│       │   └── order_repo.go
│       └── redis/
│           └── order_cache.go
│
└── main.go                    # Wiring everything together
```

### Dependency injection ở main.go

```go
func main() {
    // Infrastructure
    db := setupDB()
    cache := setupRedis()
    eventBus := eventbus.New()

    // Outbound adapters (implement domain interfaces)
    orderRepo := postgres.NewOrderRepository(db)
    orderCache := redis.NewOrderCache(cache)

    // Application layer (use cases)
    placeOrderHandler := commands.NewPlaceOrderHandler(orderRepo, eventBus)
    getOrderHandler := queries.NewGetOrderHandler(orderRepo, orderCache)

    // Inbound adapters
    httpHandler := httpAdapter.NewOrderHandler(placeOrderHandler, getOrderHandler)

    // Compose
    mux := http.NewServeMux()
    mux.Handle("/api/orders", httpHandler)

    log.Fatal(http.ListenAndServe(":8080", mux))
}
```

---

## 3️⃣ CLEAN ARCHITECTURE

Giống hexagonal nhưng phân chia thành concentric circles:

```
    ┌────────────────────────────────┐
    │         Frameworks & DB        │  ← outermost
    │   ┌────────────────────────┐   │
    │   │     Interface Adapters  │   │
    │   │   ┌────────────────┐   │   │
    │   │   │  Application   │   │   │
    │   │   │  ┌──────────┐  │   │   │
    │   │   │  │ Entities │  │   │   │  ← innermost (no deps)
    │   │   │  └──────────┘  │   │   │
    │   │   └────────────────┘   │   │
    │   └────────────────────────┘   │
    └────────────────────────────────┘
```

**Dependency Rule**: Dependencies chỉ trỏ vào trong (toward entities), không bao giờ ra ngoài.

---

## 4️⃣ EVENT SOURCING

Thay vì lưu current state, lưu **chuỗi events** đã xảy ra.

```go
// Mỗi event là một "fact" đã xảy ra, không thể thay đổi
type EventStore interface {
    Append(ctx context.Context, streamID string, events []Event, expectedVersion int) error
    Load(ctx context.Context, streamID string) ([]Event, error)
}

// Rebuild state từ events
func RebuildOrder(events []Event) (*Order, error) {
    order := &Order{}
    for _, event := range events {
        if err := order.Apply(event); err != nil {
            return nil, err
        }
    }
    return order, nil
}

func (o *Order) Apply(event Event) error {
    switch e := event.(type) {
    case OrderCreated:
        o.id = e.OrderID
        o.customerID = e.CustomerID
        o.status = StatusPending
    case ItemAdded:
        o.items = append(o.items, OrderItem{
            ProductID: e.ProductID,
            Qty:       e.Qty,
            Price:     e.Price,
        })
        o.total = o.total.MustAdd(e.Price.Multiply(e.Qty))
    case OrderPlaced:
        o.status = StatusPlaced
    case OrderCancelled:
        o.status = StatusCancelled
    default:
        return fmt.Errorf("unknown event type: %T", event)
    }
    return nil
}

// Lợi ích của Event Sourcing:
// ✅ Full audit trail - biết mọi thay đổi
// ✅ Time travel - rebuild state tại bất kỳ thời điểm nào
// ✅ Decoupled projections - tạo nhiều read models từ cùng events
// ✅ Event replay - debug, testing, migration
// ❌ Complexity cao hơn
// ❌ Query phức tạp hơn (cần CQRS)
```

---

## 5️⃣ SO SÁNH CÁC KIẾN TRÚC

| Tiêu chí       | Layered    | Hexagonal  | Clean      | Event Sourcing |
| -------------- | ---------- | ---------- | ---------- | -------------- |
| Complexity     | Thấp       | Trung bình | Cao        | Rất cao        |
| Testability    | Trung bình | Cao        | Rất cao    | Cao            |
| Flexibility    | Thấp       | Cao        | Rất cao    | Cao            |
| Learning curve | Thấp       | Trung bình | Cao        | Rất cao        |
| Phù hợp với    | CRUD apps  | Services   | Enterprise | Audit-heavy    |

### Khi nào dùng gì?

- **Layered**: CRUD apps đơn giản, team nhỏ, deadline gấp
- **Hexagonal**: Microservices, cần swap infrastructure dễ
- **Clean**: Enterprise apps, domain phức tạp, team lớn
- **Event Sourcing**: Audit trail cần thiết, cần time-travel debugging

---

## 🧠 QUIZ - FINAL 10 CÂU HỎI

1. Entity và Value Object khác nhau ở điểm gì cơ bản nhất?
2. Aggregate Root có vai trò gì trong DDD?
3. Dependency Rule trong Clean Architecture là gì?
4. Tại sao Repository trong domain layer chỉ là interface?
5. Event Sourcing khác CRUD ở điểm cốt lõi nào?
6. Hexagonal Architecture gọi là "Ports & Adapters" - Port là gì, Adapter là gì?
7. Domain Event và Integration Event khác nhau thế nào?
8. Trong CQRS, read model và write model có thể dùng khác DB không?
9. Snapshot trong Event Sourcing dùng để làm gì?
10. Khi nào không nên dùng DDD (over-engineering)?

---

## 📌 KEY TAKEAWAYS

- DDD: focus on domain, model business concepts explicitly
- Entity = identity + mutable; Value Object = no identity + immutable
- Hexagonal: core domain không phụ thuộc framework hay DB
- Dependency Rule: code outer biết code inner, không ngược lại
- Event Sourcing: events là source of truth, không phải current state
- Chọn kiến trúc phù hợp với domain complexity, đừng over-engineer
