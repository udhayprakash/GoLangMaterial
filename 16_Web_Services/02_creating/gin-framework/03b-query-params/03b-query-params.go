package main

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Grocery represents a grocery item
type Grocery2 struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

var groceries = []Grocery2{
	{ID: 1, Name: "Apples", Price: 1.99, Category: "fruits"},
	{ID: 2, Name: "Bananas", Price: 0.99, Category: "fruits"},
	{ID: 3, Name: "Oranges", Price: 2.49, Category: "fruits"},
	{ID: 4, Name: "Carrots", Price: 0.49, Category: "vegetables"},
	{ID: 5, Name: "Potatoes", Price: 0.99, Category: "vegetables"},
	{ID: 6, Name: "Rice", Price: 1.29, Category: "cereal"},
}

func main() {
	server := gin.Default()

	// Endpoint to filter groceries
	server.GET("/groceries/", func(ctx *gin.Context) {
		// Extract query parameters
		minPriceStr := ctx.Query("minPrice")
		maxPriceStr := ctx.Query("maxPrice")
		category := strings.ToLower(ctx.Query("category")) // Convert category to lowercase

		// If no filters are provided, return all groceries
		if minPriceStr == "" && maxPriceStr == "" && category == "" {
			ctx.JSON(200, groceries)
			return
		}

		var filteredGroceries []Grocery2

		// Convert minPrice and maxPrice to float64
		var minPrice, maxPrice float64
		var err error

		if minPriceStr != "" {
			minPrice, err = strconv.ParseFloat(minPriceStr, 64)
			if err != nil {
				ctx.JSON(400, gin.H{"error": "Invalid minPrice"})
				return
			}
		}

		if maxPriceStr != "" {
			maxPrice, err = strconv.ParseFloat(maxPriceStr, 64)
			if err != nil {
				ctx.JSON(400, gin.H{"error": "Invalid maxPrice"})
				return
			}
		}

		// Filter groceries
		for _, item := range groceries {
			// Check price range
			priceInRange := true
			if minPriceStr != "" && item.Price < minPrice {
				priceInRange = false
			}
			if maxPriceStr != "" && item.Price > maxPrice {
				priceInRange = false
			}

			// Check category (case-insensitive)
			categoryMatch := true
			if category != "" && strings.ToLower(item.Category) != category {
				categoryMatch = false
			}

			// Add to filtered list if both conditions are met
			if priceInRange && categoryMatch {
				filteredGroceries = append(filteredGroceries, item)
			}
		}

		// Return filtered groceries
		ctx.JSON(200, filteredGroceries)
	})

	server.Run() // Listen and serve on http://localhost:8080
}

/*

$ curl http://localhost:8080/groceries/
[{"id":1,"name":"Apples","price":1.99,"category":"fruits"},{"id":2,"name":"Bananas","price":0.99,"category":"fruits"},{"id":3,"name":"Oranges","price":2.49,"category":"fruits"},{"id":4,"name":"Carrots","price":0.49,"category":"vegetables"},{"id":5,"name":"Potatoes","price":0.99,"category":"vegetables"},{"id":6,"name":"Rice","price":1.29,"category":"cereal"}]

$ curl http://localhost:8080/groceries/?minPrice=1.00
[{"id":1,"name":"Apples","price":1.99,"category":"fruits"},{"id":3,"name":"Oranges","price":2.49,"category":"fruits"},{"id":6,"name":"Rice","price":1.29,"category":"cereal"}]

$ curl http://localhost:8080/groceries/?maxPrice=0.5
[{"id":4,"name":"Carrots","price":0.49,"category":"vegetables"}]

// invalid input
$ curl http://localhost:8080/groceries/?maxPrice=0.5&minPrice=1.00
null

$ curl http://localhost:8080/groceries/?maxPrice=1&minPrice=0.5
[{"id":2,"name":"Bananas","price":0.99,"category":"fruits"},{"id":5,"name":"Potatoes","price":0.99,"category":"vegetables"}]

$ curl http://localhost:8080/groceries/?maxPrice=1&minPrice=0.5&category=fruits
[{"id":2,"name":"Bananas","price":0.99,"category":"fruits"}]
*/