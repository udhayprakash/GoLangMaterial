package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/test", func(ctx *gin.Context) {
		// ctx.String(200, "Welcome to the Grocery Store!")
		ctx.XML(200, gin.H{
			"message": "Welcome to the Grocery Store!",
		})
	})
	server.Run("localhost:8000")

	gin.SetMode(gin.ReleaseMode)
}

/*
~ curl http://localhost:8000/test

	<map>
	<message>Welcome to the Grocery Store!</message>
	</map>

*/
