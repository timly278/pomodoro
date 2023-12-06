package pomodo

import db "pomodoro/db/sqlc"

type pomoHandlers struct {
	store db.Store
}

func NewPomoHandlers(store db.Store) *pomoHandlers {
	return &pomoHandlers{store: store}
}
