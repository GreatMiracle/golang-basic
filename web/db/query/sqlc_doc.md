Dưới đây là tất cả các cú pháp thông dụng mà bạn có thể sử dụng trong file .sql khi làm việc với sqlc để tự động sinh mã Go từ các truy vấn SQL:

1. -- name: query_name :one
   Mục đích: Truy vấn này sẽ trả về đúng một dòng kết quả. Trong Go, hàm tương ứng sẽ trả về một struct.
   Ví dụ:
   sql
   Copy code
   -- name: GetAccountByID :one
   SELECT id, owner, balance, currency, created_at
   FROM accounts
   WHERE id = $1;
   Hàm Go sinh ra sẽ trả về một đối tượng Account.
2. -- name: query_name :many
   Mục đích: Truy vấn này trả về nhiều dòng kết quả. Trong Go, hàm sẽ trả về một slice của các struct.
   Ví dụ:
   sql
   Copy code
   -- name: GetListAccount :many
   SELECT id, owner, balance, currency, created_at
   FROM accounts
   LIMIT $1 OFFSET $2;
   Hàm Go sẽ trả về một slice []Account.
3. -- name: query_name :exec
   Mục đích: Sử dụng cho các truy vấn chỉ thực thi mà không cần trả về kết quả, như UPDATE, DELETE, INSERT.
   Ví dụ:
   sql
   Copy code
   -- name: UpdateAccountBalance :exec
   UPDATE accounts
   SET balance = $2
   WHERE id = $1;
   Hàm Go sẽ chỉ thực thi truy vấn, không trả về dữ liệu.
4. -- name: query_name :execrows
   Mục đích: Tương tự như :exec, nhưng hàm Go sinh ra sẽ trả về số lượng dòng bị ảnh hưởng sau khi thực thi (affected rows).
   Ví dụ:
   sql
   Copy code
   -- name: DeleteAccount :execrows
   DELETE FROM accounts
   WHERE id = $1;
   Hàm Go sẽ trả về số lượng dòng bị xóa.
5. -- name: query_name :scalar
   Mục đích: Truy vấn này trả về một giá trị đơn lẻ (ví dụ: tổng, trung bình, đếm). Trong Go, hàm sẽ trả về một giá trị đơn.
   Ví dụ:
   sql
   Copy code
   -- name: CountAccounts :scalar
   SELECT COUNT(*)
   FROM accounts;
   Hàm Go sẽ trả về giá trị int hoặc int64 tùy theo kết quả truy vấn.
6. -- name: query_name :batchexec
   Mục đích: Dùng cho các truy vấn batch (INSERT, UPDATE, DELETE nhiều dòng cùng lúc). Hàm Go sẽ thực thi truy vấn với nhiều giá trị đầu vào.
   Ví dụ:
   sql
   Copy code
   -- name: BatchInsertAccounts :batchexec
   INSERT INTO accounts (owner, balance, currency)
   VALUES ($1, $2, $3);
   Hàm Go sẽ cho phép truyền nhiều bộ giá trị để thực hiện truy vấn hàng loạt.
   Các cú pháp đặc biệt (overrides):
   sqlc hỗ trợ thêm cú pháp để ghi đè các kiểu dữ liệu trả về hoặc kiểu dữ liệu của các tham số đầu vào.

Ví dụ:
yaml
Copy code
overrides:
- db_type: "timestamptz"
  go_type: "time.Time"
- db_type: "uuid"
  go_type: "github.com/google/uuid.UUID"
  db_type: Định nghĩa kiểu dữ liệu trong SQL.
  go_type: Định nghĩa kiểu dữ liệu tương ứng trong Go.
  Kết luận
  -- name: <query_name> :one cho một dòng dữ liệu.
  -- name: <query_name> :many cho nhiều dòng dữ liệu.
  -- name: <query_name> :exec cho truy vấn thực thi không trả về kết quả.
  -- name: <query_name> :execrows cho truy vấn thực thi trả về số dòng bị ảnh hưởng.
  -- name: <query_name> :scalar cho truy vấn trả về giá trị đơn lẻ.
  -- name: <query_name> :batchexec cho thực thi batch nhiều truy vấn.