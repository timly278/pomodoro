package main

import (
	"crypto/tls"
	"database/sql"
	"log"
	db "pomodoro/db/sqlc"
	"pomodoro/server"
	"pomodoro/util"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	gomail "gopkg.in/mail.v2"
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
	// Create Redis Client
	redisDb := redis.NewClient(&redis.Options{
		Addr:     config.RedisClientAddress,
		Password: config.RedisDbPassword,
		DB:       config.RedisDb,
	})

	// Settings for SMTP server
	dialer := gomail.NewDialer(config.AppSmtpHost, config.AppSmtpPort, config.AppEmail, config.AppPassword)
	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	store := db.NewSQLStore(dataBase)
	server, err := server.NewServer(store, &config, dialer, redisDb)
	if err != nil {
		log.Println("cannot create server", err)
	}

	server.Run(config.ServerAddress)

}
