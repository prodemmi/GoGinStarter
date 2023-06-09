package user

import (
	"GoGinStarter/app/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AddRememberTokenMigration = gormigrate.Migration{
	ID: "add_remember_token",
	Migrate: func(tx *gorm.DB) error {
		return tx.Migrator().AddColumn(&models.User{}, "remember_token")
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropColumn(&models.User{}, "remember_token")
	},
}
