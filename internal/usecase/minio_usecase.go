package usecase

import (
	"context"
	"golectro-user/internal/model"
	"golectro-user/internal/repository"
	"golectro-user/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type MinioUseCase struct {
	Repo     *repository.MinioRepository
	Validate *validator.Validate
	Log      *logrus.Logger
}

func NewMinioUsecase(repo *repository.MinioRepository, validate *validator.Validate, log *logrus.Logger) *MinioUseCase {
	return &MinioUseCase{
		Repo:     repo,
		Validate: validate,
		Log:      log,
	}
}

func (u *MinioUseCase) ValidateRequest(request *model.UploadFileRequest) error {
	if err := u.Validate.Struct(request); err != nil {
		u.Log.WithError(err).Error("Invalid input format")
		message := utils.TranslateValidationError(u.Validate, err)
		return utils.WrapMessageAsError(message)
	}
	return nil
}

func (u *MinioUseCase) Upload(ctx context.Context, input model.UploadFileInput) error {
	return u.Repo.UploadFile(ctx, input)
}

func (u *MinioUseCase) GetPresignedURL(ctx context.Context, input model.PresignedURLInput) (string, error) {
	return u.Repo.GeneratePresignedURL(ctx, input)
}

func (u *MinioUseCase) Delete(ctx context.Context, bucket, objectKey string) error {
	return u.Repo.DeleteFile(ctx, bucket, objectKey)
}
