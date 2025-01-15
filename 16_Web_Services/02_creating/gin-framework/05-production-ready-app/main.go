package main

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "os"
    "strconv"
)

type Grocery struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

var (
    groceries []Grocery
    log       *logrus.Logger
)

// init function for initializing global configurations
func init() {
    // Set Gin to release mode
    gin.SetMode(gin.ReleaseMode)

    // Initialize logger
    log = logrus.New()
    log.Out = os.Stdout
    log.SetFormatter(&logrus.JSONFormatter{})

    // Initialize groceries (optional, if you want to pre-populate data)
    groceries = []Grocery{
        {ID: 1, Name: "Apples", Price: 1.99},
        {ID: 2, Name: "Bananas", Price: 0.99},
    }
}

func main() {
    r := gin.New()
    r.Use(gin.LoggerWithWriter(log.Out))
    r.Use(gin.Recovery())

    // Health check endpoint
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

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

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Info("Starting server on :" + port)
    r.Run(":" + port)
}

/*
NOTE: 
	Avoid reaidng environment variables (PORT, etc) from init().
	place in main(), as it depends on runtime conditions
	
	overusing init() can lead to reliance on global state, which makes
	testing and debugging harder

*/