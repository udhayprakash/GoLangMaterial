package main

import (
	"gin-crud-app/internal/auth"
	"gin-crud-app/internal/config"
	"gin-crud-app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Gin
	r := gin.Default()

	// Set up authentication middleware
	r.POST("/login", auth.Login)
	authGroup := r.Group("/")
	authGroup.Use(auth.AuthMiddleware())

	// Set up grocery routes
	groceryHandler := handlers.NewGroceryHandler()
	authGroup.GET("/groceries", groceryHandler.GetAllGroceries)
	authGroup.GET("/groceries/:id", groceryHandler.GetGroceryByID)
	authGroup.POST("/groceries", auth.RoleMiddleware("admin"), groceryHandler.AddGrocery)
	authGroup.PUT("/groceries/:id", auth.RoleMiddleware("admin"), groceryHandler.UpdateGrocery)
	authGroup.DELETE("/groceries/:id", auth.RoleMiddleware("admin"), groceryHandler.DeleteGrocery)

	// Start the server
	r.Run(":" + cfg.Port)
}
