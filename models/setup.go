package models

import (
	"crypto/md5"
	"encoding/hex"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to db")
	}

	print("got here")

	database.AutoMigrate(&ErLog{})
	database.AutoMigrate(&IgnoreList{})

	database.Exec("PRAGMA journal_mode=WAL;")
	database.Exec("PRAGMA synchronous=1")

	DB = database
}
