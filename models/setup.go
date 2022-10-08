package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	print("got here")

	database.AutoMigrate(&ErLog{})

	DB = database
}