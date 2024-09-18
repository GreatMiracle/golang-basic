// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

/*Giải thích:

Đây là một interface tên là DBTX dùng để định nghĩa những phương thức cơ bản mà bất kỳ đối tượng nào đại diện cho kết nối database cần phải có.
Exec: Thực thi một câu lệnh SQL (chẳng hạn như INSERT, UPDATE, hoặc DELETE). Trả về pgconn.CommandTag chứa thông tin về kết quả của câu lệnh, và lỗi nếu có.
Query: Dùng để truy vấn nhiều kết quả từ database (ví dụ như SELECT). Trả về pgx.Rows, đại diện cho các dòng dữ liệu kết quả.
QueryRow: Dùng để truy vấn một dòng kết quả duy nhất (ví dụ như SELECT ... LIMIT 1). Trả về pgx.Row, đại diện cho một dòng dữ liệu kết quả.*/

type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}


/*Hàm này tạo mới một instance của Queries với đối tượng database được truyền vào (db).
Hàm này giúp khởi tạo các truy vấn SQL trong đối tượng Queries.*/
func New(db DBTX) *Queries {
	return &Queries{db: db}
}



//Queries: Đây là cấu trúc chứa các phương thức truy vấn SQL. db là kết nối database mà Queries sẽ sử dụng để thực hiện các truy vấn.
type Queries struct {
	db DBTX
}



//WithTx: Phương thức này cho phép tạo ra một đối tượng Queries mới, với giao dịch đang diễn ra (tx).
//Khi thực hiện trong một transaction, bạn có thể sử dụng giao dịch này để thực hiện các câu lệnh SQL mà không cần khởi tạo lại Queries.
func (q *Queries) WithTx(tx pgx.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}
