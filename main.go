package main

import (
	"context"
	"log"
	"pomodoro/api/delivery"
	"pomodoro/api/delivery/auth-handlers"
	jobs "pomodoro/api/delivery/job-handlers"
	authservice "pomodoro/api/service/auth-service"
	jobservice "pomodoro/api/service/job-service"
	userservice "pomodoro/api/service/user-service"
	_ "pomodoro/docs"
	"pomodoro/server"
	mdw "pomodoro/shared/middleware"
	"time"

	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

//	@title			Pomodoro API
//	@version		1.0
//	@description	Pomodoro Application Api Server. This app helps people study and work at better productivity.
//	@description	Backend Language: Golang
//	@description	Database: PostgreSQL, Redis
//	@description	Framework: Gin, sqlc, uber-fx (for dependency injection), uber-zap (for logging files)
//	@description	Deployment: AWS EC2, Nginx

//	@contact.name	Ly Tu
//	@contact.url	https://github.com/timly278/pomodoro
//	@contact.email	timly278@gmail.com

// @host		18.140.71.34
// @BasePath	/api/v1
func main() {

	app := fxApp()
	// In a typical application, we could just use app.Run() here. Since we
	// don't want this example to run forever, we'll use the more-explicit Start
	// and Stop.
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}

}

func fxApp() *fx.App {
	return fx.New(
		fx.Provide(
			server.NewServer,
			jobservice.NewJobService,
			authservice.NewAuthService,
			userservice.NewUserService,
			jobs.NewJobHandlers,
			auth.NewAuthHandlers,
			mdw.New,
		),
		fx.Invoke(
			delivery.MapAuthRoutes,
			delivery.MapJobsRoutes,
			server.Run,
		),
	)

}
