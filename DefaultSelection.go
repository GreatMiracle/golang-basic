package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case i := <-ch:
		// Không nhận được giá trị vì ch chưa có dữ liệu
		fmt.Println("Nhận giá trị từ channel:", i)
	default:
		// Default case sẽ chạy ngay vì channel ch không có dữ liệu
		fmt.Println("Không có dữ liệu nào sẵn sàng nhận từ channel")
	}
}
