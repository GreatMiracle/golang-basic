package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Danh sách các website cần kiểm tra
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Tạo channel để nhận thông tin từ các Goroutines
	c := make(chan string)

	// Khởi chạy Goroutines cho từng link
	for _, link := range links {
		go checkLink(link, c)
	}

	// Nghe thông tin từ channel và kiểm tra lại các liên kết
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

// Hàm kiểm tra trạng thái của website
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- fmt.Sprintf("Error: %s might be down!", link)
		return
	}
	c <- fmt.Sprintf("Success: %s is up!", link)
}
