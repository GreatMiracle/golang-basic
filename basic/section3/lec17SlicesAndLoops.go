package main

import "fmt"

func main() {
	// Tạo một slice của các lá bài (string)
	cards := []string{"Ace of Diamonds", "Five of Diamonds"}

	// In ra các lá bài ban đầu
	fmt.Println("Initial cards:", cards)

	// Thêm một lá bài mới vào slice bằng append
	cards = append(cards, "Six of Spades")

	// In ra slice sau khi thêm
	fmt.Println("After appending:", cards)

	// Lặp qua từng lá bài trong slice và in ra
	for i, card := range cards {
		fmt.Printf("Card %d: %s\n", i, card)
	}
}
