// controllers/taskController.go
package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	database "crud/src/databases"
	"crud/src/models"
)

var tasksCollection *mongo.Collection

func init() {
	// Assuming you've connected to the database elsewhere in your code
	// and obtained the 'tasks' collection
	// For example, you might do this in your main function
	// database.ConnectDatabase()
	tasksCollection = database.TasksCollection
}

// GetTasks returns all tasks
func GetAllTasks(c *gin.Context) {
	// Check if tasksCollection is nil
	if tasksCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Collection not found!"})
		return
	}

	// Use tasksCollection to fetch tasks from MongoDB
	cursor, err := tasksCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tasks not found!"})
		return
	}
	defer cursor.Close(context.Background())

	var tasks []models.Task
	if err := cursor.All(context.Background(), &tasks); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error!"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetOneTask returns a specific task by ID
func GetOneTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Check if tasksCollection is nil
	if tasksCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var task models.Task
	err = tasksCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// CreateTask creates a new task
func CreateOneTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure that tasksCollection is not nil before using it
	if database.TasksCollection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	_, err := database.TasksCollection.InsertOne(context.Background(), newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

// UpdateTask updates an existing task by ID
func UpdateOneTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedFields map[string]interface{}
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	update := bson.M{"$set": updatedFields}
	// Find and update the task
	err = tasksCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, update).Decode(&updatedFields)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, updatedFields)
}

// DeleteTask deletes a task by ID
func DeleteOneTask(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	_, err = tasksCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
