// routes/taskRoutes.go
package routes

import (
	controllers "crud/src/controllers"

	"github.com/gin-gonic/gin"
)

func setupTaskRoutes(router *gin.Engine) {
	taskRouter := router.Group("/tasks")
	{
		taskRouter.GET("", controllers.GetAllTasks)
		taskRouter.GET("/:id", controllers.GetOneTask)
		taskRouter.POST("", controllers.CreateOneTask)
		taskRouter.PATCH("/:id", controllers.UpdateOneTask)
		taskRouter.DELETE("/:id", controllers.DeleteOneTask)
	}
}
