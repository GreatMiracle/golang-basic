package main

import (
	"fmt"
	"io"
	"net/http"
)

// MyWriter là một loại cài đặt giao diện Writer
type MyWriter struct{}

// Write thực hiện phương thức của giao diện Writer
func (mw MyWriter) Write(p []byte) (n int, err error) {
	// Ghi dữ liệu ra màn hình
	return fmt.Print(string(p))
}

func main() {

	// Tạo một MyWriter
	myWriter := MyWriter{}

	// Gửi yêu cầu GET đến một URL
	resp, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Sao chép dữ liệu từ resp.Body (Reader) đến myWriter (Writer)
	_, err = io.Copy(myWriter, resp.Body) // myWriter cài đặt Writer
	if err != nil {
		panic(err)
	}

}
