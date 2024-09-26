package funcDeck

import (
	"math/rand"
	"time"
)

// Hàm với receiver trộn bộ bài
func (d Deck) ShuffleDeck() {

	// Thiết lập seed cho số ngẫu nhiên
	//rand.Seed(time.Now().UnixNano())
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	for i := range d {
		// Tạo một số ngẫu nhiên từ 0 đến chiều dài của bộ bài
		newPosition := r.Intn(len(d))

		// Hoán đổi vị trí của phần tử
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
