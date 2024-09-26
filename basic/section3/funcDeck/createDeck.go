package funcDeck

// Định nghĩa kiểu 'Deck' là một slice của string
type Deck []string

// Hàm tạo mới một bộ bài
func NewDeck() Deck {
	cards := Deck{}

	// Sử dụng hai slice để kết hợp tạo ra bộ bài
	cardSuits := []string{"♤", "♥", "♦", "♧"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	// Tạo từng lá bài
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" "+suit+"\n")
		}
	}

	return cards
}
