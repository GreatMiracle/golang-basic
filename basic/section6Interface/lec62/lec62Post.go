package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Tạo dữ liệu dưới dạng JSON
	jsonData := []byte(`{"name": "VietNamHN", "age": 26}`)

	// Thực hiện POST request
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonData))

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
