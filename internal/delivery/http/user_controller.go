package http

import (
	"golectro-user/internal/constants"
	"golectro-user/internal/delivery/http/middleware"
	"golectro-user/internal/usecase"
	"golectro-user/internal/utils"
	"net/http"

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

	result, err := c.UseCase.Sync(ctx, auth)
	if err != nil {
		c.Log.WithError(err).Error("Failed to sync user")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.FailedSyncUser, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusCreated, constants.UserSynced, result)
	ctx.JSON(res.StatusCode, res)
}
