package config

import (
	"golectro-user/internal/delivery/grpc"
	"golectro-user/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func StartGRPC(viper *viper.Viper, db *gorm.DB, validate *validator.Validate, log *logrus.Logger) {
	addressUC := usecase.NewAddressUsecase(db, log, validate, nil)
	port := viper.GetInt("GRPC_PORT")
	grpc.StartGRPCServer(addressUC, port, viper)
}
