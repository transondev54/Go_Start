# 📋 CODE REVIEW: Contact Manager (Day 1 - Mini Project 3)

## ✅ ĐIỂM MẠNH

### 1. **Cấu trúc tốt**

- Tách function rõ ràng (addContact, updateContact, deleteContact, v.v.)
- Menu loop xử lý các lựa chọn
- JSON marshaling/unmarshaling cơ bản

### 2. **Features hoàn chỉnh**

- ✅ Add contact
- ✅ View all
- ✅ Edit (update)
- ✅ Delete
- ✅ File persistence (JSON)
- ✅ Menu loop

### 3. **Error handling có cơ bản**

- Kiểm tra contact not found
- Try-catch style cho file operations

### 4. **JSON integration**

- Marshal contact to JSON
- Save/load từ file

---

## ⚠️ VẤN ĐỀ & CÁCH SỬA

### 🔴 **VẤN ĐỀ 1: updateContact không cập nhật được**

**Nguyên nhân:** Sử dụng **value receiver** thay vì **pointer receiver**

```go
// ❌ SAAAI
func (c Contact) UpdateName(name string) {
    c.Name = name  // Chỉ thay đổi copy
}
```

**Cách sửa:**

```go
// ✅ ĐÚNG
func (c *Contact) UpdateName(name string) {
    c.Name = name  // Update thực tế
}
```

Trong `updateContact`, dùng:

```go
contacts[i].Name = updatedContact.Name  // Directly update slice
```

hoặc refactor thành:

```go
func (c *Contact) Update(updated Contact) {
    if updated.Name != "" {
        c.Name = updated.Name
    }
    // ...
}
```

---

### 🔴 **VẤN ĐỀ 2: Missing case 3 & 6**

```go
// ❌ Trong main(), thiếu:
case 3:
    // View Contact by ID
case 6:
    // Search Contact
```

**Cách sửa:**

```go
case 3:
    fmt.Println("View Contact by ID")
    fmt.Print("Enter ID: ")
    var id int
    fmt.Scan(&id)

    found := false
    for _, contact := range contacts {
        if contact.ID == id {
            fmt.Printf("ID: %d, Name: %s, Phone: %s, Email: %s\n",
                contact.ID, contact.Name, contact.Phone, contact.Email)
            found = true
            break
        }
    }
    if !found {
        fmt.Println("Contact not found.")
    }

case 6:
    fmt.Println("Search Contact")
    fmt.Print("Enter name to search: ")
    var searchName string
    fmt.Scan(&searchName)

    found := false
    for _, contact := range contacts {
        if strings.Contains(strings.ToLower(contact.Name),
                          strings.ToLower(searchName)) {
            fmt.Printf("ID: %d, Name: %s, Phone: %s, Email: %s\n",
                contact.ID, contact.Name, contact.Phone, contact.Email)
            found = true
        }
    }
    if !found {
        fmt.Println("No contacts found.")
    }
```

---

### 🟡 **VẤN ĐỀ 3: Input validation**

**Hiện tại:** Không kiểm tra empty input

```go
// ❌ Không validate
fmt.Print("Enter Name: ")
var name string
fmt.Scan(&name)
contact := ContactAddDTO{Name: name, ...}
addContact(contact)
```

**Cách sửa:**

```go
// ✅ Validate input
fmt.Print("Enter Name: ")
var name string
fmt.Scan(&name)

if name == "" {
    fmt.Println("Error: Name cannot be empty")
    continue
}

// Validate email format
if !strings.Contains(email, "@") {
    fmt.Println("Error: Invalid email format")
    continue
}
```

---

### 🟡 **VẤN ĐỀ 4: Kiểm tra file errors kỹ hơn**

```go
// ❌ Hiện tại
if contactsData, err := os.ReadFile("contacts.json"); err == nil {
    json.Unmarshal(contactsData, &contacts)
    // Ignore unmarshal error!
}
```

**Cách sửa:**

```go
// ✅ Kiểm tra lỗi kỹ
if contactsData, err := os.ReadFile("contacts.json"); err == nil {
    if err := json.Unmarshal(contactsData, &contacts); err != nil {
        fmt.Println("Error parsing JSON:", err)
    }
} else if !errors.Is(err, os.ErrNotExist) {
    fmt.Println("Error reading file:", err)
}
```

---

### 🟡 **VẤN ĐỀ 5: Code organization**

**Hiện tại:** Tất cả logic trong `main()`

