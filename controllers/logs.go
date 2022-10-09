package controllers

import (
	"github.com/Nevin1901/arlog/models"
	"github.com/gin-gonic/gin"
)

type SearchParams struct {
	Field	string	`json:"field"`
	Value	string	`json:"value"`
}

func SearchController(c *gin.Context) {
	var params SearchParams
	paramErr := c.ShouldBindJSON(&params)

	if paramErr != nil {
		c.String(400, "Incorrect params")
		return
	}

	var logs []models.ErLog
	println(params.Field)
	// WARNING: this line might not be safe to use
	result := models.DB.Where("message LIKE ? OR extra_data LIKE ?", "%" + params.Value + "%", "%" + params.Value + "%").Find(&logs)

	if result.Error != nil {
		c.String(400, "Error in getting logs")
		return
	}

	c.JSON(200, logs)
}