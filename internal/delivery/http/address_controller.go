package http

import (
	"golectro-user/internal/constants"
	"golectro-user/internal/delivery/http/middleware"
	"golectro-user/internal/model"
	"golectro-user/internal/usecase"
	"golectro-user/internal/utils"
	"math"
	"net/http"
	"strconv"

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

	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	addresses, total, err := c.UseCase.GetAddressesByUserID(ctx, auth.ID, limit, offset)
	if err != nil {
		c.Log.WithError(err).Error("Failed to get addresses")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.FailedGetAddresses, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	pagination := model.PageMetadata{
		CurrentPage: page,
		PageSize:    limit,
		TotalPage:   int64(totalPages),
		TotalItem:   total,
		HasNext:     page < totalPages,
		HasPrevious: page > 1,
	}

	res := utils.SuccessWithPaginationResponse(ctx, http.StatusOK, constants.AddressesRetrieved, addresses, pagination)
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

func (c *AddressController) UpdateAddress(ctx *gin.Context) {
	auth := middleware.GetUser(ctx)
	addressID := ctx.Param("id")
	if addressID == "" {
		c.Log.Error("Address ID is required")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.InvalidRequestData, nil)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	request := new(model.UserAddressRequest)
	if err := ctx.ShouldBindJSON(request); err != nil {
		c.Log.WithError(err).Error("Invalid request data")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.InvalidRequestData, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	addressId, err := utils.ParseUUID(addressID)
	if err != nil {
		c.Log.WithError(err).Error("Invalid address ID format")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.InvalidRequestData, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	result, err := c.UseCase.UpdateAddress(ctx, request, addressId, auth.ID)
	if err != nil {
		c.Log.WithError(err).Error("Failed to update address")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.FailedUpdateAddress, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusOK, constants.UpdateAddressSuccess, result)
	ctx.JSON(res.StatusCode, res)
}

func (c *AddressController) SetDefaultAddress(ctx *gin.Context) {
	auth := middleware.GetUser(ctx)
	addressID := ctx.Param("id")
	if addressID == "" {
		c.Log.Error("Address ID is required")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.InvalidRequestData, nil)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	addressId, err := utils.ParseUUID(addressID)
	if err != nil {
		c.Log.WithError(err).Error("Invalid address ID format")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.InvalidRequestData, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	if err := c.UseCase.SetDefaultAddress(ctx, addressId, auth.ID); err != nil {
		c.Log.WithError(err).Error("Failed to set default address")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.FailedUpdateAddress, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusOK, constants.SetDefaultAddressSuccess, true)
	ctx.JSON(res.StatusCode, res)
}

func (c *AddressController) DeleteAddress(ctx *gin.Context) {
	auth := middleware.GetUser(ctx)
	addressID := ctx.Param("id")
	if addressID == "" {
		c.Log.Error("Address ID is required")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.InvalidRequestData, nil)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	addressId, err := utils.ParseUUID(addressID)
	if err != nil {
		c.Log.WithError(err).Error("Invalid address ID format")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.InvalidRequestData, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	if err := c.UseCase.DeleteAddress(ctx, addressId, auth.ID); err != nil {
		c.Log.WithError(err).Error("Failed to delete address")
		res := utils.FailedResponse(ctx, http.StatusBadRequest, constants.FailedDeleteAddress, err)
		ctx.AbortWithStatusJSON(res.StatusCode, res)
		return
	}

	res := utils.SuccessResponse(ctx, http.StatusOK, constants.AddressDeleted, true)
	ctx.JSON(res.StatusCode, res)
}
