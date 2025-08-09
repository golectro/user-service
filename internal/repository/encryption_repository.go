package repository

import (
	"golectro-user/internal/entity"

	"github.com/sirupsen/logrus"
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