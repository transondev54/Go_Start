package main

import (
	"errors"
	"fmt"
)
func divide(x, y float64)(float64, error) {
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	}

func getCoordinates2() (float64, float64) {
    return 3.5, 4.5
}

// Dùng named returns
func getCoordinates() (x, y float64) {
    x = 3.5
    y = 4.5
    return
}
func main() {
    // // Integers
    // age := 25
    // count := int64(1000000)

    // // Floats
    // pi := math.Pi
    // height := 1.75

    // // String
    // name := "Ngọc"

    // // Boolean
    // isStudent := true

    // // Print
    // fmt.Println("=== Variables Demo ===")
    // fmt.Printf("Name: %s\n", name)
    // fmt.Printf("Age: %d\n", age)
    // fmt.Printf("Height: %.2f m\n", height)
    // fmt.Printf("Is Student: %v\n", isStudent)
    // fmt.Printf("Pi: %.9f\n", pi)
    // fmt.Printf("Count: %d\n" , count)

// 	arr := [3]string{"a", "b", "c"}
// 	for i, v := range arr {
//     fmt.Printf("%d: %s\n", i, v)
// }


	// Gọi



// Gọi
x, y := getCoordinates()
fmt.Printf("X: %.1f, Y: %.1f\n", x, y)

}