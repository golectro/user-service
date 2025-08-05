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

func (uc *UserController) SyncUser(ctx *gin.Context) {
	auth := middleware.GetUser(ctx)

	result, err := uc.UseCase.Sync(ctx, auth)
	if err != nil {
		uc.Log.WithError(err).Error("Failed to sync user")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.FailedSyncUser, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusCreated, constants.UserSynced, result)
	ctx.JSON(res.StatusCode, res)
}

func (uc *UserController) UploadAvatar(ctx *gin.Context) {
	auth := middleware.GetUser(ctx)

	uploadedFile := ctx.MustGet("uploadedFile").(map[string]any)

	result, err := uc.UseCase.UploadAvatar(ctx, auth, uploadedFile)
	if err != nil {
		uc.Log.WithError(err).Error("Failed to upload avatar")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.FailedUploadAvatar, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusCreated, constants.AvatarUploaded, result)
	ctx.JSON(res.StatusCode, res)
}
