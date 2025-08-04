package usecase

import (
	"context"
	"golectro-user/internal/constants"
	"golectro-user/internal/entity"
	"golectro-user/internal/model"
	"golectro-user/internal/model/converter"
	"golectro-user/internal/repository"
	"golectro-user/internal/utils"

	"github.com/go-playground/validator/v10"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

func (uc *UserUseCase) Sync(ctx context.Context, auth *model.Auth) (*model.UserSyncResponse, error) {
	if err := uc.Validate.Struct(auth); err != nil {
		uc.Log.WithError(err).Error("Invalid input format")
		message := utils.TranslateValidationError(uc.Validate, err)
		return nil, utils.WrapMessageAsError(message)
	}

	user := &entity.User{
		ID:       auth.ID,
		Username: auth.Username,
		Email:    auth.Email,
		Roles:    auth.Roles,
	}

	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := uc.UserRepository.Sync(tx, user); err != nil {
		uc.Log.WithError(err).Error("Failed to create user")
		return nil, utils.WrapMessageAsError(constants.FailedSyncUser, err)
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.WithError(err).Error("Failed to commit transaction")
		return nil, utils.WrapMessageAsError(constants.FailedSyncUser, err)
	}

	return converter.UserToResponse(user), nil
}
