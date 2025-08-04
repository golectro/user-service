package migrations

import (
	"golectro-user/internal/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}
