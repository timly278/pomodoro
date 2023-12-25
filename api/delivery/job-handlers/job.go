package jobs

import (
	"pomodoro/api/delivery"
	"pomodoro/api/service"
)

type jobHandlers struct {
	jobService  service.JobsService
	userService service.UserService
}

var _ delivery.JobHandlers = (*jobHandlers)(nil)

func NewJobHandlers(job service.JobsService, user service.UserService) delivery.JobHandlers {
	return &jobHandlers{jobService: job, userService: user}
}
