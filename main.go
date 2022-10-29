package main

import (
	"strconv"

	"github.com/Nevin1901/arlog/controllers"
	"github.com/Nevin1901/arlog/models"
	"github.com/Nevin1901/arlog/routines"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDB("test.db")
	go routines.SetupSync()

	store := gormsessions.NewStore(models.DB, true, []byte("Secret"))

	r := gin.Default()
	r.Use(sessions.Sessions("mysession", store))
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int

		v := session.Get("count")

		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}

		session.Set("count", count)
		session.Save()

		c.String(200, strconv.Itoa(count))
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
