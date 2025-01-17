package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	
	server.GET("/", func(ctx *gin.Context) {
        ctx.String(200, "Welcome to the Grocery Store!")
	})

	// server.Run("localhost:8000")
	server.Run(":8000")
}

// Logger middleware is added by default
// Also, a middleware that handles any panic, with 500 error

// curl http://localhost:8080/
