package jobservice

import (
	"context"
	"database/sql"
	"pomodoro/api/delivery"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"
)

type jobService struct {
	store db.Store
}

func NewJobService(store db.Store) *jobService {
	return &jobService{store: store}
}

// update password
// TODO: improve this function to be able to reveive any change from clients
func (u *jobService) UpdateUserSetting(ctx context.Context, userId int64, req *delivery.UpdateUserSettingRequest) (*response.UserSettingResponse, error) {
	user, err := u.store.UpdateUser(ctx, db.UpdateUserParams{
		ID: sql.NullInt64{
			Int64: userId,
			Valid: true,
		},
		Username: sql.NullString{
			String: req.Username,
			Valid:  true,
		},
		AlarmSound: sql.NullString{
			String: req.AlarmSound,
			Valid:  true,
		},
		RepeatAlarm: sql.NullInt32{
			Int32: req.RepeatAlarm,
			Valid: true,
		},
	})
	if err != nil {
		return nil, err
	}

	rsp := response.UserSettingResponse{
		Username:    user.Username,
		AlarmSound:  user.AlarmSound,
		RepeatAlarm: user.RepeatAlarm,
	}

	return &rsp, nil
}
