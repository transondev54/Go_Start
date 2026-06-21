# MINI PROJECTS GUIDE

## 📋 Giới thiệu

Có 3 mini projects trong Day 1 Go Learning Plan:

1. **BMI Calculator** (1 giờ) - Dễ
2. **Number Guessing Game** (1.5 giờ) - Bình thường
3. **Contact Manager** (2 giờ) - Khó

---

## 🎯 Quy trình

### Bước 1: Tạo thư mục

```bash
cd mini_projects/01_BMI_Calculator
mkdir source
cd source
go mod init bmi_calculator
code main.go
```

### Bước 2: Đọc README.md

Mỗi project có README chi tiết với:

- Mô tả & yêu cầu
- Ví dụ chạy
- Hướng dẫn bước bước
- Test cases
- Bonus features

### Bước 3: Viết code

- Gõ tay (không copy-paste)
- Thử modify & experiment
- Kiểm tra error messages

### Bước 4: Test

- Test với các input khác nhau
- Xử lý edge cases
- Kiểm tra error handling

### Bước 5: Submit

- Lưu file `main.go`
- Test lại lần cuối
- Submit để nhận feedback

---

## ✅ Checklist cho mỗi project

- [ ] Đã tạo thư mục
- [ ] Đã `go mod init`
- [ ] Đã đọc README.md
- [ ] Đã viết main.go
- [ ] Code compile không error
- [ ] Chạy được chương trình
- [ ] Test với ít nhất 3 test cases
- [ ] Xử lý errors
- [ ] Code sạch (formatting)
- [ ] Commit với git (optional)

---

## 📊 Scoring Rubric

### Mỗi project được chấm 0-100:

**Functionality (40%)**

- Tất cả tính năng hoạt động
- Output đúng format
- Edge cases xử lý

**Code Quality (30%)**

- Code sạch, dễ đọc
- Naming conventions
- DRY principle
- Comments khi cần

**Error Handling (20%)**

- Validate input
- Handle edge cases
- Error messages rõ ràng

**Bonus Features (10%)**

- Extra features
- Tối ưu code
- UX improvements

---

## 🌟 Expected Output Examples

### BMI Calculator

```
Nhập cân nặng (kg): 70
Nhập chiều cao (m): 1.75
=== KẾT QUẢ BMI ===
Cân nặng: 70.00 kg
Chiều cao: 1.75 m
BMI: 22.86
Phân loại: BÌNH THƯỜNG ✓
```

### Number Guessing Game

```
Chào mừng đến với Number Guessing Game!
Máy tính đã pick một số từ 1-100.

Lần đoán 1: 50
Số cao hơn. Thử lại!

Lần đoán 2: 75
Số thấp hơn. Thử lại!

Lần đoán 3: 63
Chúc mừng! Đúng rồi!
Bạn tìm ra trong 3 lần đoán. 🎉
```

### Contact Manager

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
```

---

## 🔧 Troubleshooting

### Error: `go: command not found`

- Go chưa cài hoặc PATH chưa setup
- Cài đặt lại Go từ golang.org

### Error: `package fmt is not in GOROOT`

- Quên `import "fmt"`
- Thêm: `import "fmt"`

### Error: `undefined: Scanln`

- Viết thiếu package name
- Đổi `Scanln()` thành `fmt.Scanln()`

### Error: `division by zero`

- Kiểm tra input validation
- Thêm check: `if divisor == 0`

### Chương trình không in gì

- Check output statement (fmt.Print...)
- Chạy lại với `go run main.go`

---

## 💡 TIPS

1. **Testing**: Test khi code từng function
2. **Error checking**: Kiểm tra `err != nil` sau mỗi I/O
3. **Formatting**: Dùng `fmt.Printf()` để format output
4. **Naming**: Dùng tên meaningful cho variables
5. **Functions**: Break code thành small functions
6. **Comments**: Thêm comment cho complex logic

---

## 🚀 WORKFLOW

```
1. Tạo project folder
           ↓
2. go mod init <name>
           ↓
3. Tạo main.go
           ↓
4. Viết package main & import
           ↓
5. Viết helper functions
           ↓
6. Viết func main()
           ↓
7. go run main.go (test)
           ↓
8. Fix errors & refine
           ↓
9. Final test
           ↓
10. Done! ✅
```

---

## 📁 File Structure

Mỗi project nên có:

```
project_name/
├── go.mod              # Tự động tạo
├── main.go             # Your code
├── README.md           # Documentation (provided)
└── (optional) utils.go # Helper functions
```

---

## 🎯 GOALS

Sau khi hoàn thành 3 projects:

- ✅ Hiểu luồng lập trình Go
- ✅ Tập thói quen code tốt
- ✅ Có confidence viết Go program
- ✅ Sẵn sàng cho Lesson tiếp theo

---

## 📧 SUBMISSION

Khi xong, hãy:

1. Kiểm tra code một lần nữa
2. Test tất cả edge cases
3. Ghi nhận bất cứ vấn đề
4. Submit files để nhận feedback

---

**Happy Coding! 💻**
