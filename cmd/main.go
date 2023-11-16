package main

import (
	"database/sql"
	"log"
	"pomodoro/api"
	db "pomodoro/db/sqlc"
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
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Println("cannot create server", err)
	}

	server.Setup()
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Println("can not start server.", err)
	}
}
