# Lesson 3: Advanced Database Patterns

## 📖 Nội dung bài học

1. Repository Pattern
2. Database Migrations
3. Transactions & ACID
4. Saga Pattern (distributed transactions)
5. Connection Pooling & Performance
6. Practical examples với `database/sql`

---

## 1️⃣ REPOSITORY PATTERN

Repository pattern tách logic truy cập dữ liệu ra khỏi business logic.

```go
// Domain interface (không phụ thuộc DB cụ thể)
type OrderRepository interface {
    FindByID(ctx context.Context, id string) (*Order, error)
    FindByCustomer(ctx context.Context, customerID string) ([]*Order, error)
    Save(ctx context.Context, order *Order) error
    Update(ctx context.Context, order *Order) error
    Delete(ctx context.Context, id string) error
}

// PostgreSQL implementation
type postgresOrderRepo struct {
    db *sql.DB
}

func NewPostgresOrderRepo(db *sql.DB) OrderRepository {
    return &postgresOrderRepo{db: db}
}

func (r *postgresOrderRepo) FindByID(ctx context.Context, id string) (*Order, error) {
    row := r.db.QueryRowContext(ctx,
        `SELECT id, customer_id, status, total, created_at FROM orders WHERE id = $1`,
        id,
    )

    var o Order
    err := row.Scan(&o.ID, &o.CustomerID, &o.Status, &o.Total, &o.CreatedAt)
    if err == sql.ErrNoRows {
        return nil, ErrNotFound
    }
    if err != nil {
        return nil, fmt.Errorf("query order: %w", err)
    }
    return &o, nil
}

func (r *postgresOrderRepo) Save(ctx context.Context, order *Order) error {
    _, err := r.db.ExecContext(ctx,
        `INSERT INTO orders (id, customer_id, status, total, created_at)
         VALUES ($1, $2, $3, $4, $5)`,
        order.ID, order.CustomerID, order.Status, order.Total, order.CreatedAt,
    )
    if err != nil {
        return fmt.Errorf("insert order: %w", err)
    }
    return nil
}

// In-memory implementation (dùng cho testing)
type inMemoryOrderRepo struct {
    mu     sync.RWMutex
    orders map[string]*Order
}

func NewInMemoryOrderRepo() OrderRepository {
    return &inMemoryOrderRepo{orders: make(map[string]*Order)}
}

func (r *inMemoryOrderRepo) FindByID(_ context.Context, id string) (*Order, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    o, ok := r.orders[id]
    if !ok {
        return nil, ErrNotFound
    }
    return o, nil
}
```

---

## 2️⃣ DATABASE MIGRATIONS

### Migration files pattern

```
migrations/
    001_create_orders.up.sql
    001_create_orders.down.sql
    002_add_order_items.up.sql
    002_add_order_items.down.sql
    003_add_order_index.up.sql
    003_add_order_index.down.sql
```

### SQL migration files

```sql
-- 001_create_orders.up.sql
CREATE TABLE IF NOT EXISTS orders (
    id          VARCHAR(36)    PRIMARY KEY,
    customer_id VARCHAR(36)    NOT NULL,
    status      VARCHAR(20)    NOT NULL DEFAULT 'pending',
    total       DECIMAL(10,2)  NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ    NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ    NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_orders_customer_id ON orders(customer_id);
CREATE INDEX idx_orders_status      ON orders(status);

-- 001_create_orders.down.sql
DROP TABLE IF EXISTS orders;
```

### Simple migration runner trong Go

```go
package migration

import (
    "database/sql"
    "fmt"
    "os"
    "path/filepath"
    "sort"
    "strings"
)

type Migrator struct {
    db         *sql.DB
    migrateDir string
}

func New(db *sql.DB, dir string) *Migrator {
    return &Migrator{db: db, migrateDir: dir}
}

// EnsureMigrationsTable tạo bảng tracking nếu chưa có
func (m *Migrator) EnsureMigrationsTable() error {
    _, err := m.db.Exec(`
        CREATE TABLE IF NOT EXISTS schema_migrations (
            version    VARCHAR(255) PRIMARY KEY,
            applied_at TIMESTAMPTZ  NOT NULL DEFAULT NOW()
        )
    `)
    return err
}

// Up chạy tất cả migrations chưa applied
func (m *Migrator) Up() error {
    if err := m.EnsureMigrationsTable(); err != nil {
        return err
    }

    files, _ := filepath.Glob(filepath.Join(m.migrateDir, "*.up.sql"))
    sort.Strings(files)

    for _, file := range files {
        version := strings.TrimSuffix(filepath.Base(file), ".up.sql")

        // Check if already applied
        var count int
        m.db.QueryRow(`SELECT COUNT(*) FROM schema_migrations WHERE version = $1`, version).Scan(&count)
        if count > 0 {
            continue
        }

        // Run migration
        content, err := os.ReadFile(file)
        if err != nil {
            return fmt.Errorf("read %s: %w", file, err)
        }

        if _, err := m.db.Exec(string(content)); err != nil {
            return fmt.Errorf("apply %s: %w", version, err)
        }

        m.db.Exec(`INSERT INTO schema_migrations (version) VALUES ($1)`, version)
        fmt.Printf("✅ Applied migration: %s\n", version)
    }
    return nil
}
```

