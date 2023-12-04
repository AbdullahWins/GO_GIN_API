package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get the value of the PORT variable or use a default value (5005)
	port := getEnv("PORT", "5005")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		response := Response{Message: "Hello, World!"}
		c.JSON(200, response)
	})

	router.GET("/api", func(c *gin.Context) {
		response := Response{Message: "Hello world from api!"}
		c.JSON(200, response)
	})

	router.Run(fmt.Sprintf(":%s", port))
}

// getEnv retrieves the value of an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
