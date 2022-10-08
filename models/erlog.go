package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ErLog struct {
	gorm.Model
	LogType 	string			`json:"logType"`
	Message		string			`json:"message"`
	Title		string			`json:"title"`
	ExtraData	datatypes.JSON	`json:"extraData"`
}