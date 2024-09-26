package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Tạo yêu cầu DELETE
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/posts/1", nil)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Gửi yêu cầu
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// In trạng thái phản hồi từ server
	fmt.Println("Response status:", resp.Status)

}
