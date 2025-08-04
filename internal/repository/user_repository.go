package repository

import (
	"errors"
	"golectro-user/internal/entity"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (r *UserRepository) Sync(db *gorm.DB, user *entity.User) error {
	if user.ID == uuid.Nil {
		r.Log.Error("User ID is nil")
		return errors.New("user ID cannot be nil")
	}

	existingUser, err := r.FindByID(db, user.ID)
	if err != nil {
		r.Log.WithError(err).Error("Failed to find user by ID")
		return err
	}

	if existingUser == nil {
		if err := r.Create(db, user); err != nil {
			r.Log.WithError(err).Error("Failed to create new user")
			return err
		}
	} else {
		user.CreatedAt = existingUser.CreatedAt

		if err := r.Update(db, user); err != nil {
			r.Log.WithError(err).Error("Failed to update existing user")
			return err
		}
	}

	return nil
}

func (r *UserRepository) FindByEmail(db *gorm.DB, email string) (*entity.User, error) {
	var user entity.User
	err := db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (r *UserRepository) FindByID(db *gorm.DB, id uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := db.Where("id = ?", id).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}
