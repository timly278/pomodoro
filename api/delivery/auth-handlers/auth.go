package auth

import (
	"pomodoro/api/delivery"
	"pomodoro/api/service"

	"go.uber.org/zap"
)

type authHandlers struct {
	authService service.AuthService
	userService service.UserService
	logger      *zap.Logger
}

var _ delivery.AuthHandlers = (*authHandlers)(nil)

func NewAuthHandlers(authService service.AuthService, userService service.UserService, logger *zap.Logger) delivery.AuthHandlers {
	return &authHandlers{authService: authService, userService: userService, logger: logger}
}
