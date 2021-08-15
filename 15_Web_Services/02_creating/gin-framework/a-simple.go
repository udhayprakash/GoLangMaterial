package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to Go and Gin!")
	})
	r.Run()
}


// curl http://localhost:8080/