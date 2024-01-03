package jobs

import (
	"pomodoro/api/delivery"
	"pomodoro/api/service"

	"go.uber.org/zap"
)

type jobHandlers struct {
	jobService  service.JobsService
	userService service.UserService
	logger      *zap.Logger
}

// H is shortcut of map[string]any

var _ delivery.JobHandlers = (*jobHandlers)(nil)

func NewJobHandlers(job service.JobsService, user service.UserService, logger *zap.Logger) delivery.JobHandlers {
	return &jobHandlers{jobService: job, userService: user, logger: logger}
}