**Cách sửa:** Tách thành struct & methods:

```go
type ContactManager struct {
    contacts []Contact
    nextID   int
    filename string
}

func NewContactManager(filename string) *ContactManager {
    cm := &ContactManager{filename: filename}
    cm.Load()
    return cm
}

func (cm *ContactManager) Add(name, phone, email string) error {
    if name == "" {
        return errors.New("name cannot be empty")
    }
    // ...
}

func (cm *ContactManager) Save() error {
    data, _ := json.Marshal(cm.contacts)
    return os.WriteFile(cm.filename, data, 0644)
}

func (cm *ContactManager) ShowMenu() {
    for {
        // Menu logic
    }
}
```

---

## 📝 REFACTORED CODE STRUCTURE

```go
package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "strings"
    "time"
)

// ============ Types ============
type Contact struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Phone     string    `json:"phone"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

type ContactManager struct {
    contacts []Contact
    nextID   int
    filename string
}

// ============ Constructor ============
func NewContactManager(filename string) *ContactManager {
    cm := &ContactManager{
        filename: filename,
        nextID:   1,
    }
    _ = cm.Load() // Ignore if file doesn't exist
    return cm
}

// ============ Methods ============
func (cm *ContactManager) Add(name, phone, email string) (*Contact, error) {
    if name == "" {
        return nil, errors.New("name cannot be empty")
    }
    if email == "" {
        return nil, errors.New("email cannot be empty")
    }
    if !strings.Contains(email, "@") {
        return nil, errors.New("invalid email format")
    }

    contact := Contact{
        ID:        cm.nextID,
        Name:      name,
        Phone:     phone,
        Email:     email,
        CreatedAt: time.Now(),
    }

    cm.contacts = append(cm.contacts, contact)
    cm.nextID++
    cm.Save()

    return &contact, nil
}

func (cm *ContactManager) GetByID(id int) (*Contact, error) {
    for i := range cm.contacts {
        if cm.contacts[i].ID == id {
            return &cm.contacts[i], nil
        }
    }
    return nil, errors.New("contact not found")
}

func (cm *ContactManager) Update(id int, name, phone, email string) error {
    contact, err := cm.GetByID(id)
    if err != nil {
        return err
    }

    if name != "" {
        contact.Name = name
    }
    if phone != "" {
        contact.Phone = phone
    }
    if email != "" && strings.Contains(email, "@") {
        contact.Email = email
    }

    return cm.Save()
}

func (cm *ContactManager) Delete(id int) error {
    for i, contact := range cm.contacts {
        if contact.ID == id {
            cm.contacts = append(cm.contacts[:i], cm.contacts[i+1:]...)
            return cm.Save()
        }
    }
    return errors.New("contact not found")
}

func (cm *ContactManager) GetAll() []Contact {
    return cm.contacts
}

func (cm *ContactManager) Search(name string) []Contact {
    var results []Contact
    for _, contact := range cm.contacts {
        if strings.Contains(strings.ToLower(contact.Name),
                           strings.ToLower(name)) {
            results = append(results, contact)
        }
    }
    return results
}

func (cm *ContactManager) Save() error {
    data, err := json.MarshalIndent(cm.contacts, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(cm.filename, data, 0644)
}

func (cm *ContactManager) Load() error {
    data, err := os.ReadFile(cm.filename)
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            return nil // File not found is OK
        }
        return err
    }

    if err := json.Unmarshal(data, &cm.contacts); err != nil {
        return err
    }

    if len(cm.contacts) > 0 {
        cm.nextID = cm.contacts[len(cm.contacts)-1].ID + 1
    }

    return nil
}

// ============ Main Menu ============
func (cm *ContactManager) ShowMenu() {
    for {
        fmt.Println("\n╔════════════════════════════╗")
        fmt.Println("║   Contact Manager v1.0     ║")
        fmt.Println("╚════════════════════════════╝")
        fmt.Println("1. Add Contact")
        fmt.Println("2. View All")
        fmt.Println("3. View Contact (by ID)")
        fmt.Println("4. Edit Contact")
        fmt.Println("5. Delete Contact")
        fmt.Println("6. Search Contact")
        fmt.Println("7. Exit")

        var choice int
        fmt.Print("Enter your choice: ")
        fmt.Scan(&choice)

        switch choice {
        case 1:
            cm.handleAdd()
        case 2:
            cm.handleViewAll()
        case 3:
            cm.handleViewByID()
        case 4:
            cm.handleEdit()
        case 5:
            cm.handleDelete()
        case 6:
            cm.handleSearch()
        case 7:
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("❌ Invalid choice")
        }
    }
}

