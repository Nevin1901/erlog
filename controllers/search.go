package controllers

import (
	"github.com/Nevin1901/arlog/models"
	"github.com/gin-gonic/gin"
)

func AllLogsController(c *gin.Context) {
	var logs []models.ErLog
	result := models.DB.Find(&logs)

	if result.Error != nil {
		println("Error")
		c.String(400, "Error")
		return
	}

	c.JSON(200, logs)
}

func SearchByMessageController(c *gin.Context) {
	id := c.Param("id")
	print(id)
	message := c.Param("message")

	var logs []models.ErLog
	result := models.DB.Find(&logs, &models.ErLog{Message: message})

	if result.Error != nil {
		c.String(400, "Error getting logs")
		return
	}

	c.JSON(200, logs)
}