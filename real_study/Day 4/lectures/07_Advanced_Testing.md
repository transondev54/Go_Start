# Lesson 7: Advanced Testing

## 📖 Nội dung bài học

1. Testing pyramid
2. Table-driven tests
3. Mock generation với `mockgen`
4. Interface-based mocking thủ công
5. Integration tests
6. Fuzzing (Go 1.18+)
7. Test helpers & utilities

---

## 1️⃣ TESTING PYRAMID

```
         /\
        /  \      E2E Tests (ít, chậm, đắt)
       / E2E\
      /──────\
     /        \   Integration Tests (vừa phải)
    / Integration\
   /──────────────\
  /                \
 /   Unit Tests     \  (nhiều, nhanh, rẻ)
/────────────────────\
```

- **Unit tests**: Test từng function/method độc lập, dùng mocks
- **Integration tests**: Test kết hợp nhiều components (DB, HTTP server)
- **E2E tests**: Test toàn bộ flow từ client đến DB

---

## 2️⃣ TABLE-DRIVEN TESTS

```go
func TestCalculateDiscount(t *testing.T) {
    tests := []struct {
        name     string
        amount   float64
        code     string
        want     float64
        wantErr  bool
    }{
        {
            name:    "10% discount code",
            amount:  100.0,
            code:    "SAVE10",
            want:    90.0,
            wantErr: false,
        },
        {
            name:    "50% discount code",
            amount:  200.0,
            code:    "HALF",
            want:    100.0,
            wantErr: false,
        },
        {
            name:    "invalid code",
            amount:  100.0,
            code:    "INVALID",
            want:    0,
            wantErr: true,
        },
        {
            name:    "zero amount",
            amount:  0,
            code:    "SAVE10",
            want:    0,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := CalculateDiscount(tt.amount, tt.code)

            if (err != nil) != tt.wantErr {
                t.Errorf("CalculateDiscount() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("CalculateDiscount() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

---

## 3️⃣ INTERFACE-BASED MOCKING

Viết mock thủ công không cần tool:

```go
// Interface cần mock
type OrderRepository interface {
    FindByID(ctx context.Context, id string) (*Order, error)
    Save(ctx context.Context, order *Order) error
}

// Mock implementation
type MockOrderRepository struct {
    FindByIDFunc func(ctx context.Context, id string) (*Order, error)
    SaveFunc     func(ctx context.Context, order *Order) error

    // Track calls
    FindByIDCalls []string
    SaveCalls     []*Order
}

func (m *MockOrderRepository) FindByID(ctx context.Context, id string) (*Order, error) {
    m.FindByIDCalls = append(m.FindByIDCalls, id)
    if m.FindByIDFunc != nil {
        return m.FindByIDFunc(ctx, id)
    }
    return nil, nil
}

func (m *MockOrderRepository) Save(ctx context.Context, order *Order) error {
    m.SaveCalls = append(m.SaveCalls, order)
    if m.SaveFunc != nil {
        return m.SaveFunc(ctx, order)
    }
    return nil
}

// Sử dụng trong test
func TestOrderService_PlaceOrder(t *testing.T) {
    mockRepo := &MockOrderRepository{
        SaveFunc: func(ctx context.Context, order *Order) error {
            return nil // happy path
        },
    }

    svc := NewOrderService(mockRepo)
    _, err := svc.PlaceOrder(context.Background(), PlaceOrderInput{
        CustomerID: "cust-1",
        Items:      []OrderItem{{ProductID: "p1", Qty: 2}},
    })

    if err != nil {
        t.Fatalf("PlaceOrder() error: %v", err)
    }

    // Assert mock was called
    if len(mockRepo.SaveCalls) != 1 {
        t.Errorf("expected Save to be called once, got %d", len(mockRepo.SaveCalls))
    }
}
```

### Mockgen (tự động generate)

```go
//go:generate mockgen -source=repository.go -destination=mocks/mock_repository.go

// Sau khi generate:
import "myapp/mocks"
import "github.com/golang/mock/gomock"

func TestWithMockgen(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := mocks.NewMockOrderRepository(ctrl)

    // Set expectations
    mockRepo.EXPECT().
        FindByID(gomock.Any(), "order-1").
        Return(&Order{ID: "order-1"}, nil).
        Times(1)

    svc := NewOrderService(mockRepo)
    order, err := svc.GetOrder(context.Background(), "order-1")

    if err != nil || order.ID != "order-1" {
        t.Fail()
    }
}
```

---

## 4️⃣ INTEGRATION TESTS

```go
//go:build integration

