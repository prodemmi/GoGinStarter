package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDriver struct{}

func (d *MySQLDriver) Connect(dsn string) (*gorm.DB, error) {
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return gormDB, nil
}
