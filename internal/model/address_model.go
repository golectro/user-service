package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserAddressRequest struct {
		Label       string `json:"label" validate:"required,max=50"`
		Recipient   string `json:"recipient" validate:"required,max=100"`
		Phone       string `json:"phone" validate:"required,max=20"`
		AddressLine string `json:"address_line" validate:"required,max=255"`
		City        string `json:"city" validate:"required,max=100"`
		Province    string `json:"province" validate:"required,max=100"`
		PostalCode  string `json:"postal_code" validate:"required,max=10"`
		IsDefault   bool   `json:"is_default"`
	}
	UserAddressResponse struct {
		ID          uuid.UUID `json:"id"`
		Label       string    `json:"label"`
		Recipient   string    `json:"recipient"`
		Phone       string    `json:"phone"`
		AddressLine string    `json:"address_line"`
		City        string    `json:"city"`
		Province    string    `json:"province"`
		PostalCode  string    `json:"postal_code"`
		IsDefault   bool      `json:"is_default"`
		Encrypted   bool      `json:"encrypted"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	EditUserAddressRequest struct {
		Label       string `json:"label" validate:"omitempty,min=3,max=50"`
		Recipient   string `json:"recipient" validate:"omitempty,min=3,max=100"`
		Phone       string `json:"phone" validate:"omitempty,min=8,max=20"`
		AddressLine string `json:"address_line" validate:"omitempty,min=5,max=255"`
		City        string `json:"city" validate:"omitempty,min=3,max=100"`
		Province    string `json:"province" validate:"omitempty,min=3,max=100"`
		PostalCode  string `json:"postal_code" validate:"omitempty,min=4,max=10"`
		IsDefault   bool   `json:"is_default"`
	}

	AddressEncryptionKeyResponse struct {
		ID        uuid.UUID `json:"id"`
		AddressID uuid.UUID `json:"address_id"`
		Key       string    `json:"key"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
