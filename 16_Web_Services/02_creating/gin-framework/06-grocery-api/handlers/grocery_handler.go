package handlers

import (
    "github.com/gin-gonic/gin"
    "strconv"
    "grocery-api/models"
)

var groceries = []models.Grocery{
    {ID: 1, Name: "Apples", Price: 1.99},
    {ID: 2, Name: "Bananas", Price: 0.99},
}

func GetGroceries(c *gin.Context) {
    c.JSON(200, groceries)
}

func GetGroceryByID(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for _, item := range groceries {
        if item.ID == id {
            c.JSON(200, item)
            return
        }
    }
    c.JSON(404, gin.H{"error": "Item not found"})
}

func AddGrocery(c *gin.Context) {
    var newItem models.Grocery
    if err := c.ShouldBindJSON(&newItem); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }
    groceries = append(groceries, newItem)
    c.JSON(200, gin.H{"message": "Item added", "groceries": groceries})
}

func UpdateGrocery(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var updatedItem models.Grocery
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
}

func DeleteGrocery(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    for i, item := range groceries {
        if item.ID == id {
            groceries = append(groceries[:i], groceries[i+1:]...)
            c.JSON(200, gin.H{"message": "Item deleted", "groceries": groceries})
            return
        }
    }
    c.JSON(404, gin.H{"error": "Item not found"})
}