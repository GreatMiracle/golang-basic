package main

import "fmt"

// Hàm adder trả về một closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	pos := adder()
	fmt.Println(pos(1)) // 1
	fmt.Println(pos(2)) // 3
	fmt.Println(pos(3)) // 6

	neg := adder()
	fmt.Println(neg(-1)) // -1
	fmt.Println(neg(-2)) // -3
	fmt.Println(neg(-3)) // -6

}
