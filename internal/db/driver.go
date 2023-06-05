package db

import "gorm.io/gorm"

type Driver interface {
	Connect(dsn string) (*gorm.DB, error)
}
