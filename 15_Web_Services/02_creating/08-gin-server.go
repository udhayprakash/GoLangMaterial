package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a gin router with default middleware
	router := gin.Default()
	// A handler for GET request on /example
	router.GET("/ping1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong 1",
		}) // gin.H is a shortcut for map[string]interface{}
	})

	router.GET("/ping2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong 2",
		}) // gin.H is a shortcut for map[string]interface{}
	})
	router.Run() // listen and serve on port 8080
}

/*
Gin
	- high-performance HTTP web framework
	- allows you to build web applications and microservices
	- contains a set of commonly used functionalities (e.g.,
		routing, middleware support, rendering, etc.) that
		reduce boilerplate code and make it simpler to build
		web applications.

*/
