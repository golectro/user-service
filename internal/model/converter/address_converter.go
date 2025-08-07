package converter

import (
	"golectro-user/internal/entity"
	"golectro-user/internal/model"
)

func ToUserAddressResponses(addresses []entity.Address) []model.UserAddressResponse {
	var responses []model.UserAddressResponse
	for _, address := range addresses {
		responses = append(responses, model.UserAddressResponse{
			ID:          address.ID,
			Label:       address.Label,
			Recipient:   address.Recipient,
			Phone:       address.Phone,
			AddressLine: address.AddressLine,
			City:        address.City,
			Province:    address.Province,
			PostalCode:  address.PostalCode,
			IsDefault:   address.IsDefault,
			Encrypted:   address.Encrypted,
			CreatedAt:   address.CreatedAt,
			UpdatedAt:   address.UpdatedAt,
		})
	}
	return responses
}

func ToUserAddressResponse(address *entity.Address) *model.UserAddressResponse {
	if address == nil {
		return nil
	}
	return &model.UserAddressResponse{
		ID:          address.ID,
		Label:       address.Label,
		Recipient:   address.Recipient,
		Phone:       address.Phone,
		AddressLine: address.AddressLine,
		City:        address.City,
		Province:    address.Province,
		PostalCode:  address.PostalCode,
		IsDefault:   address.IsDefault,
		Encrypted:   address.Encrypted,
		CreatedAt:   address.CreatedAt,
		UpdatedAt:   address.UpdatedAt,
	}
}
