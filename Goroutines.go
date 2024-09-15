package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond) // Tạm dừng goroutine 100ms
	}
}

func main() {
	// Khởi chạy sayHello() trong một goroutine mới
	go sayHello()

	// Goroutine chính thực hiện công việc khác
	for i := 0; i < 5; i++ {
		fmt.Println("Main Goroutine")
		time.Sleep(150 * time.Millisecond)
	}
}
