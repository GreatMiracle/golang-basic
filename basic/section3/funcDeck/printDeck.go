package funcDeck

import "fmt"

// Hàm với receiver để in ra bộ bài
func (d Deck) PrintDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}
