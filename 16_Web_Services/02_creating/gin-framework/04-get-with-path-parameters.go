package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var todos []string

func Lists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": todos})
}

func ListItem(c *gin.Context) {
	errormessage := "Index out of range"
	indexstring := c.Param("index")
	if index, err := strconv.Atoi(indexstring); err == nil && index < len(todos) {
		c.JSON(http.StatusOK, gin.H{"item": todos[index]})
	} else {
		if err != nil {
			errormessage = "Number expected: " + indexstring
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errormessage})
	}
}

func main() {
	todos = append(todos, "Write the application")
	r := gin.Default()
	r.GET("/api/lists", Lists)
	r.GET("/api/lists/:index", ListItem)
	r.Run()
}

/*
Usage:
	curl -i http://localhost:8080/api/lists/0
	curl -i http://localhost:8080/api/lists/1
	curl -i http://localhost:8080/api/lists/foo
*/
