package main

import "fmt"

func main() {
	// Ví dụ với slice
	numbers := []int{10, 20, 30}

	// Lặp qua slice, lấy cả chỉ mục và giá trị
	fmt.Println("Lặp qua slice với chỉ mục và giá trị:")
	for i, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}

	// Lặp qua slice, bỏ qua chỉ mục
	fmt.Println("\nLặp qua slice, bỏ qua chỉ mục:")
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}

	// Lặp qua slice, chỉ lấy chỉ mục
	fmt.Println("\nLặp qua slice, chỉ lấy chỉ mục:")
	for i := range numbers {
		fmt.Printf("Index: %d\n", i)
	}
}
