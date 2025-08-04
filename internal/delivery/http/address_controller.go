package http

import (
	"golectro-user/internal/constants"
	"golectro-user/internal/delivery/http/middleware"
	"golectro-user/internal/usecase"
	"golectro-user/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AddressController struct {
	Log     *logrus.Logger
	UseCase *usecase.AddressUseCase
}

func NewAddressController(useCase *usecase.AddressUseCase, log *logrus.Logger) *AddressController {
	return &AddressController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *AddressController) GetAddressByUserID(ctx *gin.Context) {
	auth := middleware.GetUser(ctx)

	addresses, err := c.UseCase.GetAddressesByUserID(ctx, auth.ID)
	if err != nil {
		c.Log.WithError(err).Error("Failed to get addresses")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.FailedGetAddresses, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AddressesRetrieved, addresses)
	ctx.JSON(res.StatusCode, res)
}
