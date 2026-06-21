package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	fmt.Println("=== Chào mừng đến với trò chơi đoán số! ===")
	fmt.Println("Tôi đã chọn một số ngẫu nhiên từ 1 đến số bạn thích. Bạn có thể đoán nó không?")
	var max int
	fmt.Print("Nhập số lớn nhất cho phạm vi đoán (ví dụ: 100): ")
	fmt.Scanln(&max)
    rand.Seed(time.Now().UnixNano())
    secretNumber := rand.Intn(max) + 1

    attempts := 0
    guessed := false

    for !guessed {
        var guess int

        fmt.Print("Nhập số của bạn (1-" + fmt.Sprint(max) + "): ")
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