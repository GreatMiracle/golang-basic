package main

import "fmt"

// Định nghĩa kiểu Vertex
type Vertex struct {
	X, Y int
}

// Triển khai phương thức String() cho kiểu Vertex
func (v Vertex) String() string {
	return fmt.Sprintf("Vertex(X: %d, Y: %d)", v.X, v.Y)
}

func main() {
	v := Vertex{3, 4}
	// In đối tượng v, sẽ tự động gọi phương thức String()
	fmt.Println(v)
}
