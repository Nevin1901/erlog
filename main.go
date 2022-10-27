package main

import (
	"github.com/Nevin1901/arlog/controllers"
	"github.com/Nevin1901/arlog/models"
	"github.com/Nevin1901/arlog/routines"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// a super fast databse which basically stores a ton of stuff in memory, and then periodically writes to file when it can (when load is low)
// sqlite is not meant for processing large records, and clickhouse is a pain in the ass to set up. We are not using docker images just to host our server

// TODO: add benchmark
// and use the key value store and periodically append to db
// alternatively I could just create a database in memory, and append to that (benchmark insert times for single records) and then periodically append to larger db every 100 seconds

func main() {
	// var slice []int
	// var wg sync.WaitGroup

	// queue := make(chan int, 1)

	// for i := 0; i < 100; i++ {
	// 	go func(i int) {
	// 		wg.Add(1)
	// 		queue <- i
	// 	}(i)
	// }

	// go func() {
	// 	for t := range queue {
	// 		slice = append(slice, t)
	// 		wg.Done()
	// 	}
	// }()

	// wg.Wait()
	// fmt.Println(slice)

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
