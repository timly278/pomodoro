package jobservice

import (
	"pomodoro/api/service"
	db "pomodoro/db/sqlc"
)

type jobService struct {
	store db.Store
}

func NewJobService(store db.Store) service.JobsService {
	return &jobService{store: store}
}
