package usecase

import (
	"context"
	"golectro-user/internal/constants"
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
