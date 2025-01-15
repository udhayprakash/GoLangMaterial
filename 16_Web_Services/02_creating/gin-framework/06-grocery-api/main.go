package main

import (
	"grocery-api/middlewares"
	"grocery-api/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	// Initialize any global configurations or variables here
}

func main() {
	r := gin.Default()

	// Apply middlewares
	r.Use(middlewares.LoggerMiddleware())
	r.Use(middlewares.AuthMiddleware())
	r.Use(middlewares.CORSMiddleware())

	// Setup routes
	routes.SetupGroceryRoutes(r)

	// Start the server
	r.Run()
}
