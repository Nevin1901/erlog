package models

import (
	"time"

	"gorm.io/gorm"
)

type IgnoreList struct {
	ID		uint	`json:"id" gorm:"primarykey"`
	CreatedAt 	time.Time		`json:"createdAt"`
	UpdatedAt 	time.Time		`json:"updatedAt"`
	DeletedAt 	gorm.DeletedAt	`json:"deletedAt" gorm:"index"`
	Hash		string			`json:"hash"`
}