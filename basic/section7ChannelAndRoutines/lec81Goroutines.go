package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // Thông báo công việc đã hoàn thành khi hàm kết thúc

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond) // Chờ 0.5 giây giữa mỗi lần in số
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done() // Thông báo công việc đã hoàn thành khi hàm kết thúc

	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(500 * time.Millisecond) // Chờ 0.5 giây giữa mỗi lần in ký tự
	}
}

func main() {

	var wg sync.WaitGroup

	// Thêm 2 công việc vào WaitGroup: một cho printNumbers và một cho printLetters
	wg.Add(2)

	// Khởi chạy goroutine cho hàm printNumbers
	go printNumbers(&wg)

	// Khởi chạy goroutine cho hàm printLetters
	go printLetters(&wg)

	// Chờ tất cả các goroutines hoàn thành
	wg.Wait()

	// Đợi 3 giây để các goroutine hoàn thành
	time.Sleep(1 * time.Second)
	fmt.Println("Main goroutine finished.")
}
