package entity

import (
	"time"

	"github.com/google/uuid"
)

type AddressEncryptionKey struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	AddressID uuid.UUID `gorm:"type:char(36);not null;uniqueIndex;column:address_id" json:"address_id"`
	Key       string    `gorm:"type:text;not null;column:key" json:"key"`

	CreatedAt time.Time `gorm:"type:datetime;column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time `gorm:"type:datetime;column:updated_at;autoUpdateTime:milli"`

	Address *Address `gorm:"foreignKey:AddressID;references:ID" json:"-"`
}

func (AddressEncryptionKey) TableName() string {
	return "encryption_keys"
}