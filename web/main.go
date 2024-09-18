package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"simpleBank/api"
	db "simpleBank/db/sqlc"

	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://simple_bank:123456a@@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
