package funcDeck

import (
	"log"
	"os"
)

// Hàm đọc file
func ReadDeckFromFile(filename string) string {
	// Đọc dữ liệu từ tệp bằng os.ReadFile
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return string(data)
}
