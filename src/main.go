package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
}

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos []Todo

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := getEnv("PORT", "5005")

	router := gin.Default()

	// Routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, Response{Message: "Hello, World!"})
	})

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, Response{Message: "Hello world from api!"})
	})

	// CRUD operations
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.POST("/todos", createTodo)
	router.PATCH("/todos/:id", updateTodo)
	router.DELETE("/todos/:id", deleteTodo)

	router.Run(fmt.Sprintf(":%s", port))
}

func getTodos(c *gin.Context) {
	c.JSON(200, todos)
}

func getTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	for _, todo := range todos {
		if todo.ID == id {
			c.JSON(200, todo)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Todo not found"})
}

func createTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	newTodo.ID = len(todos) + 1
	todos = append(todos, newTodo)

	c.JSON(201, newTodo)
}

func updateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	var existingTodo *Todo
	for i, todo := range todos {
		if todo.ID == id {
			existingTodo = &todos[i]
			break
		}
	}

	if existingTodo == nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	var updatedFields map[string]interface{}
	if err := c.ShouldBindJSON(&updatedFields); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
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

	c.JSON(200, existingTodo)
}

func deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(200, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Todo not found"})
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
