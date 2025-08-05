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
	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	AddressRepository *repository.AddressRepository
}

func NewAddressUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, userRepository *repository.AddressRepository) *AddressUseCase {
	return &AddressUseCase{
		DB:                db,
		Log:               log,
		Validate:          validate,
		AddressRepository: userRepository,
	}
}

func (uc *AddressUseCase) GetAddressesByUserID(ctx context.Context, userID uuid.UUID) ([]model.UserAddressResponse, error) {
	address, err := uc.AddressRepository.FindByUserID(uc.DB.WithContext(ctx), userID)
	if err != nil {
		uc.Log.WithError(err).Error("Failed to find addresses by user ID")
		return nil, utils.WrapMessageAsError(constants.FailedGetAddresses, err)
	}

	if address == nil {
		return nil, utils.WrapMessageAsError(constants.FailedGetAddresses)
	}

	return converter.ToUserAddressResponses(address), nil

}

func (uc *AddressUseCase) CreateAddress(ctx context.Context, request *model.UserAddressRequest, userID uuid.UUID) (*model.UserAddressResponse, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := uc.Validate.Struct(request); err != nil {
		message := utils.TranslateValidationError(uc.Validate, err)
		return nil, utils.WrapMessageAsError(message)
	}

	address := &entity.Address{
		ID:          uuid.New(),
		UserID:      userID,
		Label:       request.Label,
		Recipient:   request.Recipient,
		Phone:       request.Phone,
		AddressLine: request.AddressLine,
		City:        request.City,
		Province:    request.Province,
		PostalCode:  request.PostalCode,
		IsDefault:   request.IsDefault,
	}

	if err := uc.AddressRepository.Create(tx, address); err != nil {
		uc.Log.WithError(err).Error("Failed to create address")
		return nil, utils.WrapMessageAsError(constants.FailedCreateAddress, err)
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.WithError(err).Error("Failed to commit transaction")
		return nil, utils.WrapMessageAsError(constants.FailedCreateAddress, err)
	}

	return converter.ToUserAddressResponse(address), nil
}