func (cm *ContactManager) handleAdd() {
    fmt.Print("Enter Name: ")
    var name string
    fmt.Scan(&name)

    fmt.Print("Enter Phone: ")
    var phone string
    fmt.Scan(&phone)

    fmt.Print("Enter Email: ")
    var email string
    fmt.Scan(&email)

    contact, err := cm.Add(name, phone, email)
    if err != nil {
        fmt.Println("❌ Error:", err)
        return
    }

    fmt.Printf("✅ Contact added (ID: %d)\n", contact.ID)
}

func (cm *ContactManager) handleViewAll() {
    if len(cm.contacts) == 0 {
        fmt.Println("📭 No contacts found")
        return
    }

    fmt.Println("\n📇 All Contacts:")
    for _, c := range cm.contacts {
        fmt.Printf("[%d] %s (%s) - %s\n", c.ID, c.Name, c.Phone, c.Email)
    }
}

func (cm *ContactManager) handleViewByID() {
    var id int
    fmt.Print("Enter ID: ")
    fmt.Scan(&id)

    contact, err := cm.GetByID(id)
    if err != nil {
        fmt.Println("❌ Error:", err)
        return
    }

    fmt.Printf("📇 Contact:\n")
    fmt.Printf("  ID: %d\n  Name: %s\n  Phone: %s\n  Email: %s\n",
        contact.ID, contact.Name, contact.Phone, contact.Email)
}

func (cm *ContactManager) handleEdit() {
    var id int
    fmt.Print("Enter ID: ")
    fmt.Scan(&id)

    _, err := cm.GetByID(id)
    if err != nil {
        fmt.Println("❌ Error:", err)
        return
    }

    fmt.Print("Enter new name (empty to skip): ")
    var name string
    fmt.Scan(&name)

    fmt.Print("Enter new phone (empty to skip): ")
    var phone string
    fmt.Scan(&phone)

    fmt.Print("Enter new email (empty to skip): ")
    var email string
    fmt.Scan(&email)

    if err := cm.Update(id, name, phone, email); err != nil {
        fmt.Println("❌ Error:", err)
        return
    }

    fmt.Println("✅ Contact updated")
}

func (cm *ContactManager) handleDelete() {
    var id int
    fmt.Print("Enter ID: ")
    fmt.Scan(&id)

    if err := cm.Delete(id); err != nil {
        fmt.Println("❌ Error:", err)
        return
    }

    fmt.Println("✅ Contact deleted")
}

func (cm *ContactManager) handleSearch() {
    fmt.Print("Enter name to search: ")
    var name string
    fmt.Scan(&name)

    results := cm.Search(name)
    if len(results) == 0 {
        fmt.Println("📭 No contacts found")
        return
    }

    fmt.Printf("\n🔍 Search Results (%d found):\n", len(results))
    for _, c := range results {
        fmt.Printf("[%d] %s (%s) - %s\n", c.ID, c.Name, c.Phone, c.Email)
    }
}

func main() {
    manager := NewContactManager("contacts.json")
    manager.ShowMenu()
}
```

---

## 📊 SCORING

| Aspect             | Score      | Notes                                              |
| ------------------ | ---------- | -------------------------------------------------- |
| **Functionality**  | 85/100     | Tất cả core features hoạt động, missing case 3 & 6 |
| **Code Quality**   | 70/100     | Có thể refactor tốt hơn, organize better           |
| **Error Handling** | 75/100     | Có error checking, nhưng chưa complete             |
| **Persistence**    | 90/100     | JSON save/load hoạt động tốt                       |
| **Testing**        | 60/100     | Không có unit tests                                |
| **Overall**        | **76/100** | Good start! Implement suggestions to improve       |

---

## 🎯 ĐỀ XUẤT CÁCH SỬA

1. ✅ Thêm case 3 & 6
2. ✅ Sửa pointer receivers cho Update
3. ✅ Thêm input validation
4. ✅ Refactor thành ContactManager struct
5. ✅ Thêm unit tests
6. ✅ Tốt lên error messages

---

## 🚀 NEXT STEPS

- Áp dụng những feedback này
- Hoàn thành Day 2 mini projects
- Học thêm về Interfaces & Concurrency
- Build larger projects sử dụng những kỹ năng mới
