package main

import (
	"database/sql"
	"log"
	"pomodoro/api"
	db "pomodoro/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	SERVER_ADDRESS = "0.0.0.0:8080"
	DB_DRIVER      = "postgres"
	DB_SOURCE      = "postgresql://root:tulb@localhost:5432/pomodoro?sslmode=disable"
)

func main() {
	dataBase, err := sql.Open(DB_DRIVER, DB_SOURCE)
	if err != nil {
		log.Println("cannot open data base.", err)
	}

	store := db.NewSQLStore(dataBase)
	server, err := api.NewServer(store)
	if err != nil {
		log.Println("cannot create server", err)
	}

	server.Setup()
	err = server.Start(SERVER_ADDRESS)
	if err != nil {
		log.Println("can not start server.", err)
	}
}
