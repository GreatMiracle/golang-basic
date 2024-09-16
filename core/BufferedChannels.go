package main

import "fmt"

func main() {
	// Tạo một buffered channel có bộ đệm chứa được 2 giá trị
	ch := make(chan int, 2)

	// Gửi 2 giá trị vào channel mà không bị chặn
	ch <- 1
	//ch <- 2

	fmt.Println("Đã gửi 2 giá trị vào channel")

	// Gửi thêm một giá trị sẽ bị chặn vì channel đã đầy
	// Uncomment dòng dưới để thấy việc chặn khi gửi vào channel đầy
	// ch <- 3

	// Nhận giá trị từ channel
	fmt.Println(<-ch) // Lấy ra giá trị đầu tiên (1)
	//fmt.Println(<-ch) // Lấy ra giá trị thứ hai (2)

	// Sau khi nhận hết, có thể gửi thêm giá trị vào channel
	ch <- 3
	ch <- 4
	//ch <- 5
	fmt.Println(<-ch) // Lấy ra giá trị thứ ba (3)
}
