package config

import (
	"golectro-user/internal/delivery/http"
	"golectro-user/internal/delivery/http/middleware"
	"golectro-user/internal/delivery/http/route"
	"golectro-user/internal/repository"
	"golectro-user/internal/usecase"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB         *gorm.DB
	Mongo      *mongo.Database
	App        *gin.Engine
	Redis      *redis.Client
	Log        *logrus.Logger
	Validate   *validator.Validate
	Viper      *viper.Viper
	GRPCClient *grpc.ClientConn
	Elastic    *elasticsearch.Client
	Minio      *minio.Client
}

func Bootstrap(config *BootstrapConfig) {
	userRepository := repository.NewUserRepository(config.Log)
	addressRepository := repository.NewAddressRepository(config.Log)

	userUseCase := usecase.NewUserUsecase(config.DB, config.Log, config.Validate, userRepository)
	addressUseCase := usecase.NewAddressUsecase(config.DB, config.Log, config.Validate, addressRepository)

	userController := http.NewUserController(userUseCase, config.Log)
	addressController := http.NewAddressController(addressUseCase, config.Log)

	authMiddleware := middleware.NewAuth(userUseCase, config.Viper)

	routeConfig := route.RouteConfig{
		App:               config.App,
		UserController:    userController,
		AddressController: addressController,
		AuthMiddleware:    authMiddleware,
	}
	routeConfig.Setup()
}
