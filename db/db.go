package db

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// Import driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Init() (*gorm.DB, error) {
	config := LoadConfig()
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", config.dbHost, config.dbPort, config.dbUser, config.dbName, config.dbPass)
	db, err := gorm.Open(config.dbType, connString)
	return db, err
}
