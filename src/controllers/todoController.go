// controllers/todoController.go
package controllers

import (
	"net/http"
	"strconv"

	"crud/src/models"

	"github.com/gin-gonic/gin"
)

var todos []models.Todo

// GetTodos returns all todos
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// GetTodo returns a specific todo by ID
func GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// CreateTodo creates a new todo
func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}

// UpdateTodo updates an existing todo by ID
func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var existingTodo *models.Todo
	for i, todo := range todos {
		if todo.ID == id {
			existingTodo = &todos[i]
			break
		}
	}

	if existingTodo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var updatedFields map[string]interface{}
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update only the fields provided in the request
	for key, value := range updatedFields {
		switch key {
		case "title":
			existingTodo.Title = value.(string)
		case "status":
			existingTodo.Status = value.(string)
			// Add more fields as needed
		}
	}

	c.JSON(http.StatusOK, existingTodo)
}

// DeleteTodo deletes a todo by ID
func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
