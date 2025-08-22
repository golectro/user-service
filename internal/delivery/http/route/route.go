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
	SwaggerController *http.SwaggerController
	Viper             *viper.Viper
}

func (c *RouteConfig) Setup() {
	api := c.App.Group("/api/v1")

	c.RegisterUserRoutes(api, c.Minio)
	c.RegisterAddressRoutes(api)
	c.RegisterSwaggerRoutes(c.App)
	c.RegisterCommonRoutes(c.App)
}
