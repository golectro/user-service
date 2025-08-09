package usecase

import (
	"context"
	"encoding/base64"
	"golectro-user/internal/constants"
	"golectro-user/internal/entity"
	"golectro-user/internal/model"
	"golectro-user/internal/model/converter"
	"golectro-user/internal/repository"
	"golectro-user/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AddressUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	AddressRepository *repository.AddressRepository
	EncryptionRepository *repository.EncryptionRepository
	EncryptionUsecase *EncryptionUsecase
}

func NewAddressUsecase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, userRepository *repository.AddressRepository, encryptionRepository *repository.EncryptionRepository, encryptionUsecase *EncryptionUsecase) *AddressUseCase {
	return &AddressUseCase{
		DB:                db,
		Log:               log,
		Validate:          validate,
		AddressRepository: userRepository,
		EncryptionRepository: encryptionRepository,
		EncryptionUsecase: encryptionUsecase,
	}
}

func (uc *AddressUseCase) GetAddressesByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]model.UserAddressResponse, int64, error) {
	address, total, err := uc.AddressRepository.FindByUserID(uc.DB.WithContext(ctx), userID, limit, offset)
	if err != nil {
		uc.Log.WithError(err).Error("Failed to find addresses by user ID")
		return nil, 0, utils.WrapMessageAsError(constants.FailedGetAddresses, err)
	}

	if address == nil {
		return nil, 0, utils.WrapMessageAsError(constants.FailedGetAddresses)
	}

	return converter.ToUserAddressResponses(address), total, nil

}

func (uc *AddressUseCase) CreateAddress(ctx context.Context, request *model.UserAddressRequest, userID uuid.UUID) (*model.UserAddressResponse, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := uc.Validate.Struct(request); err != nil {
		message := utils.TranslateValidationError(uc.Validate, err)
		return nil, utils.WrapMessageAsError(message)
	}

	dek, err := uc.EncryptionUsecase.GenerateDEK()
	if err != nil {
		return nil, utils.WrapMessageAsError(constants.FailedGeneateDEK, err)
	}

	encrypt := func(plaintext string) (string, error) {
		ciphertext, err := uc.EncryptionUsecase.EncryptAES_GCM([]byte(plaintext), dek)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(ciphertext), nil
	}

	labelEncrypted, err := encrypt(request.Label)
	if err != nil {
		return nil, utils.WrapMessageAsError(constants.FailedEncryptLabel, err)
	}

	recipientEncrypted, err := encrypt(request.Recipient)
	if err != nil {
		return nil, utils.WrapMessageAsError(constants.FailedEncryptRecipient, err)
	}

	phoneEncrypted, err := encrypt(request.Phone)
	if err != nil {
		return nil, utils.WrapMessageAsError(constants.FailedEncryptPhone, err)
	}

	addressLineEncrypted, err := encrypt(request.AddressLine)
	if err != nil {
		return nil, utils.WrapMessageAsError(constants.FailedEncryptAddressLine, err)
	}

	cityEncrypted, err := encrypt(request.City)
	if err != nil {
		return nil, utils.WrapMessageAsError(constants.FailedEncryptCity, err)
	}

	provinceEncrypted, err := encrypt(request.Province)
	if err != nil {
		return nil, utils.WrapMessageAsError(constants.FailedEncryptProvince, err)
	}

	postalCodeEncrypted, err := encrypt(request.PostalCode)
	if err != nil {
		return nil, utils.WrapMessageAsError(constants.FailedEncryptPostalCode, err)
	}

	encryptedDEK, err := uc.EncryptionUsecase.EncryptDEK(dek)
	if err != nil {
		return nil, utils.WrapMessageAsError(constants.FailedEncryptDEK, err)
	}

	addressID := uuid.New()
	address := &entity.Address{
		ID:            addressID,
		UserID:        userID,
		Label:         labelEncrypted,
		Recipient:     recipientEncrypted,
		Phone:         phoneEncrypted,
		AddressLine:   addressLineEncrypted,
		City:          cityEncrypted,
		Province:      provinceEncrypted,
		PostalCode:    postalCodeEncrypted,
		IsDefault:     request.IsDefault,
	}

	if err := uc.AddressRepository.Create(tx, address); err != nil {
		uc.Log.WithError(err).Error("Failed to create address")
		return nil, utils.WrapMessageAsError(constants.FailedCreateAddress, err)
	}

	keyEntity := &entity.AddressEncryptionKey{
		ID:        uuid.New(),
		AddressID: addressID,
		Key:       encryptedDEK,
	}

	if err := uc.EncryptionRepository.Create(tx, keyEntity); err != nil {
		uc.Log.WithError(err).Error("Failed to create address encryption key")
		return nil, utils.WrapMessageAsError(constants.FailedCreateAddress, err)
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.WithError(err).Error("Failed to commit transaction")
		return nil, utils.WrapMessageAsError(constants.FailedCreateAddress, err)
	}

	return converter.ToUserAddressResponse(address), nil
}

