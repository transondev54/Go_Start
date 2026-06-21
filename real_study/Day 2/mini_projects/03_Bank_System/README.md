# Mini Project 3: Bank System with Interfaces & Database

## 📝 Mô tả

Xây dựng một bank system sử dụng:

- **Interfaces** để define account types
- **Database (SQLite)** để persist data
- **Goroutines** để xử lý concurrent transactions
- **Error handling** cho edge cases

---

## 📋 Yêu cầu

### Interfaces & Types

1. **Account Interface**:

   ```go
   type Account interface {
       GetBalance() float64
       Deposit(amount float64) error
       Withdraw(amount float64) error
       GetType() string
       GetInfo() string
   }
   ```

2. **Account Types** (implement interface):
   - `SavingsAccount` - thấp lãi suất, rút tiền bị phí
   - `CheckingAccount` - không lãi, rút tự do
   - `InvestmentAccount` - rút sớm bị phạt

3. **Database Schema**:

   ```sql
   CREATE TABLE accounts (
       id INTEGER PRIMARY KEY,
       account_number TEXT UNIQUE,
       account_type TEXT,
       balance REAL,
       created_at TIMESTAMP
   );

   CREATE TABLE transactions (
       id INTEGER PRIMARY KEY,
       account_id INTEGER,
       type TEXT,
       amount REAL,
       timestamp TIMESTAMP
   );
   ```

### Tính năng bắt buộc

1. **Account Operations**:
   - Create account (Savings/Checking/Investment)
   - Deposit money
   - Withdraw money
   - Check balance
   - Transfer between accounts

2. **Validation**:
   - Không withdraw quá balance
   - Account type specific rules
   - Transaction logging

3. **Concurrent Safety**:
   - Handle simultaneous transactions
   - No race conditions
   - Use sync.Mutex hoặc channels

4. **Persistence**:
   - Lưu tất cả transactions
   - Load accounts từ DB
   - Transaction history

### Ví dụ output

```
╔═══════════════════════════════╗
║    Bank System v1.0           ║
╚═══════════════════════════════╝

1. Create Account
2. Deposit
3. Withdraw
4. Transfer
5. View Balance
6. Transaction History
7. Exit

Enter choice: 1
Select account type:
  1. Savings
  2. Checking
  3. Investment
Your choice: 1

✅ Savings Account created!
Account Number: SAV-001
Initial Balance: $0.00

---

Enter choice: 2
Enter account number: SAV-001
Enter deposit amount: $1000
✅ Deposit successful!
New balance: $1000.00

Enter choice: 3
Enter account number: SAV-001
Enter withdrawal amount: $100
⚠️  Savings account withdrawal fee: $5.00
✅ Withdrawn $100.00
New balance: $894.95

Enter choice: 6
Transaction History for SAV-001:
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
[1] DEPOSIT    +$1000.00  2024-01-15 10:30:00
[2] WITHDRAW   -$105.00   2024-01-15 10:35:00
```

---

## 🎯 Learning Objectives

- ✅ Interfaces & polymorphism
- ✅ Database design & SQL
- ✅ Goroutines & concurrent transactions
- ✅ Sync primitives (Mutex, WaitGroup)
- ✅ Error handling & validation
- ✅ Code organization

---

## 📚 Bước thực hiện

### Bước 1: Setup

```bash
mkdir bank_system
cd bank_system
go mod init bank_system
go get github.com/mattn/go-sqlite3
code main.go
```

### Bước 2: Define interfaces

```go
type Account interface {
    GetBalance() float64
    Deposit(amount float64) error
    Withdraw(amount float64) error
    GetAccountNumber() string
    GetType() string
}

type Bank interface {
    CreateAccount(accountType string) (Account, error)
    Transfer(from, to Account, amount float64) error
    GetTransaction(accountID int) []Transaction
}
```

### Bước 3: Implement types

