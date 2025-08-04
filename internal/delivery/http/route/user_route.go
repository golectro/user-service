package route

import "github.com/gin-gonic/gin"

func (c *RouteConfig) RegisterUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/users")

	user.POST("/sync", c.AuthMiddleware, c.UserController.SyncUser)
}
