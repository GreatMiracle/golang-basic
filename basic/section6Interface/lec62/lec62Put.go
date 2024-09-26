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
	jsonData := []byte(`{"name": "VietNamHN2", "age": 26}`)

	// Tạo yêu cầu PUT
	req, err := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//Khi sử dụng http.NewRequest thì phải tự quản lý header cho yêu cầu.
	//Cụ thể, nếu bạn không thiết lập Content-Type, server sẽ không biết dữ liệu bạn gửi thuộc loại gì (JSON, XML, hoặc một loại khác).
	req.Header.Set("Content-Type", "application/json")

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
