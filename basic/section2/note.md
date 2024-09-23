Một số lệnh quan trọng trong Go CLI:
go run <file>: Biên dịch và chạy ngay lập tức tệp Go.
go build <file>: Biên dịch tệp Go mà không chạy.
go fmt: Tự động định dạng mã nguồn Go.
go install: Cài đặt các gói (packages) và file thực thi.
go get: Lấy các thư viện từ bên ngoài về để sử dụng trong dự án.
go test: Chạy các tệp kiểm thử (test files).



Trong Go, có hai loại package chính:
Executable package: Đây là loại package khi được biên dịch sẽ tạo ra một tệp thực thi có thể chạy.
Reusable package: Đây là loại package không được biên dịch thành tệp thực thi mà chứa các mã tái sử dụng, chẳng hạn như thư viện hoặc các hàm phụ trợ.
Khi bạn sử dụng tên package là main, bạn đang tạo ra một executable package. Khi biên dịch, chương trình sẽ tạo ra một tệp thực thi có thể chạy được.
Nếu sử dụng tên package khác, như package apple, chương trình sẽ không tạo ra tệp thực thi khi biên dịch, mà chỉ tạo ra một package chứa mã tái sử dụng.