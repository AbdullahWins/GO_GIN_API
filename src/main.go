// main.go
package main

import (
	"fmt"
	"os"

	database "crud/src/databases"
	"crud/src/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Initialize the database connection
	err = database.ConnectDatabase()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	fmt.Println("Connected to the database!")

	port := getEnv("PORT", "5000")

	router := routes.SetupMainRoutes()

	router.Run(fmt.Sprintf(":%s", port))
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
