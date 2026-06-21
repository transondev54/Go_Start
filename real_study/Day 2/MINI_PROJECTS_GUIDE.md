# MINI PROJECTS GUIDE - DAY 2

## 📋 Giới thiệu

Có 3 mini projects trong Day 2 Go Learning Plan:

1. **Task Manager** (1.5 tiếng) - Bình thường
2. **Weather App** (2 tiếng) - Khó
3. **Bank System** (2 tiếng) - Khó (++advanced)

---

## 🎯 Quy trình chung

### Bước 1: Tạo thư mục

```bash
cd mini_projects/01_Task_Manager
mkdir task_manager
cd task_manager
go mod init task_manager
code main.go
```

### Bước 2: Đọc README.md

Mỗi project có README chi tiết với:

- Mô tả & yêu cầu
- Learning objectives
- Step-by-step implementation
- Ví dụ output
- Bonus features

### Bước 3: Viết code

- Gõ tay (không copy-paste)
- Thử modify & experiment
- Kiểm tra error messages
- Test khi viết

### Bước 4: Test

- Test với các input khác nhau
- Xử lý edge cases
- Kiểm tra error handling
- Verify file persistence (nếu có)

### Bước 5: Refactor & improve

- Clean up code
- Add comments
- Remove duplication
- Optimize if needed

### Bước 6: Submit

- Lưu file `main.go`
- Test lại lần cuối
- Commit với git
- Submit để nhận feedback

---

## 📊 PROJECT COMPARISON

| Project      | Difficulty | Time | Core Skills                |
| ------------ | ---------- | ---- | -------------------------- |
| Task Manager | ⭐⭐       | 1.5h | Pointers, JSON, File I/O   |
| Weather App  | ⭐⭐⭐     | 2h   | Goroutines, Channels, HTTP |
| Bank System  | ⭐⭐⭐⭐   | 2h   | Interfaces, Database, Sync |

---

## ✅ CHECKLIST cho mỗi project

- [ ] Tạo project directory
- [ ] `go mod init`
- [ ] Đọc README.md chi tiết
- [ ] Viết main.go từ scratch
- [ ] Code compile không error
- [ ] Chạy được chương trình
- [ ] Test với ít nhất 5 test cases
- [ ] Xử lý errors & edge cases
- [ ] Code formatted & readable
- [ ] Add comments cho functions
- [ ] Commit với git (optional)

---

## 🔄 DAY 2 WORKFLOW

```
Lesson 1 (Interfaces) → Quick practice
      ↓
Lesson 2 (Pointers)  → Task Manager project start
      ↓
Lesson 3 (Errors)    → Add error handling to Task Manager
      ↓
Lesson 4 (Goroutines) → Weather App project start
      ↓
Lesson 5 (Database)  → Bank System project start
      ↓
Lesson 6 (JSON/HTTP) → Weather App API integration
      ↓
Lesson 7 (Testing)   → Add tests to all projects
      ↓
Review & Polish      → Final touches
```

---

## 💡 TIPS FOR SUCCESS

1. **Không vội** - Hiểu concept trước khi code
2. **Gõ tay tất cả** - Không copy-paste từ README
3. **Experiment** - Modify code, try different approaches
4. **Error messages** - Read & understand lỗi
5. **Commit early** - Commit sau mỗi feature hoàn thành
6. **Test often** - Test khi viết, không chờ cuối

---

## 📚 RESOURCES

### Go Official

- [Go Tour](https://tour.golang.org)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Standard Library](https://pkg.go.dev)

### Recommended Reading

- Interfaces: [A Tour of Go - Interfaces](https://tour.golang.org/methods/9)
- Goroutines: [Concurrency - Go by Example](https://gobyexample.com/goroutines)
- Database: [database/sql](https://pkg.go.dev/database/sql)

---

## ❓ FAQ

**Q: Nếu stuck ở đâu?**

- Read error message kỹ
- Xem relevant lesson lại
- Try modify code, see what happens
- Ask for help if really stuck

**Q: Nên làm projects theo thứ tự?**

- Yes! Mỗi project learn mới concepts
- Task Manager → Weather → Bank
- Tránh jump around

**Q: Có debug tools không?**

- `fmt.Printf` cho basic debugging
- Go debugger: `dlv` (advanced)
- Print statements là OK cho now

**Q: Bao lâu để hoàn thành?**

- Task Manager: 1-1.5 giờ
- Weather App: 2 giờ
- Bank System: 2-2.5 giờ
- Total: ~5-6 giờ (không tính breaks)

---

## 🎓 LEARNING PATH

```
Day 1: Basics (Fundamentals)
├── Variables & Types
├── Control Flow
├── Functions
└── Mini Projects (3)

Day 2: Advanced (Production-ready)
├── Interfaces & Polymorphism
├── Pointers & Memory
├── Error Handling
├── Concurrency (Goroutines/Channels)
├── Persistence (Database)
└── Mini Projects (3) ← YOU ARE HERE
```

---

## 🎯 NEXT STEPS AFTER DAY 2

- Day 3: Web Framework (Gin/Echo)
- Day 4: API Development
- Day 5: Deployment & DevOps
- Day 6: Performance Optimization
- Day 7: Real-world project

---

Good luck! 🚀
