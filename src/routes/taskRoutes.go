// routes/taskRoutes.go
package routes

import (
	"crud/src/controllers"

	"github.com/gin-gonic/gin"
)

func setupTaskRoutes(router *gin.Engine) {
	taskRouter := router.Group("/tasks")
	{
		taskRouter.GET("", controllers.GetTasks)
		taskRouter.GET("/:id", controllers.GetTask)
		taskRouter.POST("", controllers.CreateTask)
		taskRouter.PATCH("/:id", controllers.UpdateTask)
		taskRouter.DELETE("/:id", controllers.DeleteTask)
	}
}
