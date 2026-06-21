# ⭐ START HERE - GO LEARNING DAY 2

## 🎯 Chào mừng đến Day 2!

Bạn đã hoàn thành **Day 1** - Nền tảng Go! 🎉

**Day 2** sẽ đưa bạn từ beginner → intermediate developer, với focus vào:

- ✅ **Interfaces** - Design patterns & polymorphism
- ✅ **Pointers** - Memory & references
- ✅ **Error handling** - Production-ready error management
- ✅ **Goroutines** - Concurrent programming
- ✅ **Database** - Persist data with SQLite
- ✅ **Testing** - Unit tests & quality assurance

---

## 📅 LỊCH TRÌNH NHANH

| Thời gian | Nội dung                     | Thời lượng |
| --------- | ---------------------------- | ---------- |
| **0h**    | Lesson 1: Interfaces         | 1 giờ      |
| **1h**    | Lesson 2: Pointers           | 1 giờ      |
| **2h**    | Mini Project 1: Task Manager | 1.5 giờ    |
| **3.5h**  | Lesson 3: Error Handling     | 1 giờ      |
| **4.5h**  | Lesson 4: Goroutines         | 1.5 giờ    |
| **6h**    | Mini Project 2: Weather App  | 2 giờ      |
| **8h**    | Lesson 5: Database           | 1 giờ      |
| **9h**    | Lesson 6: JSON & HTTP        | 1 giờ      |
| **10h**   | Mini Project 3: Bank System  | 2 giờ      |
| **12h**   | Lesson 7: Testing            | 1 giờ      |
| **13h**   | Review & Polish              | 1-2 giờ    |

---

## 📚 CÓ GÌ TRONG DAY 2

### 📖 7 Bài giảng

1. **Interfaces** - Thiết kế polymorphic systems
2. **Pointers** - Memory management & receivers
3. **Error Handling** - Custom errors & wrapping
4. **Goroutines** - Concurrent programming
5. **Database** - SQLite & SQL operations
6. **JSON & HTTP** - API integration
7. **Testing** - Unit tests & quality

### 🏗️ 3 Mini Projects

| Project          | Level    | Focus                      |
| ---------------- | -------- | -------------------------- |
| **Task Manager** | ⭐⭐     | File I/O, Pointers, JSON   |
| **Weather App**  | ⭐⭐⭐   | Goroutines, Channels, HTTP |
| **Bank System**  | ⭐⭐⭐⭐ | Interfaces, Database, Sync |

---

## 🚀 GETTING STARTED

### Step 1: Review Day 1 (5-10 min)

Hãy review lại những gì bạn học:

- Variables, types, functions
- Control flow, slices, maps
- Structs & methods
- File I/O basics

**Files to review:**

- [`Day 1/lectures/`](../Day%201/lectures)
- Your mini projects code

### Step 2: Code Review for Contact Manager

Xem feedback cho Day 1 project của bạn:

- [`Day 1/mini_projects/03_Contact_Manager/CODE_REVIEW.md`](../Day%201/mini_projects/03_Contact_Manager/CODE_REVIEW.md)

**Key fixes:**

- Pointer receivers for Update
- Add missing features (case 3, 6)
- Better error handling
- Input validation

### Step 3: Start Learning

Begin with **Lesson 1: Interfaces**

```bash
# Open the lessons
code lectures/01_Interfaces.md
```

### Step 4: Code Along

- Gõ tay tất cả code examples
- Modify & experiment
- Try-error là bình thường
- Đừng copy-paste

### Step 5: Mini Projects

Sau mỗi 2 lessons, làm mini project tương ứng:

```bash
cd mini_projects/01_Task_Manager/task_manager
go mod init task_manager
code main.go
```

---

## 💻 DEV ENVIRONMENT SETUP

### Prerequisites

```bash
# Check Go installation
go version

# Update Go (if needed)
# Visit: https://golang.org/dl
```

### Tools (Recommended)

```bash
# Install air (auto-reload)
go install github.com/cosmtrek/air@latest

# Install golangci-lint (linting)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install dlv (debugger)
go install github.com/go-delve/delve/cmd/dlv@latest
```

