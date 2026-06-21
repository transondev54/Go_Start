# Lesson 5: Database Basics - SQLite & SQL Operations

## 📖 Nội dung bài học

1. Giới thiệu SQLite
2. Kết nối database
3. CRUD operations - Create, Read, Update, Delete
4. Prepared statements & SQL injection
5. Transactions

---

## 1️⃣ GIỚI THIỆU SQLITE

### SQLite là gì?

- ✅ Lightweight, serverless database
- ✅ Lưu dữ liệu trong file `.db`
- ✅ Perfect cho desktop/mobile apps
- ✅ Built-in Go support qua `database/sql`

### Cài đặt driver

```bash
go get github.com/mattn/go-sqlite3
```

---

## 2️⃣ KẾT NỐI DATABASE

### Import packages

```go
import (
    \"database/sql\"
    _ \"github.com/mattn/go-sqlite3\"
)
```

### Tạo & mở database

```go
import (
    \"database/sql\"
    _ \"github.com/mattn/go-sqlite3\"
)

func main() {
    // Tạo/mở database
    db, err := sql.Open(\"sqlite3\", \"mydatabase.db\")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Kiểm tra connection
    if err := db.Ping(); err != nil {
        panic(err)
    }

    fmt.Println(\"Connected!\")
}
```

### Schema - Tạo bảng

```go
func CreateTableIfNotExists(db *sql.DB) error {
    createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT UNIQUE NOT NULL,
        age INTEGER
    );
    `
    _, err := db.Exec(createTableSQL)
    return err
}
```

---

## 3️⃣ CRUD OPERATIONS

### CREATE - Thêm dữ liệu

```go
func AddUser(db *sql.DB, name, email string, age int) error {
    // Prepared statement - safer
    stmt, err := db.Prepare(\"INSERT INTO users (name, email, age) VALUES (?, ?, ?)\")
    if err != nil {
        return err
    }
    defer stmt.Close()

    result, err := stmt.Exec(name, email, age)
    if err != nil {
        return err
    }

    id, _ := result.LastInsertId()
    fmt.Printf(\"User added with ID: %d\\n\", id)
    return nil
}
```

### READ - Lấy dữ liệu

```go
type User struct {
    ID    int
    Name  string
    Email string
    Age   int
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
    user := &User{}

    row := db.QueryRow(\"SELECT id, name, email, age FROM users WHERE id = ?\", id)

    err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf(\"user not found\")
        }
        return nil, err
    }

    return user, nil
}

func GetAllUsers(db *sql.DB) ([]User, error) {
    rows, err := db.Query(\"SELECT id, name, email, age FROM users\")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User

    for rows.Next() {
        user := User{}
        err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, rows.Err()
}
```

### UPDATE - Cập nhật dữ liệu

```go
func UpdateUser(db *sql.DB, id int, name, email string) error {
    stmt, err := db.Prepare(\"UPDATE users SET name = ?, email = ? WHERE id = ?\")
    if err != nil {
        return err
    }
    defer stmt.Close()

    result, err := stmt.Exec(name, email, id)
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf(\"user not found\")
    }

    return nil
}
```

### DELETE - Xóa dữ liệu

```go
func DeleteUser(db *sql.DB, id int) error {
    stmt, err := db.Prepare(\"DELETE FROM users WHERE id = ?\")
    if err != nil {
        return err
    }
    defer stmt.Close()

    result, err := stmt.Exec(id)
    if err != nil {
        return err
    }

    rowsAffected, _ := result.RowsAffected()
    if rowsAffected == 0 {
        return fmt.Errorf(\"user not found\")
    }

    return nil
}
```

---

## 4️⃣ PREPARED STATEMENTS - SQL INJECTION PREVENTION

### Sai - SQL Injection

```go
// ❌ KHÔNG LÀM ĐIỀU NÀY
func UnsafeQuery(db *sql.DB, name string) error {
    query := fmt.Sprintf(\"SELECT * FROM users WHERE name = '%s'\", name)
    // Nếu name = \"' OR '1'='1\" => Query bị hỏng!
    rows, _ := db.Query(query)
    // ...
}
```

### Đúng - Prepared Statements

```go
// ✅ LÀM CÁI NÀY
func SafeQuery(db *sql.DB, name string) error {
    // Dấu ? là placeholder
    rows, err := db.Query(\"SELECT * FROM users WHERE name = ?\", name)
    if err != nil {
        return err
    }
    defer rows.Close()

    // Go tự động escape các giá trị
    // ...
}
```

---

## 5️⃣ TRANSACTIONS

### Atomic operations

```go
func TransferMoney(db *sql.DB, fromID, toID int, amount float64) error {
    // Bắt đầu transaction
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    // Nếu có error, rollback
    defer func() {
        if err != nil {
            tx.Rollback()
        }
    }()

    // Giảm account từ
    _, err = tx.Exec(\"UPDATE accounts SET balance = balance - ? WHERE id = ?\", amount, fromID)
    if err != nil {
        return err
    }

    // Tăng account tới
    _, err = tx.Exec(\"UPDATE accounts SET balance = balance + ? WHERE id = ?\", amount, toID)
    if err != nil {
        return err
    }

    // Commit nếu mọi thứ OK
    return tx.Commit().Error
}
```

---

## 📝 TÓM TẮT DATABASE

```go
// Kết nối
db, _ := sql.Open(\"sqlite3\", \"file.db\")

// CREATE
stmt, _ := db.Prepare(\"INSERT INTO users VALUES (?, ?)\")
stmt.Exec(name, email)

// READ
row := db.QueryRow(\"SELECT * FROM users WHERE id = ?\", id)
row.Scan(&id, &name, &email)

// UPDATE
db.Exec(\"UPDATE users SET name = ? WHERE id = ?\", name, id)

// DELETE
db.Exec(\"DELETE FROM users WHERE id = ?\", id)
```

---

## 💡 BEST PRACTICES

1. **Luôn dùng prepared statements** - chống SQL injection
2. **Defer rows.Close()** - tránh resource leak
3. **Kiểm tra ErrNoRows** - phân biệt error vs no data
4. **Dùng transactions** - cho multi-step operations
5. **Connection pooling** - Go tự động quản lý

---

## 🎯 BƯỚC TIẾP THEO

- Đọc **Lesson 6** về Testing
- Tạo database cho mini project
- Thực hành CRUD operations
