package delivery

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6,max=12"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdateUserSettingRequest struct {
	Username    string `json:"username" binding:"required,alphanum"`
	AlarmSound  string `json:"alarm_sound" binding:"required"`
	RepeatAlarm int32  `json:"repeat_alarm" binding:"required,min=1,max=10"`
}

type SendCodeRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type VerificationRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=12"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
