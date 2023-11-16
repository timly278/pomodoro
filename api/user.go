package api

import (
	"net/http"
	db "pomodoro/db/sqlc"
	"pomodoro/util"
	"time"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6,max=12"`
	Email    string `json:"email" binding:"required,email"`
}

type createUserResponse struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username" `
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) createUserResponse {
	return createUserResponse{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

// CreateUser - user signing up
func (server *Server) CreateUser(ctx *gin.Context) {
	var newUser createUserRequest
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(newUser.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	createUserParams := db.CreateUserParams{
		Username:       newUser.Username,
		HashedPassword: hashedPassword,
		Email:          newUser.Email,
	}
	createdUser, err := server.store.CreateUser(ctx, createUserParams)
	if err != nil {
		//todo: handle detail error of pq here
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(createdUser)
	ctx.JSON(http.StatusOK, response{
		Message: "register successfully",
		Data:    rsp,
	})
}

func (server *Server) UserLogin(ctx *gin.Context) {

}

func (server *Server) UserLogout(ctx *gin.Context) {

}

func (server *Server) Dosomething(ctx *gin.Context) {

}
