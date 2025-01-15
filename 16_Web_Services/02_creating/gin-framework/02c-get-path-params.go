package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Grocery struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var groceries = []Grocery{
	{ID: 1, Name: "Apples", Price: 1.99},
	{ID: 2, Name: "Bananas", Price: 0.99},
}

func main() {
	server := gin.Default()

	server.GET("/groceries/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id")) // Extract path parameter
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid ID"})
			// fmt.Sprintf("Invalid ID: %v", id),
			// fmt.Sprintf("Invalid ID: %v; Error is %v", id, err),

			return
		}

		for _, item := range groceries {
			if item.ID == id {
				ctx.JSON(200, item)
				return
			}
		}
		ctx.JSON(404, gin.H{"error": "Item not found"})

	})
	server.Run() // default port is 8080
}

/*
~ curl http://localhost:8080/groceries/1
{"id":1,"name":"Apples","price":1.99}

~ curl http://localhost:8080/groceries/2
{"id":2,"name":"Bananas","price":0.99}

~ curl http://localhost:8080/groceries/3
{"error":"Item not found"}

~ curl http://localhost:8080/groceries/apple
{"error":"Invalid ID: 0"}*/
