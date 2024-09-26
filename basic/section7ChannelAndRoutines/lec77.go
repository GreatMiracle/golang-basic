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

	for _, link := range links {
		//Gửi yêu cầu HTTP GET đến website và lưu trữ phản hồi vào biến resp. Nếu có lỗi, nó sẽ được lưu vào biến err.
		resp, err := http.Get(link)
		if err != nil {
			// Nếu có lỗi, in ra thông báo
			fmt.Printf("Error: %s might be down\n", link)
			continue
		}
		defer resp.Body.Close() // Đảm bảo đóng body sau khi hoàn thành
		// Nếu yêu cầu thành công
		if resp.StatusCode == 200 {
			fmt.Printf("Success: %s is up\n", link)
		} else {
			fmt.Printf("Error: %s returned status %d\n", link, resp.StatusCode)
		}
	}

}
