package migrator

import (
	"GoGinStarter/database/migrations/otp_tokens"
	"GoGinStarter/database/migrations/user"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var Migrations = []*gormigrate.Migration{
	&user.InitMigration,
	&otp_tokens.InitMigration,
	&user.AddRememberTokenMigration,
	&user.AddMobileMigration,
}

func NewMigrator(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, Migrations)
}
