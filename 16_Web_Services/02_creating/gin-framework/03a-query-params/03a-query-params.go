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
	{ID: 3, Name: "Oranges", Price: 2.49},
	{ID: 4, Name: "Carrots", Price: 0.49},
	{ID: 5, Name: "Potatoes", Price: 0.99},
}

func main() {
	server := gin.Default()

	server.GET("/groceries/", func(ctx *gin.Context) {
		maxPriceStr := ctx.Query("maxPrice") // Extract query parameter

		// If no maxPrice is provided, return all groceries
		if maxPriceStr == "" {
			ctx.JSON(200, groceries)
			return
		}

		// Convert maxPrice to float64
		maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid maxPrice"})
			return
		}

		// Filter groceries with price less than or equal to maxPrice
		var filteredGroceries []Grocery
		for _, item := range groceries {
			if item.Price <= maxPrice {
				filteredGroceries = append(filteredGroceries, item)
			}
		}

		// Return filtered groceries
		ctx.JSON(200, filteredGroceries)

	})
	server.Run()
}

/*
~ curl http://localhost:8080/groceries/
[{"id":1,"name":"Apples","price":1.99},{"id":2,"name":"Bananas","price":0.99},{"id":3,"name":"Oranges","price":2.49},{"id":4,"name":"Carrots","price":0.49},{"id":5,"name":"Potatoes","price":0.99}]

~ curl http://localhost:8080/groceries/maxPrice
404 page not found

~ curl http://localhost:8080/groceries/?maxPrice
[{"id":1,"name":"Apples","price":1.99},{"id":2,"name":"Bananas","price":0.99},{"id":3,"name":"Oranges","price":2.49},{"id":4,"name":"Carrots","price":0.49},{"id":5,"name":"Potatoes","price":0.99}]

~ curl http://localhost:8080/groceries/?maxPrice=1
[{"id":2,"name":"Bananas","price":0.99},{"id":4,"name":"Carrots","price":0.49},{"id":5,"name":"Potatoes","price":0.99}]

~ curl http://localhost:8080/groceries/?maxPrice=0.5
[{"id":4,"name":"Carrots","price":0.49}]


Assignment:

	As a best practice, URLs should be implemented case-insensitive
	fix it in code

	~ curl http://localhost:8080/groceries/?maxprice=0.5
	[{"id":1,"name":"Apples","price":1.99},{"id":2,"name":"Bananas","price":0.99},{"id":3,"name":"Oranges","price":2.49},{"id":4,"name":"Carrots","price":0.49},{"id":5,"name":"Potatoes","price":0.99}]

*/
