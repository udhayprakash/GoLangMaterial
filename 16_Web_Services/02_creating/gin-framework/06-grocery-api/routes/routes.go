package routes

import (
    "github.com/gin-gonic/gin"
    "grocery-api/handlers"
)

func SetupGroceryRoutes(r *gin.Engine) {
    r.GET("/groceries", handlers.GetGroceries)
    r.GET("/groceries/:id", handlers.GetGroceryByID)
    r.POST("/groceries", handlers.AddGrocery)
    r.PUT("/groceries/:id", handlers.UpdateGrocery)
    r.DELETE("/groceries/:id", handlers.DeleteGrocery)
}