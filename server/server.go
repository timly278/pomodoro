package server

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	db "pomodoro/db/sqlc"
	_ "pomodoro/docs"
	"pomodoro/security"
	"pomodoro/util"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	gomail "gopkg.in/mail.v2"

	// gin-swagger middleware
	ginSwagger "github.com/swaggo/gin-swagger"

	// swagger embed files
	swaggerFiles "github.com/swaggo/files"
)

// type Server struct {
// 	store      db.Store
// 	tokenMaker security.TokenMaker
// 	config     *util.Config
// 	dialer     *gomail.Dialer
// 	redisdb    *redis.Client
// 	// logger
// }

func NewServer() (
	router *gin.Engine,
	store db.Store,
	tokenMaker security.TokenMaker,
	config *util.Config,
	dialer *gomail.Dialer,
	redisDb *redis.Client,
	err error,
) {

	config, err = util.LoadConfig(".")
	if err != nil {
		log.Println("cannot load config:", err)
		return
	}
	dataBase, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Println("cannot open data base.", err)
	}
	// Create Redis Client
	redisDb = redis.NewClient(&redis.Options{
		Addr:     config.RedisClientAddress,
		Password: config.RedisDbPassword,
		DB:       config.RedisDb,
	})

	// Settings for SMTP server
	dialer = gomail.NewDialer(config.AppSmtpHost, config.AppSmtpPort, config.AppEmail, config.AppPassword)
	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	store = db.NewSQLStore(dataBase)

	tokenMaker, err = security.NewJwtTokenMaker(config.TokenSymmetricKey)
	if err != nil {
		err = fmt.Errorf("cannot create token maker: %w", err)
		return
	}
	router = gin.Default()

	return
}

func Run(router *gin.Engine, config *util.Config) {

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(config.ServerAddress)
}
