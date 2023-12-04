package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		response := Response{Message: "Hello, World!"}
		c.JSON(200, response)
	})

	router.GET("/api", func(c *gin.Context) {
		response := Response{Message: "Hello world from api!"}
		c.JSON(200, response)
	})

	port := 5005
	router.Run(fmt.Sprintf(":%d", port))
}
