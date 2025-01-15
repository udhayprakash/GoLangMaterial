package handlers_test

import (
	"gin-crud-app/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllGroceries(t *testing.T) {
	// Setup
	r := gin.Default()
	handler := handlers.NewGroceryHandler()
	r.GET("/groceries", handler.GetAllGroceries)

	// Test request
	req, _ := http.NewRequest("GET", "/groceries", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Apples")
}
