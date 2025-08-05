package http

import (
	"golectro-user/internal/constants"
	"golectro-user/internal/delivery/http/middleware"
	"golectro-user/internal/model"
	"golectro-user/internal/usecase"
	"golectro-user/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UserController struct {
	Log          *logrus.Logger
	UserUseCase  *usecase.UserUseCase
	MinioUseCase *usecase.MinioUseCase
	Viper        *viper.Viper
}

func NewUserController(userUseCase *usecase.UserUseCase, minioUseCase *usecase.MinioUseCase, log *logrus.Logger, viper *viper.Viper) *UserController {
	return &UserController{
		Log:          log,
		UserUseCase:  userUseCase,
		MinioUseCase: minioUseCase,
		Viper:        viper,
	}
}

func (uc *UserController) SyncUser(ctx *gin.Context) {
	auth := middleware.GetUser(ctx)

	result, err := uc.UserUseCase.Sync(ctx, auth)
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

	result, err := uc.UserUseCase.UploadAvatar(ctx, auth, uploadedFile)
	if err != nil {
		uc.Log.WithError(err).Error("Failed to upload avatar")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.FailedUploadAvatar, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusCreated, constants.AvatarUploaded, result)
	ctx.JSON(res.StatusCode, res)
}

func (uc *UserController) DownloadAvatar(ctx *gin.Context) {
	auth := middleware.GetUser(ctx)

	result, err := uc.UserUseCase.FindUserByID(ctx, auth.ID)
	if err != nil {
		uc.Log.WithError(err).Error("Failed to find user by ID")
		res := utils.FailedResponse(ctx, http.StatusInternalServerError, constants.FailedFindUserByID, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	if result.AvatarObject == "" {
		uc.Log.Error("User avatar not found")
		res := utils.FailedResponse(ctx, http.StatusNotFound, constants.UserNotFound, nil)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	avatarURL, err := uc.MinioUseCase.GetPresignedURL(ctx, model.PresignedURLInput{
		Bucket:    uc.Viper.GetString("MINIO_BUCKET_AVATAR"),
		ObjectKey: result.AvatarObject,
		Expiry:    int64((time.Hour * 24).Seconds()),
	})
	if err != nil {
		uc.Log.WithError(err).Error("Failed to get presigned URL for avatar")
		res := utils.FailedResponse(ctx, http.StatusInternalServerError, constants.FailedGetPresignedURL, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AvatarDownloaded, gin.H{"url": avatarURL})
	ctx.JSON(res.StatusCode, res)
}
