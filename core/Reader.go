package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// Tạo một reader từ chuỗi
	r := strings.NewReader("Hello, Go!")

	// Tạo buffer chứa 8 byte
	buf := make([]byte, 8)

	// Đọc dữ liệu từ reader cho đến khi gặp io.EOF
	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
	}

}
