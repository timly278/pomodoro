package jobservice

import (
	"context"
	"pomodoro/api/delivery"
	db "pomodoro/db/sqlc"
)

func (p *jobService) CreateNewType(ctx context.Context, userId int64, req *delivery.CreateNewTypeRequest) (db.Type, error) {

	return p.store.CreateNewType(ctx, db.CreateNewTypeParams{
		UserID:            userId,
		Name:              req.Name,
		Color:             req.Color,
		Goalperday:        req.Goalperday,
		Duration:          req.Duration,
		Shortbreak:        req.Shortbreak,
		Longbreak:         req.Longbreak,
		Longbreakinterval: req.Longbreakinterval,
		AutostartBreak:    req.AutostartBreak,
	})
}

func (p *jobService) GetTypes(ctx context.Context, userId int64) ([]db.Type, error) {
	return p.store.ListTypes(ctx, userId)
}

func (p *jobService) UpdateType(ctx context.Context, userId int64, typeId int64, req *delivery.CreateNewTypeRequest) (db.Type, error) {

	return p.store.UpdateTypeById(ctx, db.UpdateTypeByIdParams{
		ID:                typeId,
		UserID:            userId,
		Name:              req.Name,
		Color:             req.Color,
		Goalperday:        req.Goalperday,
		Duration:          req.Duration,
		Shortbreak:        req.Shortbreak,
		Longbreak:         req.Longbreak,
		Longbreakinterval: req.Longbreakinterval,
		AutostartBreak:    req.AutostartBreak,
	})

}
