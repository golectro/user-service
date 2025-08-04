package usecase

import (
	"golectro-user/internal/repository"

	"github.com/go-playground/validator/v10"

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
