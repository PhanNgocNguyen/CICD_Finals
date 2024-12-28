package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

var items = []Item{
	{ID: "1", Name: "Item 1", Price: "400"},
	{ID: "2", Name: "Item 2", Price: "500"},
	{ID: "3", Name: "Item 3", Price: "700"},
}

func main() {
	router := gin.Default()

	router.GET("/hello", hello)

	router.GET("/items", getItems)
	router.GET("/items/:id", getItemByID)
	router.POST("/items", createItem)
	router.PUT("/items/:id", updateItem)
	router.DELETE("/items/:id", deleteItem)

	router.Run(":4000")
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "It's working"})
}

func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, items)
}

func getItemByID(c *gin.Context) {
	id := c.Param("id")
	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item does not exist"})
}

func createItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err == nil {
		items = append(items, newItem)
		c.JSON(http.StatusCreated, newItem)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func updateItem(c *gin.Context) {
	id := c.Param("id")
	var updatedItem Item
	if err := c.ShouldBindJSON(&updatedItem); err == nil {
		for i, item := range items {
			if item.ID == id {
				items[i].Name = updatedItem.Name
				items[i].Price = updatedItem.Price
				c.JSON(http.StatusOK, items[i])
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func deleteItem(c *gin.Context) {
	id := c.Param("id")
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}
