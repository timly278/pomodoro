package jobservice

import db "pomodoro/db/sqlc"

type jobService struct {
	store db.Store
}

func NewJobService(store db.Store) *jobService {
	return &jobService{store: store}
}
