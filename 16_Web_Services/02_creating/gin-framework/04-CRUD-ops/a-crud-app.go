package main

import (
    "github.com/gin-gonic/gin"
    "strconv"
)

type Grocery struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

var groceries = []Grocery{
    {ID: 1, Name: "Apples", Price: 1.99},
    {ID: 2, Name: "Bananas", Price: 0.99},
}

func main() {
    r := gin.Default()

    // Get all groceries
    r.GET("/groceries", func(c *gin.Context) {
        c.JSON(200, groceries)
    })

    // Get a specific grocery by ID
    r.GET("/groceries/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        for _, item := range groceries {
            if item.ID == id {
                c.JSON(200, item)
                return
            }
        }
        c.JSON(404, gin.H{"error": "Item not found"})
    })

    // Add a new grocery
    r.POST("/groceries", func(c *gin.Context) {
        var newItem Grocery
        if err := c.ShouldBindJSON(&newItem); err != nil {
            c.JSON(400, gin.H{"error": "Invalid input"})
            return
        }
        groceries = append(groceries, newItem)
        c.JSON(200, gin.H{"message": "Item added", "groceries": groceries})
    })

    // Update a grocery
    r.PUT("/groceries/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        var updatedItem Grocery
        if err := c.ShouldBindJSON(&updatedItem); err != nil {
            c.JSON(400, gin.H{"error": "Invalid input"})
            return
        }
        for i, item := range groceries {
            if item.ID == id {
                groceries[i] = updatedItem
                c.JSON(200, gin.H{"message": "Item updated", "groceries": groceries})
                return
            }
        }
        c.JSON(404, gin.H{"error": "Item not found"})
    })

    // Delete a grocery
    r.DELETE("/groceries/:id", func(c *gin.Context) {
        id, _ := strconv.Atoi(c.Param("id"))
        for i, item := range groceries {
            if item.ID == id {
                groceries = append(groceries[:i], groceries[i+1:]...)
                c.JSON(200, gin.H{"message": "Item deleted", "groceries": groceries})
                return
            }
        }
        c.JSON(404, gin.H{"error": "Item not found"})
    })

    r.Run()
}


/*

[GIN-debug] GET    /groceries      --> main.main.func1 (3 handlers)

(3 handlers) means:

	Logger Middleware
	Recovery Middleware
	Your custom route handler

*/