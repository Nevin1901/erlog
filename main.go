package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)


type LogRequestBody struct {
	logType string
	title string
	info string
	extraData string
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello")
	})
	r.POST("/", func(c *gin.Context) {
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