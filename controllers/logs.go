package controllers

import (
	"strconv"

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

func LogIdxController(c *gin.Context) {
	var log []models.ErLog
	id := c.Param("id")

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.String(200, "Failed convering str to int")
		return
	}

	result := models.DB.Find(&log, &models.ErLog{ID: intId})

	if result.Error != nil {
		c.String(400, "Error")
		return
	}

	c.String(200, "OK")
}

type MessageCount struct {
	ID		int		`json:"id"`
	Title	string	`json:"title"`
	Message	string	`json:"message"`
	Num		int		`json:"num"`
}

func CountController(c *gin.Context) {
	var count []MessageCount

	result := models.DB.Raw("SELECT id, title, message, COUNT(*) AS `num` FROM er_logs GROUP BY message").Scan(&count)
	if result.Error != nil {
		c.String(400, "Error getting logs")
		return
	}

	c.JSON(200, count)
}
