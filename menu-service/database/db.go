package database

import (
"log"
"menu-service/models"
"gorm.io/driver/postgres"
"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	if err = DB.AutoMigrate(&models.MenuItem{}); err != nil {
		return err
	}

	log.Println("Menu-service database connected and migrated")
	return nil
}
