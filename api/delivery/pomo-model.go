package delivery

type CreatePomodoroRequest struct {
	TypeID      int64 `json:"type_id" binding:"required,min=1"`
	TaskID      int64 `json:"task_id"`
	FocusDegree int32 `json:"focus_degree" binding:"required,min=1,max=5"`
}
