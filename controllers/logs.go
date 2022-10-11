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
	println(id)

	i64, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.String(400, "Failed convering str to int")
		return
	}

	intId := uint(i64)


	result := models.DB.Raw("SELECT * FROM er_logs WHERE deleted_at IS NULL AND message = (SELECT message FROM er_logs WHERE id = ?)", intId).Scan(&log)
	// result := models.DB.Find(&log, &models.ErLog{ID: intId})

	if result.Error != nil {
		c.String(400, "Error")
		return
	}

	c.JSON(200, log)
}

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

	hash := models.GetMD5Hash(log.Message)
	ignore_log := models.IgnoreList{Hash: hash}
	models.DB.Create(&ignore_log)

	var deletedLogs []models.ErLog
	models.DB.Where("message = ?", log.Message).Delete(&deletedLogs)
	c.String(200, "Ok")
}

type MessageCount struct {
	ID		int		`json:"id"`
	Title	string	`json:"title"`
	Message	string	`json:"message"`
	Num		int		`json:"num"`
}

func CountController(c *gin.Context) {
	var count []MessageCount

	result := models.DB.Raw("SELECT id, title, message, COUNT(*) AS `num` FROM er_logs WHERE deleted_at IS NULL GROUP BY message").Scan(&count)
	if result.Error != nil {
		c.String(400, "Error getting logs")
		return
	}

	c.JSON(200, count)
}
