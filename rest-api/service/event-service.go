package service

import (
	"database/sql"
	"fmt"
	"rest-api/db"
	"rest-api/models"
)

// Tạo một slice lưu trữ các sự kiện
var events []models.Event

// Phương thức Save để lưu sự kiện vào slice
func SaveEvent(e *models.Event) error {
	// Câu truy vấn để chèn dữ liệu vào bảng "events"
	query := `
		INSERT INTO events (name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
		RETURNING id
	`

	// Chuẩn bị câu truy vấn
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)

	// Lấy ID của sự kiện vừa được chèn vào bằng cách quét (scan) giá trị trả về
	var lastInsertID int64
	err = stmt.QueryRow(e.Name, e.Description, e.Location, e.DateTime, e.UserID).Scan(&lastInsertID)
	if err != nil {
		return fmt.Errorf("failed to execute statement and retrieve last insert ID: %w", err)
	}

	// Gán ID cho struct Event
	e.ID = int(lastInsertID)
	return nil
}

func GetAllEvents() ([]models.Event, error) {

	query := "SELECT id, name, description, location, dateTime, user_id FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("lỗi khi thực hiện truy vấn: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	var events []models.Event
	for rows.Next() {

		var event models.Event

		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, fmt.Errorf("lỗi khi quét dữ liệu: %w", err)
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("lỗi khi duyệt qua các dòng: %w", err)
	}

	return events, err
}
