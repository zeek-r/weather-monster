package db

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/zeek-r/weather-monster/models"

	// Import driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Init() (*gorm.DB, error) {
	config := LoadConfig()
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.dbHost, config.dbPort, config.dbUser, config.dbName, config.dbPass)
	db, err := gorm.Open(config.dbType, connString)

	if config.debug == true {
		// create tables in debug mode only, not recommened for prod
		db.AutoMigrate(&models.City{})
		db.AutoMigrate(&models.Temperature{})
		db.AutoMigrate(&models.Webhook{})
	}
	return db, err
}
