package model

import (
	"time"

	"github.com/google/uuid"
)

type UserAddressRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type UserAddressResponse struct {
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
