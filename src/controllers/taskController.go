// controllers/taskController.go
package controllers

import (
	"net/http"
	"strconv"

	"crud/src/models"

	"github.com/gin-gonic/gin"
)

var tasks []models.Task

// GetTasks returns all tasks
func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, tasks)
}

// GetTask returns a specific task by ID
func GetTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// CreateTask creates a new task
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

// UpdateTask updates an existing task by ID
func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var existingTask *models.Task
	for i, task := range tasks {
		if task.ID == id {
			existingTask = &tasks[i]
			break
		}
	}

	if existingTask == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
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
			existingTask.Title = value.(string)
		case "status":
			existingTask.Status = value.(string)
			// Add more fields as needed
		}
	}

	c.JSON(http.StatusOK, existingTask)
}

// DeleteTask deletes a task by ID
func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
