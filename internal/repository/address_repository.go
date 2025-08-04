package repository

import (
	"errors"
	"golectro-user/internal/entity"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressRepository struct {
	Repository[entity.Address]
	Log *logrus.Logger
}

func NewAddressRepository(log *logrus.Logger) *AddressRepository {
	return &AddressRepository{
		Log: log,
	}
}

func (r *AddressRepository) FindByUserID(db *gorm.DB, userID uuid.UUID) ([]entity.Address, error) {
	var addresses []entity.Address
	err := db.Where("user_id = ?", userID).Find(&addresses).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		r.Log.WithError(err).Error("Failed to find addresses by user ID")
		return nil, err
	}

	return addresses, nil
}
