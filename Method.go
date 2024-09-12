package main

import (
	"fmt"
)

// Định nghĩa kiểu mới dựa trên float64
type MyFloat float64

// Method cho MyFloat
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var f MyFloat = -7.5
	fmt.Println(f.Abs()) // Kết quả: 7.5
}
