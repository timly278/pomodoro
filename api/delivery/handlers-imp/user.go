package handlers

import (
	"net/http"
	"pomodoro/api/delivery"
	"pomodoro/api/service"
	logging "pomodoro/api/service/service-imp"
	db "pomodoro/db/sqlc"
	"pomodoro/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type UserHandlers struct {
	userService service.User
}

var _ delivery.UserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(store db.Store, redisdb *redis.Client) *UserHandlers {
	userService := logging.NewUserLogging(store, redisdb)
	return &UserHandlers{userService: userService}
}

func (u *UserHandlers) CreateUser(ctx *gin.Context) {
	var req delivery.CreateUserRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	user, err := u.userService.CreateUser(ctx, req)
	if err != nil {
		// TODO: handle specific error i.e sql.NoRowErr
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	rsp := response.NewUserResponse(user)
	ctx.JSON(http.StatusOK, response.Response{
		Message: "create new user successfully",
		Data:    rsp,
	})
}

func (u *UserHandlers) Login(ctx *gin.Context) {
	var req delivery.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

}

func (u *UserHandlers) Logout(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, response.Response{
		Message: "not implemented feature",
	})
}

func (u *UserHandlers) UpdatePassword(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, response.Response{
		Message: "not implemented feature",
	})
}

func (u *UserHandlers) UpdateUserSetting(ctx *gin.Context) {
	var req delivery.UpdateUserSettingRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}

	newSetting, err := u.userService.UpdateUserSetting(ctx, getUserId(ctx), &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, response.Response{
		Message: "update setting successfully",
		Data:    newSetting,
	})
}
