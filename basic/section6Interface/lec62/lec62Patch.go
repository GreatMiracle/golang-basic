package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// URL của API bạn muốn gửi yêu cầu PATCH
	url := "https://example.com/api/resource/123"

	// Dữ liệu JSON cần gửi trong body
	data := map[string]interface{}{
		"username": "newusername",
		"email":    "newemail@example.com",
	}

	// Mã hóa dữ liệu thành JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Tạo yêu cầu PATCH
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Thiết lập Content-Type cho body là JSON
	req.Header.Set("Content-Type", "application/json")

	// Thực hiện yêu cầu bằng client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Xử lý phản hồi từ server
	fmt.Println("Response Status:", resp.Status)
}