```go
type SavingsAccount struct {
    ID             int
    AccountNumber  string
    Balance        float64
    WithdrawalFee  float64
    mu             sync.Mutex
}

func (sa *SavingsAccount) Withdraw(amount float64) error {
    sa.mu.Lock()
    defer sa.mu.Unlock()

    totalAmount := amount + sa.WithdrawalFee
    if totalAmount > sa.Balance {
        return fmt.Errorf("insufficient balance")
    }

    sa.Balance -= totalAmount
    return nil
}
```

### Bước 4: Database layer

```go
func CreateDatabase() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "bank.db")
    // ... create schema
    return db, nil
}

func SaveTransaction(db *sql.DB, tx *Transaction) error {
    // Insert transaction
}

func GetTransactions(db *sql.DB, accountID int) []Transaction {
    // Query transactions
}
```

### Bước 5: Bank implementation

```go
type SimpleBank struct {
    accounts map[string]Account
    mu       sync.Mutex
    db       *sql.DB
}

func (b *SimpleBank) Transfer(from, to Account, amount float64) error {
    b.mu.Lock()
    defer b.mu.Unlock()

    if err := from.Withdraw(amount); err != nil {
        return err
    }
    if err := to.Deposit(amount); err != nil {
        from.Deposit(amount) // Rollback
        return err
    }

    return b.saveTransactions(from, to, amount)
}
```

### Bước 6: Concurrent transactions test

```go
func (b *SimpleBank) ProcessConcurrentTransactions(transactions []TransactionRequest) {
    var wg sync.WaitGroup
    results := make(chan error, len(transactions))

    for _, tx := range transactions {
        wg.Add(1)
        go func(t TransactionRequest) {
            defer wg.Done()
            err := b.Transfer(t.From, t.To, t.Amount)
            results <- err
        }(tx)
    }

    wg.Wait()
    close(results)
}
```

---

## 🏦 Account Types Specs

### 1. Savings Account

- Withdrawal fee: $5
- Interest rate: 0.5% monthly
- Min balance: $0
- Max withdrawals: unlimited (but with fee)

### 2. Checking Account

- No fees
- No interest
- Can go negative: No
- Overdraft protection: No

### 3. Investment Account

- Early withdrawal penalty: 10% of amount
- Interest rate: 2% quarterly
- Min balance: $100
- Lock-in period: 90 days

---

## 📦 Bonus Features

- [ ] Interest calculation & auto-deposit
- [ ] Monthly statements
- [ ] Account freeze/lock
- [ ] Multiple owners
- [ ] Loan management
- [ ] Bill pay
- [ ] Investment portfolio tracking
- [ ] Rate comparison

---

## ✅ Checklist

- [ ] Project setup & DB connection
- [ ] Interfaces defined
- [ ] 3 account types implemented
- [ ] CRUD operations
- [ ] Validation rules
- [ ] Concurrent transaction handling
- [ ] Error handling
- [ ] Transaction logging
- [ ] Persistence to database
- [ ] Load from database on startup
- [ ] Tested (concurrent safety)

---

## 🔍 Test Scenarios

```
1. Create account → verify balance = 0
2. Deposit → verify balance increased
3. Withdraw → verify balance decreased & fee applied
4. Transfer between accounts → verify both updated
5. Concurrent transactions (100 deposits/withdrawals) → verify consistency
6. Restart program → verify data persisted
7. Insufficient balance → error
8. Invalid amounts → error
```

---

## 📊 Scoring Rubric (0-100)

- **Functionality (35%)**: All operations work, rules enforced
- **Interfaces (15%)**: Proper design, polymorphism
- **Database (15%)**: Persistence, queries
- **Concurrency (15%)**: Thread-safe, no race conditions
- **Error Handling (10%)**: Validation, edge cases
- **Testing (10%)**: Concurrent tests pass

---

## 🔗 Resources

- Interfaces: [effective-go#interfaces](https://golang.org/doc/effective_go#interfaces)
- SQLite: [go-sqlite3](https://github.com/mattn/go-sqlite3)
- Sync: [sync package](https://pkg.go.dev/sync)
