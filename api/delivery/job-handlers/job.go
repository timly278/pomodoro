package jobs

import (
	"pomodoro/api/delivery"
	"pomodoro/api/service"
	jobservice "pomodoro/api/service/job-service"
	userservice "pomodoro/api/service/user-service"
	db "pomodoro/db/sqlc"
)

type jobHandlers struct {
	jobService  service.JobsService
	userService service.UserService
}

var _ delivery.JobHandlers = (*jobHandlers)(nil)

func NewPomoHandlers(store db.Store) *jobHandlers {
	job := jobservice.NewJobService(store)
	user := userservice.NewUserService(store)
	return &jobHandlers{jobService: job, userService: user}
}
