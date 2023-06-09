package user

import (
	"GoGinStarter/app/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var AddMobileMigration = gormigrate.Migration{
	ID: "add_mobile",
	Migrate: func(tx *gorm.DB) error {
		return tx.Migrator().AddColumn(&models.User{}, "mobile")
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropColumn(&models.User{}, "mobile")
	},
}
