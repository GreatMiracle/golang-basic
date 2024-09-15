package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine gửi dữ liệu vào channel ch1 sau 2 giây
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from ch1"
	}()

	// Goroutine gửi dữ liệu vào channel ch2 sau 1 giây
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()

	// Sử dụng select để chờ đợi trên cả 2 channel
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
