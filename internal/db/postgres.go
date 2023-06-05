package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDriver struct{}

func (d *PostgresDriver) Connect(dsn string) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return gormDB, nil
}
