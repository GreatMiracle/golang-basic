package db

import (
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://simple_bank:123456a@@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

//func TestMain(m *testing.M) {
//
//	// Mở kết nối DB
//	conn, err := sql.Open(dbDriver, dbSource)
//	if err != nil {
//		log.Fatal("cannot connect to db:", err)
//	}
//	// Khởi tạo context
//	testQueries = New(conn)
//	os.Exit(m.Run())
//}

func TestMain(m *testing.M) {
	// Mở kết nối DB bằng pgx
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close(context.Background())

	// Truyền kết nối vào trong Queries
	testQueries = New(conn)
	os.Exit(m.Run())
}
