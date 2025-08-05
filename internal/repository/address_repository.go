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

func (r *AddressRepository) FindByUserID(db *gorm.DB, userID uuid.UUID, limit, offset int) ([]entity.Address, int64, error) {
	var addresses []entity.Address
	var total int64

	query := db.Model(&entity.Address{}).Where("user_id = ?", userID)

	// Hitung total
	if err := query.Count(&total).Error; err != nil {
		r.Log.WithError(err).Error("Failed to count addresses by user ID")
		return nil, 0, err
	}

	// Ambil data paginated
	err := query.Limit(limit).Offset(offset).Find(&addresses).Error
	if err != nil {
		r.Log.WithError(err).Error("Failed to find paginated addresses by user ID")
		return nil, 0, err
	}

	return addresses, total, nil
}

func (r *AddressRepository) FindByID(db *gorm.DB, addressID uuid.UUID) (*entity.Address, error) {
	var address entity.Address
	err := db.First(&address, "id = ?", addressID).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		r.Log.WithError(err).Error("Failed to find address by ID")
		return nil, err
	}

	return &address, nil
}

func (r *AddressRepository) UnsetOtherDefaultAddresses(db *gorm.DB, userID uuid.UUID, addressID uuid.UUID) error {
	err := db.Model(&entity.Address{}).
		Where("user_id = ? AND id != ?", userID, addressID).
		Update("is_default", false).Error

	if err != nil {
		r.Log.WithError(err).Error("Failed to unset other default addresses")
		return err
	}

	return nil
}
