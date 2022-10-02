package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		name := c.Param("name")
		fmt.Printf("%v", name)
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