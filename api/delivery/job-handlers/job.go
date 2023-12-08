package pomodo

import (
	"pomodoro/api/service"
	jobservice "pomodoro/api/service/job-service"
	db "pomodoro/db/sqlc"
)

type jobHandlers struct {
	jobService service.JobsService
}

func NewPomoHandlers(store db.Store) *jobHandlers {
	job := jobservice.NewJobService(store)
	return &jobHandlers{jobService: job}
}
