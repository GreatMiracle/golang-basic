package funcDeck

import (
	"fmt"
	"os"
	"path/filepath"
)

// Hàm lưu bộ bài vào tệp
func (d Deck) SaveDeckToFile(fileName string) error {

	// Lấy đường dẫn thư mục từ tên file
	dir := filepath.Dir(fileName)

	// Kiểm tra nếu thư mục chưa tồn tại, thì tạo nó
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755) // Tạo thư mục với quyền đọc/ghi cho chủ sở hữu và quyền đọc cho người khác
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	data := []byte(fmt.Sprintf("%v", d))
	return os.WriteFile(fileName, data, 0777)
}
