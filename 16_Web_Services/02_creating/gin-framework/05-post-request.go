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

func AddListItem(c *gin.Context) {
	item := c.PostForm("item")
	todos = append(todos, item)
	c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(len(todos)-1))
}

func main() {
	todos = append(todos, "Write the application")
	r := gin.Default()
	r.GET("/api/lists", Lists)
	r.GET("/api/lists/:index", ListItem)
	r.POST("/api/lists", AddListItem)
	r.Run()
}

/*
Usage:
	curl -i -X POST -d "item=apple" http://localhost:8080/api/lists
	curl -i http://localhost:8080/api/lists

	curl -i -X POST -d "item=mango" http://localhost:8080/api/lists
	curl -i -X POST -d "item=banana" http://localhost:8080/api/lists
	curl -i http://localhost:8080/api/lists

*/
