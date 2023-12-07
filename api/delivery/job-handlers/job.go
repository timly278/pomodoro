package pomodo

import (
	"pomodoro/api/service"
	pomoservice "pomodoro/api/service/job-service"
	db "pomodoro/db/sqlc"
)

type jobHandlers struct {
	pomoService service.PomodoService
}

func NewPomoHandlers(store db.Store) *jobHandlers {
	pomo := pomoservice.NewPomoService(store)
	return &jobHandlers{pomoService: pomo}
}
