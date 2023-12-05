// routes/mainRoutes.go
package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupMainRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	// Include other route setups
	setupTodoRoutes(router)
	setupTaskRoutes(router)

	return router
}
