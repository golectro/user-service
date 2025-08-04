package route

import (
	"golectro-user/internal/delivery/http"

	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App            *gin.Engine
	AuthMiddleware gin.HandlerFunc
	UserController *http.UserController
}

func (c *RouteConfig) Setup() {
	api := c.App.Group("/api")

	c.RegisterUserRoutes(api)
	c.RegisterCommonRoutes(c.App)
}