### Useful Commands

```bash
# Run program
go run main.go

# Build
go build -o myapp

# Test
go test
go test -v
go test -cover

# Format
go fmt ./...

# Lint
golangci-lint run

# Vet (static analyzer)
go vet ./...
```

---

## 📖 HOW TO USE LESSONS

### Mỗi lesson có

✅ Definitions & concepts
✅ Code examples (gõ tay!)
✅ Common mistakes
✅ Best practices
✅ Real-world examples

### Study method

```
1. READ the theory
2. TYPE examples in your editor
3. RUN & test code
4. MODIFY & experiment
5. UNDERSTAND the output
6. MOVE to next section
```

---

## 🎯 DAILY GOALS

### Day 2 Morning (0-6 hours)

- [ ] Complete Lesson 1 (Interfaces)
- [ ] Complete Lesson 2 (Pointers)
- [ ] Complete Lesson 3 (Error Handling)
- [ ] Start Mini Project 1 (Task Manager)

### Day 2 Afternoon (6-12 hours)

- [ ] Finish Mini Project 1
- [ ] Complete Lesson 4 (Goroutines)
- [ ] Complete Lesson 5 (Database)
- [ ] Complete Lesson 6 (JSON & HTTP)
- [ ] Start Mini Project 2 (Weather App)

### Day 2 Evening (12-16 hours)

- [ ] Finish Mini Project 2
- [ ] Complete Lesson 7 (Testing)
- [ ] Start Mini Project 3 (Bank System)
- [ ] Write tests for projects

---

## 📋 CHECKLIST BEFORE START

- [ ] Go installed & working
- [ ] Editor configured (VS Code + Go extension)
- [ ] Day 1 code reviewed
- [ ] Understand Day 1 concepts
- [ ] Ready to learn advanced concepts?

---

## ❓ FREQUENTLY ASKED

**Q: Cần phải làm tất cả projects?**

- Recommended: Yes, tất cả 3
- Minimum: Ít nhất 2

**Q: Có thể bỏ qua lessons?**

- No, lessons có prerequisites
- Mỗi lesson build on previous
- Don't skip!

**Q: Bao lâu để hoàn thành Day 2?**

- 12-16 hours total
- With breaks: 2-3 days

**Q: Nếu không hiểu?**

1. Read lesson lại
2. Check examples
3. Try coding
4. Google specific concept
5. Ask for help

**Q: Có deadline?**

- No, learn at your own pace
- Consistency > speed
- Better deep understanding than rushing

---

## 🎓 SUCCESS TIPS

✅ **Be consistent** - Code mỗi ngày
✅ **Type everything** - Không copy-paste
✅ **Understand errors** - Don't just fix, understand
✅ **Experiment** - Modify code, break things
✅ **Take breaks** - Every 90 min
✅ **Review code** - Others' + your own
✅ **Ask questions** - When stuck

---

## 📚 LEARNING MATERIALS

### Inside this folder

- `lectures/` - 7 detailed lessons
- `mini_projects/` - 3 projects with READMEs
- `DAY_2_LEARNING_PLAN.md` - Detailed curriculum

### External resources

- [Go Tour](https://tour.golang.org) - Interactive
- [Effective Go](https://golang.org/doc/effective_go) - Best practices
- [Go by Example](https://gobyexample.com) - Practical examples

---

## 🚀 LET'S GET STARTED!

When you're ready:

```bash
# Navigate to Day 2
cd Day\ 2

# Open first lesson
code lectures/01_Interfaces.md

# Read, learn, code along!
```

---

## 🎉 YOU'VE GOT THIS!

Day 2 is more challenging, but you're ready. You've completed the fundamentals!

By the end of Day 2, you'll understand:

- How to design systems with interfaces
- How to write concurrent code
- How to handle errors gracefully
- How to persist data
- How to test your code

**Let's build production-ready Go code!** 💪

---

Questions? Check the lesson files or review code examples.

Happy learning! 🚀
