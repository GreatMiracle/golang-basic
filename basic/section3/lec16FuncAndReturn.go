package main

import "fmt"

func newCardStr() string {
    return "Five of Diamonds" // Trả về chuỗi "Five of Diamonds"
}

func newCardInt() int {
    return 12 // Trả về số nguyên 12
}

func main() {
    // Khai báo biến card và gán giá trị bằng kết quả trả về từ hàm newCard
    card := newCardStr()
    fmt.Println(card) // In ra màn hình: Five of Diamonds

	card1 := newCardInt()
    fmt.Println(card1) // In ra màn hình: Five of Diamonds
}