package jobservice

import db "pomodoro/db/sqlc"

type jobService struct {
	store db.Store
}

func NewPomoService(store db.Store) *jobService {
	return &jobService{store: store}
}
