package main

import (
	"io/ioutil"
	"net/http"

	"github.com/Nevin1901/arlog/controllers"
	"github.com/Nevin1901/arlog/models"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// a super fast databse which basically stores a ton of stuff in memory, and then periodically writes to file when it can (when load is low)
// sqlite is not meant for processing large records, and clickhouse is a pain in the ass to set up. We are not using docker images just to host our server

func main() {
	models.ConnectDB()
	print("main")
	print("got here")

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.POST("/", func(c *gin.Context) {
		var logObj models.ErLog

		errLog := c.ShouldBindJSON(&logObj);

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

		// fmt.Printf("%+v\n", logObj)
		models.DB.Create(&logObj)
		println("done")
		c.String(200, "OK")
		return

		name := c.Param("name")

		jsonData, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			println(err)
			c.String(http.StatusOK, "Error reading request body")
			return
		}

		str1 := string(jsonData[:])

		if len(str1) == 0 {
			c.String(200, "no body")
			return
		}

		println(str1)

		c.JSON(http.StatusOK, gin.H {
			"message": name,
		})
	})

	r.POST("/all", func(c *gin.Context) {
		var logs []models.ErLog
		result := models.DB.Find(&logs)

		if result.Error != nil {
			println("Error")
			c.String(400, "Error")
			return
		}

		c.JSON(200, logs)
	})

	r.POST("/count", controllers.CountController)
	r.POST("/logs/:id", controllers.LogIdxController)
	r.POST("/ignore/:id", controllers.IgnoreLogController)
	r.POST("/search", controllers.SearchController)
	r.POST("/rotate", controllers.RotateLogController)
	r.POST("/ignored", controllers.GetIgnoredController)

	r.Run()
}
