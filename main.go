package main

import (
	"github.com/Nevin1901/arlog/controllers"
	"github.com/Nevin1901/arlog/models"
	"github.com/Nevin1901/arlog/routines"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB("test.db")
	go routines.SetupSync()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.POST("/", controllers.AddLogController)
	r.POST("/all", controllers.AllLogsController)
	r.POST("/count", controllers.CountController)
	r.POST("/logs/:id", controllers.LogIdxController)
	r.POST("/ignore/:id", controllers.IgnoreLogController)
	r.POST("/unignore", controllers.UnIgnoreController)
	r.POST("/search", controllers.SearchController)
	r.POST("/rotate", controllers.RotateLogController)
	r.POST("/ignored", controllers.GetIgnoredController)

	// quit <- true to quit
	r.Run()
}
