package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Gửi yêu cầu HTTP đến một URL
	resp, err := http.Get("https://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close() // Đảm bảo đóng resp.Body khi không còn sử dụng

	// Sử dụng io.Copy để sao chép dữ liệu từ resp.Body đến os.Stdout
	bytesCopied, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Println("Error copying data:", err)
		return
	}

	fmt.Printf("Total bytes copied: %d\n", bytesCopied) // In ra số byte đã sao chép
}

/*Bạn có thể sử dụng io.Copy để thay thế cho việc sử dụng make và đọc dữ liệu từ io.Reader bằng phương thức Read.
	Đây là một số điểm chính để hiểu rõ hơn về cách hoạt động của io.Copy và sự thay thế này:

Sự thay thế giữa io.Copy và Read với make
Đơn giản hóa mã:

Khi bạn sử dụng io.Copy, bạn không cần phải tạo một byte slice bằng cách sử dụng make, rồi sau đó gọi hàm Read để đọc dữ liệu.
	io.Copy sẽ tự động xử lý việc đọc và ghi dữ liệu cho bạn.
io.Copy sẽ tiếp tục đọc từ io.Reader (như resp.Body) và ghi vào một io.Writer (như os.Stdout) cho đến khi không còn dữ liệu nào để đọc hoặc có lỗi xảy ra.
Tối ưu hóa:

io.Copy được tối ưu hóa để sao chép dữ liệu một cách hiệu quả, thường sử dụng các bộ đệm (buffer) bên trong để giảm số lần gọi hàm và tăng tốc độ sao chép.
Giao diện thống nhất:

io.Copy sử dụng các giao diện io.Reader và io.Writer, cho phép bạn làm việc với bất kỳ nguồn dữ liệu nào mà triển khai giao diện đó,
không chỉ là resp.Body hoặc os.Stdout*/
