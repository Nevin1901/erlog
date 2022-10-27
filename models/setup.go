package models

import (
	"crypto/md5"
	"encoding/hex"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func ConnectDB(db_name string) {
	database, err := gorm.Open(sqlite.Open(db_name), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to db")
	}

	database.AutoMigrate(&ErLog{})
	database.AutoMigrate(&IgnoreList{})

	// thank you https://phiresky.github.io/blog/2020/sqlite-performance-tuning/
	database.Exec("PRAGMA journal_mode=WAL;")
	database.Exec("PRAGMA synchronous=normal;")
	database.Exec("PRAGMA temp_store=memory;")
	database.Exec("PRAGMA mmap_size=30000000000;")


	DB = database
}
