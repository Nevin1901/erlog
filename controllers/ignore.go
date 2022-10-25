package controllers

import (
	"strconv"

	"github.com/Nevin1901/arlog/models"
	"github.com/gin-gonic/gin"
)


func IgnoreLogController(c *gin.Context) {
	id := c.Param("id")

	i64, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.String(400, "Failed converting str to int")
		return
	}

	intId := uint(i64)

	var log models.ErLog

	result := models.DB.First(&log, &models.ErLog{ID: intId})

	if result.Error != nil {
		c.String(400, "Error looking up log")
		return
	}

	var preview_length int
	length := len(log.Message)

	if length < 40 {
		preview_length = len(log.Message)
	} else {
		preview_length = 40
	}

	preview := log.Message[0:preview_length]

	if preview_length == 40 {
		preview += "..."
	}

	hash := models.GetMD5Hash(log.Message)
	ignore_log := models.IgnoreList{Hash: hash, Preview: preview}
	models.DB.Create(&ignore_log)

	var deletedLogs []models.ErLog
	models.DB.Where("message = ?", log.Message).Delete(&deletedLogs)
	c.String(200, "Ok")
}

func GetIgnoredController(c *gin.Context) {
	var ignored []models.IgnoreList
	result := models.DB.Find(&ignored, &models.IgnoreList{})

	if result.Error != nil {
		c.String(400, "Error retrieving ignored")
		return
	}

	c.JSON(200, ignored)
}

type UnIgnoreParams struct {
	Hash	string	`json:"hash"`
}
func UnIgnoreController(c *gin.Context) {
	var params UnIgnoreParams
	result := c.ShouldBindJSON(&params)

	if result != nil {
		c.String(400, "Expected Hash");
		return
	}

	var log []models.IgnoreList

	models.DB.
	Model(&models.IgnoreList{}).
	Where("hash = ?", params.Hash).
	Delete(&log)

	c.String(200, "Ok")
}