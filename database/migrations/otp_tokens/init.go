package otp_tokens

import (
	"GoGinStarter/app/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var InitMigration = gormigrate.Migration{
	ID: "otp_tokens_init",
	Migrate: func(tx *gorm.DB) error {
		return tx.Migrator().CreateTable(&models.OtpTokens{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("otp_tokens")
	},
}
