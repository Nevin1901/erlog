package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ErLog struct {
	ID			uint			`json:"id" gorm:"primarykey"`
	CreatedAt	time.Time		`json:"createdAt"`
	UpdatedAt	time.Time		`json:"updatedAt"`
	DeletedAt	gorm.DeletedAt	`json:"deletedAt" gorm:"index"`
	LogType		string			`json:"logType"`
	Message		string			`json:"message"`
	Title		string			`json:"title"`
	ExtraData	datatypes.JSON	`json:"extraData"`
}
