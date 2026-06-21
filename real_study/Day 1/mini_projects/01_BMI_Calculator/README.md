# Mini Project 1: BMI Calculator

## 📋 Mô tả dự án

Viết chương trình tính chỉ số khối cơ thể (BMI - Body Mass Index) và đưa ra nhận xét về tình trạng sức khỏe của người dùng.

---

## 🎯 Yêu cầu (Requirements)

### Tính năng chính

- [ ] Nhập cân nặng (kg) từ người dùng
- [ ] Nhập chiều cao (m) từ người dùng
- [ ] Tính BMI = cân nặng / (chiều cao²)
- [ ] Phân loại kết quả:
  - BMI < 18.5: Gầy (Underweight)
  - 18.5 ≤ BMI < 25: Bình thường (Normal)
  - 25 ≤ BMI < 30: Thừa cân (Overweight)
  - BMI ≥ 30: Béo phì (Obese)
- [ ] Hiển thị kết quả chi tiết

### Xử lý lỗi

- [ ] Kiểm tra input hợp lệ (không âm, > 0)
- [ ] Xử lý khi nhập text thay vì số
- [ ] Xử lý chia cho 0 (chiều cao = 0)

---

## 📝 Ví dụ chương trình

**Input:**

```
Nhập cân nặng (kg): 70
Nhập chiều cao (m): 1.75
```

**Output:**

```
=== KẾT QUẢ BMI ===
Cân nặng: 70.00 kg
Chiều cao: 1.75 m
BMI: 22.86
Phân loại: BÌNH THƯỜNG ✓
```

---

## 🔨 Hướng dẫn thực hiện

### Bước 1: Khởi tạo project

```bash
mkdir bmi_calculator
cd bmi_calculator
go mod init bmi_calculator
```

### Bước 2: Tạo main.go

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    var weight float64
    var height float64

    // TODO: Nhập input từ người dùng
    fmt.Print("Nhập cân nặng (kg): ")
    fmt.Scanln(&weight)

    fmt.Print("Nhập chiều cao (m): ")
    fmt.Scanln(&height)

    // TODO: Tính BMI

    // TODO: Phân loại

    // TODO: In kết quả
}
```

### Bước 3: Hàm tính BMI

```go
func calculateBMI(weight, height float64) (float64, error) {
    // TODO: Validate input
    // TODO: Calculate BMI
    // TODO: Return result
}
```

### Bước 4: Hàm phân loại

```go
func classifyBMI(bmi float64) string {
    switch {
    case bmi < 18.5:
        return "Gầy (Underweight)"
    case bmi < 25:
        return "Bình thường (Normal)"
    case bmi < 30:
        return "Thừa cân (Overweight)"
    default:
        return "Béo phì (Obese)"
    }
}
```

---

## 💡 Gợi ý thêm

1. **Input**: Dùng `fmt.Scanln()` để nhập
2. **Validation**: Kiểm tra weight > 0 và height > 0
3. **Error handling**: Return error nếu input không hợp lệ
4. **Formatting**: Dùng `fmt.Printf("%.2f\n", bmi)` để 2 chữ số thập phân
5. **Testing**: Test với nhiều input khác nhau

---

## 🧪 Test cases

| Weight | Height | BMI   | Classification |
| ------ | ------ | ----- | -------------- |
| 70     | 1.75   | 22.86 | Normal         |
| 50     | 1.60   | 19.53 | Normal         |
| 60     | 1.50   | 26.67 | Overweight     |
| 40     | 1.80   | 12.35 | Underweight    |
| 100    | 1.70   | 34.59 | Obese          |

---

## 🌟 Bonus (thêm điểm)

- [ ] Tính BMI cho nhiều người (loop)
- [ ] Lưu kết quả vào file
- [ ] Tính lượng cân nặng cần tăng/giảm để đạt BMI bình thường
- [ ] Gợi ý ăn uống dựa trên phân loại

---

## 📊 Đánh giá

| Tiêu chí       | Điểm |
| -------------- | ---- |
| Functionality  | 40%  |
| Code Quality   | 30%  |
| Error Handling | 20%  |
| Bonus Features | 10%  |

---

## 🎓 Kiến thức áp dụng

- ✅ Variables & Types
- ✅ Functions & Return values
- ✅ If/Else conditions
- ✅ Input/Output (fmt, Scanln)
- ✅ Error handling
- ✅ String formatting

---

**Deadline:** 1 giờ
**Start:** Sau Lesson 3 ✓
