package controllers

import (
	"net/http"
	"strconv"

	"github.com/Nevin1901/arlog/models"
	"github.com/Nevin1901/arlog/routines"
	"github.com/gin-gonic/gin"
)

type SearchParams struct {
	Value	string	`json:"value"`
}

func AddLogController(c *gin.Context) {
	var logObj models.ErLog

	errLog := c.ShouldBindJSON(&logObj)

	if errLog != nil {
		println("Failed binding body")
		c.String(http.StatusBadRequest, "Failed binding body")
		return
	}

	hash := models.GetMD5Hash(logObj.Message)
	var exists bool
	result := models.DB.Model(&models.IgnoreList{}).
		Select("count(*) > 0").
		Where("hash = ?", hash).
		Find(&exists).Error

	if result != nil {
		c.String(400, "Error")
		return
	}

	if exists == true {
		println("ignored")
		c.String(200, "Ignored")
		return
	}

	models.DB.Create(&logObj)
	go routines.AppendLog(2)
	println("done")
	c.String(200, "OK")
}

func SearchController(c *gin.Context) {
	var params SearchParams
	paramErr := c.ShouldBindJSON(&params)

	if paramErr != nil {
		c.String(400, "Incorrect params")
		return
	}

	var logs []models.ErLog
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

type MessageCount struct {
	ID		int		`json:"id"`
	Title	string	`json:"title"`
	Message	string	`json:"message"`
	Num		int		`json:"num"`
	LogType string	`json:"logType"`
}

func CountController(c *gin.Context) {
	var count []MessageCount
	search := c.Query("search")
	println(search)

	result := models.DB.Raw("SELECT id, title, message, log_type, COUNT(*) AS `num` FROM er_logs WHERE (message LIKE ? OR extra_data LIKE ? OR title LIKE ?) AND deleted_at IS NULL GROUP BY message", "%"+search+"%", "%"+search+"%", "%"+search+"%").Scan(&count)
	if result.Error != nil {
		c.String(400, "Error getting logs")
		return
	}

	c.JSON(200, count)
}

func RotateLogController(c *gin.Context) {
	max_log_amt := 2

	// str_amount := strconv.Itoa(max_log_amt)

	var count int64
	result := models.DB.Find(&models.ErLog{}).Count(&count)

	delete_amt := int(count) - max_log_amt

	if result.Error != nil {
		c.String(400, "Error counting logs")
		return
	}

	var to_delete []models.ErLog
	models.DB.Order("id ASC").Limit(delete_amt).Find(&to_delete)
	models.DB.Delete(&to_delete)

	// models.DB.Exec("DELETE FROM er_logs WHERE id in (SELECT id FROM er_logs ORDER BY id ASC LIMIT ?)", str_amount)

	c.String(200, "Deleted Logs")
}
