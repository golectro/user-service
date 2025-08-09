package repository

import (
	"errors"
	"golectro-user/internal/entity"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type EncryptionRepository struct {
	Repository[entity.AddressEncryptionKey]
	Log *logrus.Logger
}

func NewEncryptionRepository(log *logrus.Logger) *EncryptionRepository {
	return &EncryptionRepository{
		Log: log,
	}
}

func (r *EncryptionRepository) FindByAddressID(db *gorm.DB, addressID uuid.UUID) (*entity.AddressEncryptionKey, error) {
	var key entity.AddressEncryptionKey
	err := db.First(&key, "address_id = ?", addressID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		r.Log.WithError(err).Error("Failed to find encryption key by address ID")
		return nil, err
	}

	return &key, nil
}
