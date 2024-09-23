package main

import "fmt"

func main() {
	// Khai báo biến kiểu tường minh
    var card string = "Ace of Spades"
    fmt.Println(card) // In ra: Ace of Spades

    // Khai báo biến với cú pháp ngắn gọn
    card1 := "Five of Diamonds"
    fmt.Println(card1) // In ra: Five of Diamonds

    // Gán lại giá trị cho biến đã khai báo
    card = "Queen of Hearts"
    fmt.Println(card) // In ra: Queen of Hearts
}