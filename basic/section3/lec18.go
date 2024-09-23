package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Định nghĩa kiểu 'Deck' là một slice của string
type Deck []string

// Hàm tạo mới một bộ bài
func newDeck() Deck {
    cards := Deck{}

    // Sử dụng hai slice để kết hợp tạo ra bộ bài
    cardSuits := []string{"♤", "♥", "♦", "♧"}
    cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

    // Tạo từng lá bài
    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value+" of "+suit)
        }
    }

    return cards
}


// Hàm với receiver để in ra bộ bài
func (d Deck) print() {
    for _, card := range d {
        fmt.Println(card)
    }
}

// Hàm với receiver trộn bộ bài
func (d Deck) shuffle() {
    rand.Seed(time.Now().UnixNano())

    for i := range d {
        newPosition := rand.Intn(len(d))
        d[i], d[newPosition] = d[newPosition], d[i]
    }
}

func(d Deck)deal (handSize int) (Deck, Deck){
	hand := d[:handSize]
	remainingDeck := d[handSize:]

	return hand, remainingDeck
}

func main() {
	 // Tạo mới một bộ bài
	 cards := newDeck()
    //  Gọi hàm print để in ra các lá bài
	//  cards.print()

	 
	 fmt.Println("Shuffling the deck...")
	 cards.shuffle()

	 // Gọi hàm deal chia bài
	 fmt.Println("Deal the deck...")
	 hand, _:=cards.deal(5)
	 fmt.Println(hand)


	   // In lại bộ bài đã được trộn
	//    cards.print()
 

}