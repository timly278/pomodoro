package main

import (
	"database/sql"
	"log"
	db "pomodoro/db/sqlc"
	"pomodoro/server"
	"pomodoro/util"

	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Println("cannot load config:", err)
		return
	}
	dataBase, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Println("cannot open data base.", err)
	}

	store := db.NewSQLStore(dataBase)
	server, err := server.NewServer(store, &config)
	if err != nil {
		log.Println("cannot create server", err)
	}

	server.Run(config.ServerAddress)

}
