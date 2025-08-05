package route

import (
	"golectro-user/internal/delivery/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
)

type RouteConfig struct {
	App               *gin.Engine
	Minio             *minio.Client
	AuthMiddleware    gin.HandlerFunc
	UserController    *http.UserController
	AddressController *http.AddressController
	Viper             *viper.Viper
}

func (c *RouteConfig) Setup() {
	api := c.App.Group("/api")

	c.RegisterUserRoutes(api, c.Minio)
	c.RegisterCommonRoutes(c.App)
	c.RegisterAddressRoutes(api)
}
