package main

import (
	"errors"
	"fmt"
)

// Hàm chia hai số nguyên, trả về kết quả và lỗi nếu có
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// Thử chia 10 cho 2, không có lỗi xảy ra
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Thử chia 10 cho 0, sẽ trả về lỗi
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
