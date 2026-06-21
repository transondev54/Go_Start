# Mini Project 1: Task Manager with File Persistence

## 📝 Mô tả

Xây dựng một task manager app với các tính năng:

- Thêm tasks (title, description, status)
- Xem tất cả tasks
- Filter tasks by status (pending/completed)
- Đánh dấu task hoàn thành
- Xóa tasks
- **Lưu & tải tasks từ JSON file**

---

## 📋 Yêu cầu

### Tính năng bắt buộc

1. **Task struct** với fields:
   - `ID` (int)
   - `Title` (string)
   - `Description` (string)
   - `Status` (string - "pending" hoặc "completed")
   - `CreatedAt` (time.Time)

2. **File I/O**:
   - Lưu tasks vào `tasks.json` sau mỗi thay đổi
   - Load tasks từ `tasks.json` khi startup
   - Handle case file không tồn tại

3. **CRUD Operations**:
   - Add task
   - View all tasks
   - View pending tasks
   - View completed tasks
   - Mark task as completed
   - Delete task

4. **Error Handling**:
   - Validate input (không để title rỗng)
   - Handle file operations
   - Custom error types

### Ví dụ output

```
╔═══════════════════════════════╗
║     Task Manager v1.0         ║
╚═══════════════════════════════╝

1. Add Task
2. View All Tasks
3. View Pending Tasks
4. View Completed Tasks
5. Mark Task Complete
6. Delete Task
7. Exit

Enter choice: 1
Enter task title: Learn Go Interfaces
Enter description: Study interfaces and polymorphism
Task added (ID: 1)

Enter choice: 2
[1] Learn Go Interfaces (pending)
    Learn interfaces and polymorphism
    Created: 2024-01-15 10:30:00

[2] Buy Groceries (completed)
    Milk, eggs, bread
    Created: 2024-01-15 09:00:00
```

---

## 🎯 Learning Objectives

- ✅ Pointers & method receivers
- ✅ JSON marshaling/unmarshaling
- ✅ File I/O operations
- ✅ Error handling & custom errors
- ✅ Code organization & functions
- ✅ Testing basics

---

## 📚 Bước thực hiện

### Bước 1: Setup project

```bash
mkdir task_manager
cd task_manager
go mod init task_manager
code main.go
```

### Bước 2: Định nghĩa structs

```go
type Task struct {
    ID          int
    Title       string
    Description string
    Status      string
    CreatedAt   time.Time
}

type TaskManager struct {
    tasks  []Task
    nextID int
}
```

### Bước 3: Implement functions

- `NewTaskManager()` - khởi tạo
- `AddTask(title, description string)` - thêm task
- `MarkComplete(id int)` - đánh dấu hoàn thành
- `DeleteTask(id int)` - xóa task
- `GetPendingTasks()` - lấy pending tasks
- `SaveToFile(filename string)` - lưu JSON
- `LoadFromFile(filename string)` - tải JSON

### Bước 4: Main menu

```go
func (tm *TaskManager) ShowMenu() {
    for {
        // Hiển thị menu options
        // Đọc input
        // Execute operations
    }
}
```

### Bước 5: Testing

```go
func TestAddTask(t *testing.T) {
    tm := NewTaskManager()
    tm.AddTask("Test", "Description")

    if len(tm.tasks) != 1 {
        t.Error("Task not added")
    }
}
```

---

## 📦 Bonus Features

- [ ] Search tasks by title
- [ ] Edit task description
- [ ] Due dates & reminders
- [ ] Task priorities
- [ ] Statistics (completed %, total tasks)
- [ ] Sort tasks by date/status
- [ ] CSV export

---

## ✅ Checklist

- [ ] Project tạo & go mod setup
- [ ] Struct định nghĩa
- [ ] Add task functionality
- [ ] File save/load
- [ ] Tất cả CRUD operations
- [ ] Menu loop hoạt động
- [ ] Error handling
- [ ] Tested (ít nhất 3 test cases)
- [ ] Code clean & formatted
- [ ] Commit với git

---

## 🔍 Testing Guide

```bash
# Run program
go run main.go

# Test cases
1. Add task → Save to file → Restart program → Verify task loaded
2. Add multiple tasks → Filter by status → Verify results
3. Mark task complete → Verify status changed
4. Delete task → Verify removed from list
```

---

## 📊 Scoring Rubric (0-100)

- **Functionality (40%)**: Tất cả features hoạt động
- **Code Quality (30%)**: Clean, organized, readable
- **Error Handling (20%)**: Validate input, handle edge cases
- **Testing (10%)**: Unit tests cho core functions
- **Bonus (10%)**: Extra features, nice UX

---

## 🔗 Resources

- JSON in Go: [encoding/json](https://pkg.go.dev/encoding/json)
- File I/O: [os package](https://pkg.go.dev/os)
- Testing: [testing package](https://pkg.go.dev/testing)