func (uc *AddressUseCase) UpdateAddress(ctx context.Context, request *model.UserAddressRequest, addressID uuid.UUID, userID uuid.UUID) (*model.UserAddressResponse, error) {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := uc.Validate.Struct(request); err != nil {
		message := utils.TranslateValidationError(uc.Validate, err)
		return nil, utils.WrapMessageAsError(message)
	}

	address, err := uc.AddressRepository.FindByID(tx, addressID)
	if err != nil {
		uc.Log.WithError(err).Error("Failed to find address by ID")
		return nil, utils.WrapMessageAsError(constants.FailedGetAddresses, err)
	}

	if address == nil || address.UserID != userID {
		return nil, utils.WrapMessageAsError(constants.AddressNotFound)
	}

	address.Label = request.Label
	address.Recipient = request.Recipient
	address.Phone = request.Phone
	address.AddressLine = request.AddressLine
	address.City = request.City
	address.Province = request.Province
	address.PostalCode = request.PostalCode
	address.IsDefault = request.IsDefault

	if err := uc.AddressRepository.Update(tx, address); err != nil {
		uc.Log.WithError(err).Error("Failed to update address")
		return nil, utils.WrapMessageAsError(constants.FailedUpdateAddress, err)
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.WithError(err).Error("Failed to commit transaction")
		return nil, utils.WrapMessageAsError(constants.FailedUpdateAddress, err)
	}

	return converter.ToUserAddressResponse(address), nil
}

func (uc *AddressUseCase) SetDefaultAddress(ctx context.Context, addressID uuid.UUID, userID uuid.UUID) error {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	address, err := uc.AddressRepository.FindByID(tx, addressID)
	if err != nil {
		uc.Log.WithError(err).Error("Failed to find address by ID")
		return utils.WrapMessageAsError(constants.FailedGetAddresses, err)
	}

	if address == nil || address.UserID != userID {
		return utils.WrapMessageAsError(constants.AddressNotFound)
	}

	address.IsDefault = true

	if err := uc.AddressRepository.Update(tx, address); err != nil {
		uc.Log.WithError(err).Error("Failed to set default address")
		return utils.WrapMessageAsError(constants.FailedSetDefaultAddress, err)
	}

	if err := uc.AddressRepository.UnsetOtherDefaultAddresses(tx, userID, addressID); err != nil {
		uc.Log.WithError(err).Error("Failed to unset other default addresses")
		return utils.WrapMessageAsError(constants.FailedSetDefaultAddress, err)
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.WithError(err).Error("Failed to commit transaction")
		return utils.WrapMessageAsError(constants.FailedSetDefaultAddress, err)
	}

	return nil
}

func (uc *AddressUseCase) DeleteAddress(ctx context.Context, addressID uuid.UUID, userID uuid.UUID) error {
	tx := uc.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	address, err := uc.AddressRepository.FindByID(tx, addressID)
	if err != nil {
		uc.Log.WithError(err).Error("Failed to find address by ID")
		return utils.WrapMessageAsError(constants.FailedGetAddresses, err)
	}

	if address == nil || address.UserID != userID {
		return utils.WrapMessageAsError(constants.AddressNotFound)
	}

	if err := uc.AddressRepository.Delete(tx, address); err != nil {
		uc.Log.WithError(err).Error("Failed to delete address")
		return utils.WrapMessageAsError(constants.FailedDeleteAddress, err)
	}

	if err := tx.Commit().Error; err != nil {
		uc.Log.WithError(err).Error("Failed to commit transaction")
		return utils.WrapMessageAsError(constants.FailedDeleteAddress, err)
	}

	return nil
}
