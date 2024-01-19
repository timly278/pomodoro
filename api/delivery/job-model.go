package delivery

import "time"

type CreatePomodoroRequest struct {
	TypeID      int64 `json:"type_id" binding:"required,min=1"`
	TaskID      int64 `json:"task_id"`
	FocusDegree int32 `json:"focus_degree" binding:"required,min=1,max=5"`
}

type GetPomodorosRequest struct {
	FromDate time.Time `form:"from_date" binding:"required" validate:"required,rfc3339"`
	ToDate   time.Time `form:"to_date" binding:"required" validate:"required,rfc3339"`
	Page     int32     `form:"page" binding:"required,min=1"`
	Limit    int32     `form:"limit" binding:"required,min=1"`
}

type GetStatisticRequest struct {
	FromDate time.Time `form:"from_date" binding:"required" validate:"required,rfc3339"`
	ToDate   time.Time `form:"to_date" binding:"required" validate:"required,rfc3339"`
}

type CreateNewTypeRequest struct {
	Name              string `json:"name" binding:"required,alphanum"`
	Color             string `json:"color" binding:"required"`
	Goalperday        int32  `json:"goal_per_day" binding:"required,min=1"`
	Duration          int32  `json:"duration" binding:"required,min=1"`
	Shortbreak        int32  `json:"shortbreak" binding:"required,min=1"`
	Longbreak         int32  `json:"longbreak" binding:"required,min=1"`
	Longbreakinterval int32  `json:"longbreakinterval" binding:"required,min=1"`
	AutostartBreak    bool   `json:"autostart_break" binding:"required,boolean"`
}

type UpdateUserSettingRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	AlarmSound  string `json:"alarm_sound" binding:"required"`
	RepeatAlarm int32  `json:"repeat_alarm" binding:"required,min=1,max=10"`
}
