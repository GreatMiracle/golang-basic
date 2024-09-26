package main

import "fmt"

// Định nghĩa một struct Card với các thuộc tính Suit và Value
type Card struct {
	Suit  string
	Value string
}

// Hàm in thông tin của lá bài
func (c Card) Print() {
	fmt.Printf("Card: %s of %s\n", c.Value, c.Suit)
}

//string{"♤", "♥", "♦", "♧"}
//string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

func main() {
	// Tạo một lá bài mới
	card1 := Card{Suit: "♤", Value: "A"}
	card2 := Card{Suit: "♦", Value: "2"}

	// In thông tin của các lá bài
	card1.Print() // Output: Card: Ace of Hearts
	card2.Print() // Output: Card: Two of Spades
}
