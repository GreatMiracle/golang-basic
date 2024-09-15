package main

import "fmt"

func sendValues(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Gui ch <- i: %d\n", i)
		ch <- i
	}
	close(ch) // Đóng channel sau khi gửi xong
}

func main() {
	ch := make(chan int)

	// Khởi chạy goroutine để gửi giá trị
	go sendValues(ch)

	// Nhận giá trị từ channel bằng vòng lặp range
	for v := range ch {
		fmt.Println("Nhận giá trị:", v)
	}

	fmt.Println("Channel đã bị đóng, không còn giá trị nào nữa.")
}
