package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	groceries := []string{"Apples", "Bananas", "Oranges"}

	server.GET("/groceries/:index", func(ctx *gin.Context) {
		index, _ := strconv.Atoi(ctx.Param("index"))
		if index < 0 || index >= len(groceries) {
			ctx.JSON(404, gin.H{"error": "Item not found"})
			return
		}
		ctx.JSON(200, gin.H{
			"groceries": groceries[index],
		})
	})
	server.Run() // default port is 8080
}

/*
~ curl http://localhost:8080/groceries:1

{"groceries":"Apples"}

*/
