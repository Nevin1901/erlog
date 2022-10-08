package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

// a super fast databse which basically stores a ton of stuff in memory, and then periodically writes to file when it can (when load is low)
// sqlite is not meant for processing large records, and clickhouse is a pain in the ass to set up. We are not using docker images just to host our server

type LogRequestBody struct {
	gorm.Model
	logType		string	`gorm:"serializer:json"`
	title		string	`gorm:"serializer:json"`
	info		string	`gorm:"serializer:json"`
	extraData	string	`gorm:"serializer:json"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	db.AutoMigrate(&LogRequestBody{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello")
	})

	r.POST("/", func(c *gin.Context) {
		logObj := LogRequestBody{}

		errLog := c.ShouldBindBodyWith(&logObj, binding.JSON);

		if errLog != nil {
			println("Failed binding body")
			c.String(http.StatusBadRequest, "Failed binding body")
			return
		}

		fmt.Printf("%+v\n", logObj)
		db.Create(&logObj)
		c.String(200, "OK")
		println("done");
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
		logs := LogRequestBody{};
		result := db.Find(&logs)
		println(result)
		fmt.Printf("%+v\n", logs)
		c.String(200, "ok")
	})

	r.GET("/:name", func(c *gin.Context) {
		name := c.Param("name")
		println("hello " + name)
		message := "hello " + name
		c.String(http.StatusOK, message)
	})

	r.Run()
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello you've requested root")
	// })

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.ListenAndServe(":8000", nil)
}