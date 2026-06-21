# Mini Project 3: Contact Manager

## 📋 Mô tả dự án

Viết ứng dụng quản lý danh bạ (liên hệ). Người dùng có thể thêm, xem, chỉnh sửa, xóa danh bạ. Dữ liệu được lưu trong memory (sử dụng slice của structs).

---

## 🎯 Yêu cầu (Requirements)

### Tính năng chính

- [ ] **Add Contact**: Thêm danh bạ mới (tên, số điện thoại, email)
- [ ] **View All**: Xem tất cả danh bạ
- [ ] **View By ID**: Xem chi tiết 1 danh bạ
- [ ] **Edit Contact**: Chỉnh sửa thông tin danh bạ
- [ ] **Delete Contact**: Xóa danh bạ
- [ ] **Search**: Tìm danh bạ theo tên
- [ ] **Menu**: Giao diện menu lựa chọn

### Data Structure

```go
type Contact struct {
    ID    int
    Name  string
    Phone string
    Email string
}
```

---

## 📝 Ví dụ chương trình

```
╔════════════════════════════╗
║   CONTACT MANAGER v1.0     ║
╚════════════════════════════╝

1. Add Contact
2. View All
3. View Contact (by ID)
4. Edit Contact
5. Delete Contact
6. Search Contact
7. Exit

Choose option: 1

Enter Name: Alice
Enter Phone: 0912345678
Enter Email: alice@example.com

Contact added successfully! (ID: 1)

Choose option: 2

======== ALL CONTACTS ========
ID: 1 | Name: Alice | Phone: 0912345678 | Email: alice@example.com

Choose option: 7
Goodbye!
```

---

## 🔨 Hướng dẫn thực hiện

### Bước 1: Khởi tạo project

```bash
mkdir contact_manager
cd contact_manager
go mod init contact_manager
```

### Bước 2: Struct definition

```go
package main

type Contact struct {
    ID    int
    Name  string
    Phone string
    Email string
}

// Global slice to store contacts
var contacts []Contact
var nextID = 1
```

### Bước 3: Core functions

```go
// Add contact
func addContact(name, phone, email string) {
    contact := Contact{
        ID:    nextID,
        Name:  name,
        Phone: phone,
        Email: email,
    }
    contacts = append(contacts, contact)
    nextID++
    fmt.Println("Contact added!")
}

// View all
func viewAll() {
    if len(contacts) == 0 {
        fmt.Println("No contacts found!")
        return
    }
    for _, c := range contacts {
        fmt.Printf("ID: %d | Name: %s | Phone: %s\n", c.ID, c.Name, c.Phone)
    }
}

// Delete contact
func deleteContact(id int) {
    for i, c := range contacts {
        if c.ID == id {
            contacts = append(contacts[:i], contacts[i+1:]...)
            fmt.Println("Contact deleted!")
            return
        }
    }
    fmt.Println("Contact not found!")
}

// Search contact
func searchContact(name string) {
    found := false
    for _, c := range contacts {
        if strings.Contains(c.Name, name) {
            fmt.Printf("ID: %d | Name: %s | Phone: %s | Email: %s\n",
                c.ID, c.Name, c.Phone, c.Email)
            found = true
        }
    }
    if !found {
        fmt.Println("No contacts found!")
    }
}
```

### Bước 4: Main menu

```go
func main() {
    for {
        fmt.Println("\n╔════════════════════════════╗")
        fmt.Println("║   CONTACT MANAGER v1.0     ║")
        fmt.Println("╚════════════════════════════╝")
        fmt.Println("1. Add Contact")
        fmt.Println("2. View All")
        fmt.Println("3. View Contact (by ID)")
        fmt.Println("4. Edit Contact")
        fmt.Println("5. Delete Contact")
        fmt.Println("6. Search Contact")
        fmt.Println("7. Exit")

        var choice int
        fmt.Print("\nChoose option: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            // Add contact
        case 2:
            // View all
        case 7:
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Invalid option!")
        }
    }
}
```

---

## 💡 Gợi ý thêm

1. **Slice manipulation**: Xóa element bằng `append(slice[:i], slice[i+1:]...)`
2. **Search**: Dùng `strings.Contains()` hoặc `strings.EqualFold()`
3. **Edit**: Tìm index, modify fields, không cần re-add
4. **ID management**: Tăng `nextID` mỗi lần thêm
5. **Input validation**: Kiểm tra empty strings

---

## 🧪 Test cases

| Case                  | Expected         |
| --------------------- | ---------------- |
| Add 1 contact         | Success          |
| View all              | Show 1 contact   |
| Search by name        | Find contact     |
| Edit contact          | Update fields    |
| Delete contact        | Remove from list |
| View all after delete | Show remaining   |

---

## 🌟 Bonus (thêm điểm)

- [ ] Lưu contacts vào file (JSON)
- [ ] Tải contacts từ file khi khởi động
- [ ] Export to CSV
- [ ] Sort by name/phone
- [ ] Validate phone number format
- [ ] Duplicate check (phone)
- [ ] Pagination (nếu nhiều contacts)
- [ ] Advanced search (multiple fields)

---

## 📊 Đánh giá

| Tiêu chí                    | Điểm |
| --------------------------- | ---- |
| Core Functionality          | 40%  |
| Code Quality & Organization | 30%  |
| User Interface              | 15%  |
| Bonus Features              | 15%  |

---

## 🎓 Kiến thức áp dụng

- ✅ Structs & slices
- ✅ Functions & methods
- ✅ Switch statements & menus
- ✅ Loops & iteration
- ✅ String operations
- ✅ Input/Output
- ✅ Error handling
- ✅ Data manipulation

---

**Deadline:** 2 giờ
**Start:** Sau Lesson 7 ✓

---

## 📁 File structure

```
contact_manager/
├── go.mod
├── main.go           # All code
├── README.md         # This file
└── contacts.json     # Data file (bonus)
```
