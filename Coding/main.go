package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// Any formats any value as a string.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
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
	fmt.Println(Any(42))   // "42"
	fmt.Println(Any(true)) // "\"Hello, World!\""
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any([]int64{x}))
	// "1"
	// "1"
	// "[]int64 0x8202b87b0"
	fmt.Println(Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"
}
