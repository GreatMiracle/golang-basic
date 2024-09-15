package main

import "fmt"

func Index[T comparable](s []T, x T) int {
	for i2, v := range s {
		if v == x {
			return i2
		}
	}
	return -1
}

func main() {
	ints := []int{10, 20, 30, 40}
	fmt.Println(Index(ints, 30))

	strs := []string{"apple", "banana", "cherry"}
	fmt.Println(Index(strs, "banana")) // Kết quả: 1

	// Nếu không tìm thấy
	fmt.Println(Index(strs, "orange")) // Kết quả: -1
}
