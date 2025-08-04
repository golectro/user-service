package http

import (
	"golectro-user/internal/delivery/http/middleware"
	"golectro-user/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, log *logrus.Logger) *UserController {
	return &UserController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *UserController) SyncUser(ctx *gin.Context) {
	auth := middleware.GetUser(ctx)

	ctx.JSON(200, gin.H{
		"message": "User synced successfully",
		"user": gin.H{
			"id":       auth.ID,
			"email":    auth.Email,
			"username": auth.Username,
			"roles":    auth.Roles,
		},
	})
}
