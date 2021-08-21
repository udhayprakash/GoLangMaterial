package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK!",
		})
	})
	server.Run("localhost:8000")

	gin.SetMode(gin.ReleaseMode)
}

/*
~curl http://localhost:8000/test
{"message":"OK!"}


By default, runs in "debug" mode,

To change to "release" mode,
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)
*/
