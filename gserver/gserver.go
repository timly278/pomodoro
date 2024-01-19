package gserver

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net"
	db "pomodoro/db/sqlc"
	gapi "pomodoro/gapi/auth-handlers"
	"pomodoro/pb"
	"pomodoro/plogger"
	"pomodoro/security"
	"pomodoro/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	gomail "gopkg.in/mail.v2"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type gRPCServer struct {
	gapi.AuthHandlers
	//JobsHandlers...
}

// Construct components of GRPC server
func NewGrpcServerComponents() (
	store db.Store,
	tokenMaker security.TokenMaker,
	config *util.Config,
	dialer *gomail.Dialer,
	redisDb *redis.Client,
	logger *zap.Logger,
	err error,
) {
	config, err = util.LoadConfig(".")
	if err != nil {
		fmt.Println("cannot load config:", err)
		return
	}
	fmt.Println("Load configuration successfully")

	logger = plogger.New(config, "pomodoro")
	defer logger.Sync()

	sugar := logger.Sugar()

	dataBase, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		sugar.Fatalf("cannot open data base.", err)
		return
	}
	sugar.Infoln("Database is opened successfully")

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
	return
}

func NewServer(auth *gapi.AuthHandlers) *gRPCServer {
	return &gRPCServer{AuthHandlers: *auth}
}

func Run(config *util.Config, server *gRPCServer) {

	grpcServer := grpc.NewServer()
	pb.RegisterAuthHandlersServer(grpcServer, server)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("started gRPC at %v\n", config.GrpcServerAddress)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}
