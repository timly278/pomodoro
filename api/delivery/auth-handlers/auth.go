package auth

import (
	"pomodoro/api/delivery"
	"pomodoro/api/service"
)

type authHandlers struct {
	authService service.AuthService
	userService service.UserService
}

var _ delivery.AuthHandlers = (*authHandlers)(nil)

func NewAuthHandlers(authService service.AuthService, userService service.UserService) delivery.AuthHandlers {
	return &authHandlers{authService: authService, userService: userService}
}
