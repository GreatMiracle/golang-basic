package main

import "fmt"

// Định nghĩa cấu trúc Triangle
type Triangle struct {
	base   float64
	height float64
}

// Định nghĩa phương thức getArea cho Triangle
func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

// Định nghĩa cấu trúc Square
type Square struct {
	side float64
}

// Định nghĩa phương thức getArea cho Square
func (s Square) getArea() float64 {
	return s.side * s.side
}

// Định nghĩa giao diện Shape
type Shape interface {
	getArea() float64
}

// Hàm printArea nhận vào một Shape và in diện tích của nó
func printArea(s Shape) {
	area := s.getArea()
	fmt.Printf("Diện tích là: %.2f\n", area)
}

func main() {
	triangle := Triangle{base: 10, height: 5}

	square := Square{side: 4}

	printArea(triangle)
	printArea(square)
}
