package main

import (
	"fmt"
	"math"
	"os"
)
func calculateBMI(weight, height float64) (float64, string) {
    // TODO: Validate input
    if weight <= 0 || height <= 0 {
        return 0, "Lỗi: Cân nặng và chiều cao phải lớn hơn 0"
    }
    // TODO: Calculate BMI
    bmi := weight / math.Pow(height, 2)
    // TODO: Return result
    return bmi, ""
}
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

func saveResult(weight, height, bmi float64, category string) error {
	content := fmt.Sprintf("=== KẾT QUẢ BMI ===\nCân nặng: %.2f kg\nChiều cao: %.2f m\nBMI: %.2f\nPhân loại: %s\n", weight, height, bmi, category)
	err := os.WriteFile("bmi_result.txt", []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("lỗi ghi file: %v", err)
	}
	return nil
}
func main() {
    var weight float64
    var height float64

    // TODO: Nhập input từ người dùng
    fmt.Print("Nhập cân nặng (kg): ")
    fmt.Scanln(&weight)

    fmt.Print("Nhập chiều cao (mét, VD: 1.65): ")
    fmt.Scanln(&height)

    // TODO: Tính BMI
    bmi, err1 := calculateBMI(weight, height)
    if err1 != "" {
        fmt.Println("Lỗi:", err1)
        return
    }

    // TODO: Phân loại
	category := classifyBMI(bmi)
    // TODO: In kết quả
    fmt.Println("\n=== KẾT QUẢ BMI ===")
    fmt.Printf("Cân nặng: %.2f kg\n", weight)
    fmt.Printf("Chiều cao: %.2f m\n", height)
    fmt.Printf("BMI: %.2f\n", bmi)
    fmt.Printf("Phân loại: %s\n", category)

    // TODO: Lưu kết quả vào file
    err := saveResult(weight, height, bmi, category)
    if err != nil {
        fmt.Println("Lỗi:", err)
        return
    }
    fmt.Println("\n✓ Kết quả đã được lưu vào file 'bmi_result.txt'")
}