package integration_test

import (
    "database/sql"
    "os"
    "testing"

    _ "github.com/lib/pq"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
    // Setup: connect to test database
    dsn := os.Getenv("TEST_DATABASE_URL")
    if dsn == "" {
        dsn = "postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable"
    }

    var err error
    testDB, err = sql.Open("postgres", dsn)
    if err != nil {
        panic("cannot connect to test db: " + err.Error())
    }

    // Run migrations
    if err := runMigrations(testDB); err != nil {
        panic("migrations failed: " + err.Error())
    }

    // Run tests
    code := m.Run()

    // Teardown
    testDB.Close()
    os.Exit(code)
}

func TestOrderRepository_Save_Integration(t *testing.T) {
    repo := NewPostgresOrderRepo(testDB)

    order := &Order{
        ID:         "test-order-1",
        CustomerID: "cust-1",
        Status:     "pending",
        Total:      150.00,
        CreatedAt:  time.Now(),
    }

    // Test Save
    if err := repo.Save(context.Background(), order); err != nil {
        t.Fatalf("Save() error: %v", err)
    }

    // Test FindByID
    found, err := repo.FindByID(context.Background(), order.ID)
    if err != nil {
        t.Fatalf("FindByID() error: %v", err)
    }
    if found.ID != order.ID {
        t.Errorf("FindByID() ID = %s, want %s", found.ID, order.ID)
    }

    // Cleanup
    testDB.Exec(`DELETE FROM orders WHERE id = $1`, order.ID)
}

// Chạy integration tests:
// go test -tags=integration ./...
```

---

## 5️⃣ FUZZING (Go 1.18+)

```go
// Fuzzing tự động generate inputs để tìm crashes
func FuzzParseOrderID(f *testing.F) {
    // Seed corpus - valid examples
    f.Add("ORD-001")
    f.Add("ORD-abc-123")
    f.Add("")

    f.Fuzz(func(t *testing.T, input string) {
        // Function không được panic với bất kỳ input nào
        result, err := ParseOrderID(input)
        if err != nil {
            return // error là expected cho invalid input
        }
        // Nếu không có error, result phải valid
        if result.String() == "" {
            t.Errorf("ParseOrderID(%q) returned empty string without error", input)
        }
    })
}

// Chạy fuzzing:
// go test -fuzz=FuzzParseOrderID -fuzztime=30s
```

---

## 6️⃣ TEST HELPERS & UTILITIES

```go
// testutil/testutil.go - shared test utilities

package testutil

import (
    "testing"
    "time"
)

// AssertError kiểm tra error thoả mãn điều kiện
func AssertError(t *testing.T, err error, msg string) {
    t.Helper()
    if err == nil {
        t.Errorf("expected error containing %q, got nil", msg)
        return
    }
    if !strings.Contains(err.Error(), msg) {
        t.Errorf("expected error containing %q, got %q", msg, err.Error())
    }
}

// Eventually retry cho đến khi condition true hoặc timeout
func Eventually(t *testing.T, condition func() bool, timeout, interval time.Duration, msg string) {
    t.Helper()
    deadline := time.Now().Add(timeout)
    for time.Now().Before(deadline) {
        if condition() {
            return
        }
        time.Sleep(interval)
    }
    t.Fatalf("condition not met within %s: %s", timeout, msg)
}

// NewTestServer tạo httptest server với cleanup
func NewTestServer(t *testing.T, handler http.Handler) *httptest.Server {
    t.Helper()
    srv := httptest.NewServer(handler)
    t.Cleanup(srv.Close)
    return srv
}
```

---

## 🧠 QUIZ - 5 CÂU HỎI

1. Tại sao unit tests nên chiếm phần lớn trong testing pyramid?
2. Lợi ích của table-driven tests so với viết test case riêng lẻ?
3. Khi nào dùng mock thủ công vs mockgen?
4. Build tag `//go:build integration` dùng để làm gì?
5. Fuzzing giúp tìm loại bugs nào mà unit tests khó tìm được?

---

## 📌 KEY TAKEAWAYS

- Testing pyramid: nhiều unit, ít integration, rất ít E2E
- Table-driven tests → dễ thêm test case, dễ đọc
- Mock qua interface → tests không phụ thuộc implementation
- Integration tests với build tags → chạy riêng, có thể skip trong CI nhanh
- Fuzzing → tìm crashes/panics với random inputs, đặc biệt tốt cho parsers
