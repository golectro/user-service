package entity

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID            uuid.UUID            `gorm:"type:char(36);primaryKey" json:"id"`
	UserID        uuid.UUID            `gorm:"type:char(36);not null;column:user_id" json:"user_id"`
	User          User                 `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Label         string               `gorm:"type:text" json:"label"`
	Recipient     string               `gorm:"type:text;not null" json:"recipient"`
	Phone         string               `gorm:"type:text;not null" json:"phone"`
	AddressLine   string               `gorm:"type:text;not null;column:address_line" json:"address_line"`
	City          string               `gorm:"type:text;not null" json:"city"`
	Province      string               `gorm:"type:text;not null" json:"province"`
	PostalCode    string               `gorm:"type:text;not null;column:postal_code" json:"postal_code"`
	IsDefault     bool                 `gorm:"type:tinyint(1);default:0;check:is_default IN (0,1);column:is_default" json:"is_default"`
	Encrypted     bool                 `gorm:"type:tinyint(1);default:1;check:encrypted IN (0,1)" json:"encrypted"`
	EncryptionKey AddressEncryptionKey `gorm:"foreignKey:AddressID;constraint:OnDelete:CASCADE" json:"encryption_key,omitempty"`
	CreatedAt     time.Time            `gorm:"type:datetime;column:created_at;autoCreateTime:milli"`
	UpdatedAt     time.Time            `gorm:"type:datetime;column:updated_at;autoUpdateTime:milli"`
}

func (Address) TableName() string {
	return "addresses"
}
