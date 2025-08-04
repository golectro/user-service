package route

import (
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App               *gin.Engine
}

func (c *RouteConfig) Setup() {
	c.RegisterCommonRoutes(c.App)
}
