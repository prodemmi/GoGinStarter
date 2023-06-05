package migrator

import (
	"GoGinStarter/database/migrations/user"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var Migrations = []*gormigrate.Migration{
	&user.InitMigration,
}

func NewMigrator(db *gorm.DB) *gormigrate.Gormigrate {
	return gormigrate.New(db, gormigrate.DefaultOptions, Migrations)
}
