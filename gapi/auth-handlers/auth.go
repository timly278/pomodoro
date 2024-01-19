package gapi

import (
	"context"
	"pomodoro/api/delivery"
	"pomodoro/api/service"
	"pomodoro/pb"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SHould this be interface or struct???
type AuthHandlers struct {
	pb.UnimplementedAuthHandlersServer
	authService service.AuthService
	userService service.UserService
	logger      *zap.Logger
}

func NewAuthHandlers(authService service.AuthService, userService service.UserService, logger *zap.Logger) *AuthHandlers {
	return &AuthHandlers{authService: authService, userService: userService, logger: logger}
}

func (a *AuthHandlers) Register(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {

	deliReq := &delivery.CreateUserRequest{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	}

	user, _, err := a.userService.CreateUser(ctx, deliReq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to create user: %s", err)
	}

	err = a.authService.SendEmailVerification(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}, nil
}

func (a *AuthHandlers) Login(context.Context, *pb.LoginRequest) (*pb.UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

