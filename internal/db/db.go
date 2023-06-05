package db

import (
	"GoGinStarter/app/models"
	"GoGinStarter/internal/config"
	"GoGinStarter/internal/log"
	"fmt"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}

func ProvideDB(config *config.Config, log log.Log) *gorm.DB {
	var dbDriver Driver

	switch config.DB.Driver {
	case "mysql":
		dbDriver = &MySQLDriver{}
	case "postgres":
		dbDriver = &PostgresDriver{}
	default:
		errMsg := "Driver " + config.DB.Driver + " not found"
		log.Error(errMsg)
	}
	fmt.Println(config.DB.DSN)
	db, err := dbDriver.Connect(config.DB.DSN)

	if err != nil {
		log.Error(err.Error())
	}

	log.Info("Starting migrations")
	if mErr := migrate(db); mErr != nil {
		log.Error(mErr.Error())
	}
	log.Info("Migration ends")

	return db
}
