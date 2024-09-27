package db

import (
	"database/sql"

	//_ "github.com/mattn/go-sqlite3"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error

	//sử dụng hàm sql.Open() để mở kết nối đến cơ sở dữ liệu.
	//tham số đầu tiên là tên driver ("sqlite3") và tham số thứ hai là đường dẫn đến file cơ sở dữ liệu.
	//DB, err = sql.Open("sqlite3", "api.db")
	DB, err = sql.Open("sqlite", "./api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	//giới hạn số lượng kết nối mở đồng thời
	DB.SetMaxOpenConns(10)

	//duy trì một số kết nối mở sẵn
	DB.SetMaxIdleConns(5)

	// Kiểm tra kết nối
	if err := DB.Ping(); err != nil {
		panic("Failed to ping database.")
	}

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table.")
	}

}
