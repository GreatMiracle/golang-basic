package main

import (
	"fmt"
	"io"
	"os"
)

// go run main.go ./golang-basic/basic/write-file/hand_poker_deck.txt
func main() {

	//// Kiểm tra nếu số lượng tham số là đủ
	//if len(os.Args) < 2 {
	//	fmt.Println("Vui lòng cung cấp tên tập tin.")
	//	return
	//}
	//filename := os.Args[1]

	filename := "/golang-basic/basic/write-file/hand_poker_deck.txt"
	// Lấy tên tập tin từ tham số dòng lệnh

	// Mở tập tin
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Lỗi khi mở tập tin:", err)
		return
	}
	defer file.Close() // Đảm bảo đóng tập tin sau khi đọc

	// Đọc nội dung tập tin
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Lỗi khi đọc tập tin:", err)
		return
	}

	// In nội dung ra màn hình
	fmt.Println(string(content))

}
