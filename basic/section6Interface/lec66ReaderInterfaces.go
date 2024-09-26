package main

import (
	"fmt"
	"io"
	"strings"
)

// Hàm printData nhận vào một `io.Reader` và đọc dữ liệu từ nó
func printData(r io.Reader) {
	buf := make([]byte, 8) // Tạo một buffer với kích thước 8 byte
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break // Kết thúc nếu đọc hết dữ liệu
		}
		if err != nil {
			fmt.Println("Error reading:", err)
			break
		}
		// In ra số byte đã đọc và dữ liệu
		fmt.Print(string(buf[:n])) // In ra nội dung đã đọc
	}
}

func main() {
	// Chuỗi input
	data := "Hello, Golang Reader!."

	// Sử dụng strings.NewReader để tạo một đối tượng triển khai Reader interface
	reader := strings.NewReader(data)

	// Gọi hàm printData với đối tượng reader
	printData(reader)
}