---

## 3️⃣ TRANSACTIONS & ACID

### Basic transaction

```go
func (r *postgresOrderRepo) CreateOrderWithItems(ctx context.Context, order *Order, items []OrderItem) error {
    tx, err := r.db.BeginTx(ctx, &sql.TxOptions{
        Isolation: sql.LevelSerializable, // hoặc LevelReadCommitted
    })
    if err != nil {
        return fmt.Errorf("begin tx: %w", err)
    }
    // Luôn rollback nếu không commit
    defer tx.Rollback()

    // Insert order
    if _, err := tx.ExecContext(ctx,
        `INSERT INTO orders (id, customer_id, status, total) VALUES ($1, $2, $3, $4)`,
        order.ID, order.CustomerID, order.Status, order.Total,
    ); err != nil {
        return fmt.Errorf("insert order: %w", err)
    }

    // Insert items
    for _, item := range items {
        if _, err := tx.ExecContext(ctx,
            `INSERT INTO order_items (order_id, product_id, qty, price) VALUES ($1, $2, $3, $4)`,
            order.ID, item.ProductID, item.Qty, item.Price,
        ); err != nil {
            return fmt.Errorf("insert item: %w", err)
        }
    }

    // Update inventory
    for _, item := range items {
        result, err := tx.ExecContext(ctx,
            `UPDATE inventory SET qty = qty - $1 WHERE product_id = $2 AND qty >= $1`,
            item.Qty, item.ProductID,
        )
        if err != nil {
            return fmt.Errorf("update inventory: %w", err)
        }
        rows, _ := result.RowsAffected()
        if rows == 0 {
            return fmt.Errorf("insufficient stock for product %s", item.ProductID)
        }
    }

    return tx.Commit()
}
```

### Transaction helper

```go
// WithTx chạy fn trong một transaction, tự rollback nếu lỗi
func WithTx(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    if err := fn(tx); err != nil {
        return err
    }
    return tx.Commit()
}

// Usage
err := WithTx(ctx, db, func(tx *sql.Tx) error {
    // tất cả operations trong cùng transaction
    return nil
})
```

---

## 4️⃣ SAGA PATTERN

Dùng khi cần transaction trải rộng nhiều services (distributed transaction).

### Choreography-based Saga

```go
// Mỗi service lắng nghe event và phát event tiếp theo
//
// 1. OrderService: OrderCreated →
// 2. InventoryService: StockReserved → (hoặc StockFailed)
// 3. PaymentService: PaymentCharged → (hoặc PaymentFailed)
// 4. ShippingService: ShipmentCreated →
// 5. OrderService: OrderCompleted

type SagaStep struct {
    Execute     func(ctx context.Context) error
    Compensate  func(ctx context.Context) error // undo nếu step sau fail
}

type Saga struct {
    steps     []SagaStep
    completed []int // index của các steps đã chạy
}

func (s *Saga) Execute(ctx context.Context) error {
    for i, step := range s.steps {
        if err := step.Execute(ctx); err != nil {
            // Compensate các steps đã chạy (theo thứ tự ngược)
            s.compensate(ctx)
            return fmt.Errorf("saga step %d failed: %w", i, err)
        }
        s.completed = append(s.completed, i)
    }
    return nil
}

func (s *Saga) compensate(ctx context.Context) {
    for i := len(s.completed) - 1; i >= 0; i-- {
        idx := s.completed[i]
        if err := s.steps[idx].Compensate(ctx); err != nil {
            fmt.Printf("compensation failed for step %d: %v\n", idx, err)
        }
    }
}
```

---

## 5️⃣ CONNECTION POOLING

```go
func SetupDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    // Connection pool configuration
    db.SetMaxOpenConns(25)                 // max concurrent connections
    db.SetMaxIdleConns(10)                 // max idle connections in pool
    db.SetConnMaxLifetime(5 * time.Minute) // max time a connection is reused
    db.SetConnMaxIdleTime(1 * time.Minute) // max time a connection sits idle

    // Verify connection
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := db.PingContext(ctx); err != nil {
        return nil, fmt.Errorf("ping db: %w", err)
    }

    return db, nil
}
```

---

## 🧠 QUIZ - 5 CÂU HỎI

1. Repository pattern mang lại lợi ích gì so với gọi DB trực tiếp?
2. Tại sao `defer tx.Rollback()` an toàn ngay cả sau khi `tx.Commit()`?
3. Sự khác biệt giữa Isolation Level `ReadCommitted` và `Serializable`?
4. Saga choreography vs orchestration: khi nào dùng cái nào?
5. `SetMaxOpenConns` và `SetMaxIdleConns` khác nhau như thế nào?

---

## 📌 KEY TAKEAWAYS

- Repository pattern → business logic không biết về DB cụ thể
- Migration files → quản lý schema changes có kiểm soát
- Luôn dùng `defer tx.Rollback()` + explicit `Commit()` pattern
- Saga thay thế 2PC cho distributed transactions
- Tune connection pool phù hợp với load của service
