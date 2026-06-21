# Mini Project 2: Number Guessing Game

## 📋 Mô tả dự án

Viết game đoán số: Máy tính pick một số ngẫu nhiên từ 1-100, người chơi phải đoán ra số đó. Mỗi lần đoán, máy sẽ cho feedback "cao hơn" hoặc "thấp hơn". Đếm số lần đoán.

---

## 🎯 Yêu cầu (Requirements)

### Tính năng chính

- [ ] Máy tính pick random number từ 1-100
- [ ] Người chơi nhập số để đoán
- [ ] Feedback:
  - Nếu đoán < số: "Số bạn đoán thấp hơn"
  - Nếu đoán > số: "Số bạn đoán cao hơn"
  - Nếu đoán = số: "Chúc mừng! Bạn đã đoán đúng!"
- [ ] Đếm số lần đoán
- [ ] Hiển thị kết quả cuối

### Xử lý lỗi

- [ ] Kiểm tra số trong range 1-100
- [ ] Xử lý input không hợp lệ
- [ ] Gợi ý nếu đoán quá chênh lệch (ví dụ: <1 hoặc >100)

---

## 📝 Ví dụ chương trình

```
Chào mừng đến với Number Guessing Game!
Máy tính đã pick một số từ 1-100. Hãy đoán xem!

Lần đoán 1: 50
Số bạn đoán cao hơn. Thử lại!

Lần đoán 2: 25
Số bạn đoán thấp hơn. Thử lại!

Lần đoán 3: 37
Chúc mừng! Bạn đã đoán đúng số 37!

Bạn tìm ra câu trả lời trong 3 lần đoán. 🎉
```

---

## 🔨 Hướng dẫn thực hiện

### Bước 1: Khởi tạo project

```bash
mkdir number_guessing_game
cd number_guessing_game
go mod init number_guessing_game
```

### Bước 2: Generate random number

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Seed random với thời gian hiện tại
    rand.Seed(time.Now().UnixNano())

    // Generate random số từ 1-100
    randomNumber := rand.Intn(100) + 1

    fmt.Println("Secret number:", randomNumber)  // Debugging
}
```

### Bước 3: Game loop

```go
func main() {
    rand.Seed(time.Now().UnixNano())
    secretNumber := rand.Intn(100) + 1

    attempts := 0
    guessed := false

    for !guessed {
        var guess int

        fmt.Print("Nhập số của bạn (1-100): ")
        _, err := fmt.Scanln(&guess)

        if err != nil {
            fmt.Println("Vui lòng nhập một số hợp lệ!")
            continue
        }

        attempts++

        if guess == secretNumber {
            fmt.Printf("Chúc mừng! Bạn đoán đúng số %d!\n", secretNumber)
            guessed = true
        } else if guess < secretNumber {
            fmt.Println("Số của bạn thấp hơn. Thử lại!")
        } else {
            fmt.Println("Số của bạn cao hơn. Thử lại!")
        }
    }

    fmt.Printf("Bạn tìm ra câu trả lời trong %d lần đoán. 🎉\n", attempts)
}
```

---

## 💡 Gợi ý thêm

1. **Random**: Import `math/rand` và seed với `time.Now().UnixNano()`
2. **Loop**: Dùng `for !guessed` hoặc vô hạn với `break`
3. **Validation**: Kiểm tra 1 ≤ guess ≤ 100
4. **Attempts**: Đếm mỗi lần lặp hợp lệ
5. **UX**: Gợi ý "Quá xa!" nếu sai > 20 số

---

## 🧪 Test cases

| Case                | Expected       |
| ------------------- | -------------- |
| Đoán lần 1 đúng     | 1 attempt      |
| Đoán từ từ          | Nhiều attempts |
| Nhập text           | Error message  |
| Nhập số ngoài 1-100 | Warning        |

---

## 🌟 Bonus (thêm điểm)

- [ ] Khó độ (dễ: 1-50, bình thường: 1-100, khó: 1-1000)
- [ ] Hint: "Hotter/Colder" dựa trên attempts trước
- [ ] High score: Lưu kỷ lục ít attempts nhất
- [ ] Multi-round: Chơi nhiều ván
- [ ] Difficulty tips

---

## 📊 Đánh giá

| Tiêu chí         | Điểm |
| ---------------- | ---- |
| Game Logic       | 40%  |
| Code Quality     | 30%  |
| Input Validation | 20%  |
| Bonus Features   | 10%  |

---

## 🎓 Kiến thức áp dụng

- ✅ Random number generation
- ✅ For/While loops
- ✅ If/Else conditions
- ✅ Input/Output
- ✅ Variables & Counters
- ✅ Error handling

---

**Deadline:** 1.5 giờ
**Start:** Sau Lesson 5 ✓
