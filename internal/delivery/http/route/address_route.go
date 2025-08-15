package route

import "github.com/gin-gonic/gin"

func (c *RouteConfig) RegisterAddressRoutes(rg *gin.RouterGroup) {
	address := rg.Group("/address")

	address.GET("", c.AuthMiddleware, c.AddressController.GetAddressByUserID)
	address.POST("", c.AuthMiddleware, c.AddressController.CreateAddress)
	address.PUT("/:id", c.AuthMiddleware, c.AddressController.UpdateAddress)
	address.PUT("/:id/set-default", c.AuthMiddleware, c.AddressController.SetDefaultAddress)
	address.DELETE("/:id", c.AuthMiddleware, c.AddressController.DeleteAddress)

	address.GET("/encryption-key/:addressID", c.AuthMiddleware, c.AddressController.GetAddressEncryptionKey)
}
