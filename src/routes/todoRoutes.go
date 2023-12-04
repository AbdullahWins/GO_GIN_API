// routes/todoRoutes.go
package routes

import (
	"crud/src/controllers"

	"github.com/gin-gonic/gin"
)

func setupTodoRoutes(router *gin.Engine) {
	todoRouter := router.Group("/todos")
	{
		todoRouter.GET("", controllers.GetTodos)
		todoRouter.GET("/:id", controllers.GetTodo)
		todoRouter.POST("", controllers.CreateTodo)
		todoRouter.PATCH("/:id", controllers.UpdateTodo)
		todoRouter.DELETE("/:id", controllers.DeleteTodo)
	}
}
