package jobservice

import (
	"context"
	"database/sql"
	"pomodoro/api/delivery"
	db "pomodoro/db/sqlc"
)

func (p *jobService) CreatePomodoro(ctx context.Context, userId int64, req *delivery.CreatePomodoroRequest) (*db.Pomodoro, error) {

	pomodoro, err := p.store.CreatePomodoro(ctx, db.CreatePomodoroParams{
		UserID: userId,
		TypeID: req.TypeID,
		TaskID: sql.NullInt64{
			Int64: req.TaskID,
			Valid: bool(req.TaskID != 0),
		},
		FocusDegree: req.FocusDegree,
	})

	return &pomodoro, err
}

func (p *jobService) GetPomodorosByDates(ctx context.Context, userId int64, req *delivery.GetPomodorosRequest) ([]db.GetPomodorosRow, error) {
	queryParams := db.GetPomodorosParams{
		UserID:   userId,
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
		FromDate: req.FromDate,
		ToDate:   req.ToDate,
	}

	return p.store.GetPomodoros(ctx, queryParams)
}

func (p *jobService) GetMinutesFocused(ctx context.Context, userId int64, req *delivery.GetStatisticRequest) (int64, error) {
	queryParams := db.GetMinutesFocusedParams{
		UserID:   userId,
		FromDate: req.FromDate,
		ToDate:   req.ToDate,
	}

	return p.store.GetMinutesFocused(ctx, queryParams)
}
func (p *jobService) GetDaysAccessed(ctx context.Context, userId int64, req *delivery.GetStatisticRequest) (int64, error) {
	queryParams := db.GetDaysAccessedParams{
		UserID:   userId,
		FromDate: req.FromDate,
		ToDate:   req.ToDate,
	}

	return p.store.GetDaysAccessed(ctx, queryParams)
}
