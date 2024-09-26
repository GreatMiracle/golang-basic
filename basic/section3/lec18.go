package main

import (
	"cards/section3/funcDeck"
	"fmt"
	"log"
)

func main() {
	// Tạo mới một bộ bài
	cards := funcDeck.NewDeck()
	//  Gọi hàm print để in ra các lá bài
	//cards.PrintDeck()

	fmt.Println("Shuffling the deck...")
	cards.ShuffleDeck()

	// Gọi hàm deal chia bài
	fmt.Println("Deal the deck...")
	hand, _ := cards.DealDeck(5)
	fmt.Println(hand)

	// In lại bộ bài đã được trộn
	//    cards.print()

	fileName := "./golang-basic/basic/write-file/hand_poker_deck.txt"
	err := hand.SaveDeckToFile(fileName)
	if err != nil {
		log.Fatalf("Error saving to file: %v", err)
	}

	// Tải bộ bài từ tệp
	loadedDeck := funcDeck.ReadDeckFromFile(fileName)
	fmt.Println("Loaded Deck:", loadedDeck)

}
