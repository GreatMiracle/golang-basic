package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// Method sử dụng pointer receiver
func (v *Vertex) Scale(factor float64) {
	v.X = v.X * factor
	v.Y = v.Y * factor
}

// Method sử dụng value receiver
func (v Vertex) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v1 Vertex) ScalePointerValue(factor float64) {
	v1.X = v1.X * factor
	v1.Y = v1.Y * factor
}

func main() {
	v := Vertex{3, 4}

	// Gọi method với pointer receiver (thay đổi giá trị ban đầu)
	v.Scale(2)
	fmt.Println(v) // Kết quả: {6 8}

	// Gọi method với value receiver (không thay đổi giá trị ban đầu)
	v1 := Vertex{3, 4}
	v1.ScalePointerValue(2)
	fmt.Println(v1)

	fmt.Println(v.Abs()) // Kết quả: 100

}
