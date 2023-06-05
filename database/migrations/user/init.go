package user

import (
	"GoGinStarter/app/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var InitMigration = gormigrate.Migration{
	ID: "user_01",
	Migrate: func(tx *gorm.DB) error {
		return tx.Migrator().CreateTable(&models.User{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users")
	},
}
