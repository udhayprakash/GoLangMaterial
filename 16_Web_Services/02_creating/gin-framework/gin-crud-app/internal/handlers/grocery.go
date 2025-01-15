package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GroceryHandler struct {
	// Add service dependencies here
}

func NewGroceryHandler() *GroceryHandler {
	return &GroceryHandler{}
}

func (h *GroceryHandler) GetAllGroceries(c *gin.Context) {
	// Replace with actual logic to fetch groceries
	c.JSON(http.StatusOK, gin.H{"groceries": []string{"Apples", "Bananas"}})
}

func (h *GroceryHandler) GetGroceryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// Replace with actual logic to fetch a grocery by ID
	c.JSON(http.StatusOK, gin.H{"id": id, "name": "Apples"})
}

func (h *GroceryHandler) AddGrocery(c *gin.Context) {
	// Replace with actual logic to add a grocery
	c.JSON(http.StatusOK, gin.H{"message": "Grocery added"})
}

func (h *GroceryHandler) UpdateGrocery(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// Replace with actual logic to update a grocery
	c.JSON(http.StatusOK, gin.H{"message": "Grocery updated", "id": id})
}

func (h *GroceryHandler) DeleteGrocery(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// Replace with actual logic to delete a grocery
	c.JSON(http.StatusOK, gin.H{"message": "Grocery deleted", "id": id})
}
