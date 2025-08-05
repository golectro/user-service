package http

import (
	"golectro-user/internal/constants"
	"golectro-user/internal/delivery/http/middleware"
	"golectro-user/internal/model"
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

func (c *AddressController) CreateAddress(ctx *gin.Context) {
	auth := middleware.GetUser(ctx)
	request := new(model.UserAddressRequest)

	if err := ctx.ShouldBindJSON(request); err != nil {
		c.Log.WithError(err).Error("Invalid request data")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.InvalidRequestData, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	result, err := c.UseCase.CreateAddress(ctx, request, auth.ID)
	if err != nil {
		c.Log.WithError(err).Error("Failed to create address")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.FailedGetAddresses, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusCreated, constants.AddressesRetrieved, result)
	ctx.JSON(res.StatusCode, res)

}
