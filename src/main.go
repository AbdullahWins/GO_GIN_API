// main.go
package main

import (
	"fmt"
	"os"

	"crud/src/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := getEnv("PORT", "5005")

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
