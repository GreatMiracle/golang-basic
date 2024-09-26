package cards

import (
	"cards/section3/funcDeck"
	"os"
	"strings"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := funcDeck.NewDeck()

	expectedDeckSize := 52
	if len(deck) != expectedDeckSize {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckSize, len(deck))
	}

	expectedFirstCard := "A ♤\n"
	if deck[0] != expectedFirstCard {
		t.Errorf("Expected first card to be %v, but got %v", expectedFirstCard, deck[0])
	}

	expectedLastCard := "K ♧\n"
	if deck[len(deck)-1] != expectedLastCard {
		t.Errorf("Expected last card to be %v, but got %v", expectedLastCard, deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Đặt tên tệp tin tạm thời
	testFileName := "_decktesting"

	// Xóa tệp tin cũ nếu có
	os.Remove(testFileName)

	// Tạo một deck mới
	deck := funcDeck.Deck{"A ♤",
		"2 ♦",
		"3 ♧",
		"4 ♥",
		"5 ♥",
	}

	// Lưu deck vào tệp tin
	err := deck.SaveDeckToFile(testFileName)
	if err != nil {
		t.Fatalf("Không thể lưu deck vào tệp tin: %v", err)
	}

	// Tải lại deck từ tệp tin
	loadedDeck := funcDeck.ReadDeckFromFile(testFileName)
	_, spaceCount := splitAndCountSpaces(loadedDeck)
	// Xác thực rằng số lượng thẻ trong deck là như mong đợi
	if (spaceCount / 2) != 4 {
		t.Errorf("Kỳ vọng 4 thẻ trong deck, nhưng có %v", len(loadedDeck))
	}

	// Dọn dẹp tệp tin tạm thời sau khi kiểm thử
	os.Remove(testFileName)
}

// Hàm cắt chuỗi tại khoảng trắng và đếm số khoảng trắng
func splitAndCountSpaces(input string) ([]string, int) {
	// Cắt chuỗi tại khoảng trắng
	words := strings.Fields(input)
	// Đếm số khoảng trắng
	spaceCount := strings.Count(input, " ")
	return words, spaceCount
}
