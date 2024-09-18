-- name: CreateAccount :one
INSERT INTO accounts (owner,
                      balance,
                      currency)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetAccount :one
SELECT id, owner, balance, currency, created_at FROM accounts
WHERE id = $1
;

-- name: ListAccounts :many
SELECT *
FROM accounts
WHERE owner = $1
ORDER BY id LIMIT $2
OFFSET $3;

-- name: UpdateAccount :one
UPDATE accounts
SET balance =$2
WHERE id = $1
RETURNING *
;

-- name: UpdateAccountNotReturning :exec
UPDATE accounts
SET owner =$2
WHERE id = $1
;

-- name: AddAccountBalance :one
UPDATE accounts
SET balance = balance + sqlc.arg(amount)
WHERE id = sqlc.arg(id) RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
;

-- Dòng -- name: GetListAccount :many trong câu lệnh SQL đóng vai trò rất quan trọng khi sử dụng sqlc để tự động sinh mã (code) trong Go.
-- Nó giúp xác định tên và kiểu trả về của truy vấn SQL trong quá trình tạo code Go từ SQL file. Hãy phân tích kỹ hơn:
--
-- Ý nghĩa của -- name: GetListAccount :many
-- -- name:: Đây là cú pháp của sqlc để đặt tên cho một truy vấn SQL. Trong ví dụ này, tên truy vấn là GetListAccount.
-- :many: Định nghĩa rằng truy vấn này trả về nhiều kết quả (nhiều dòng).
-- Điều này có nghĩa là hàm tương ứng trong Go sẽ trả về một slice (mảng) của các đối tượng Account (hoặc struct tương ứng).

