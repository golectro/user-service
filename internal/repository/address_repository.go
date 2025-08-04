package repository

import (
	"golectro-user/internal/entity"

	"github.com/sirupsen/logrus"
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
