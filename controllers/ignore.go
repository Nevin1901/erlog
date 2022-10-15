package controllers

import (
	"github.com/Nevin1901/arlog/models"
	"github.com/gin-gonic/gin"
)

func GetIgnoredController(c *gin.Context) {
	var ignored []models.IgnoreList
	result := models.DB.Find(&ignored, &models.IgnoreList{})

	if result.Error != nil {
		c.String(400, "Error retrieving ignored")
		return
	}

	c.JSON(200, ignored)
}