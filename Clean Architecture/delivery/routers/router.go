package routers

import (
	"github.com/gin-gonic/gin"

	infrastructure "github.com/EyasuGet/task-manager-mongo/Infrastructure"
	usecases "github.com/EyasuGet/task-manager-mongo/Usecases"
	"github.com/EyasuGet/task-manager-mongo/delivery/controllers"
)

func SetupRouter(userCtrl *controllers.UserController, tasksCtrl *controllers.TaskController, userUsecase *usecases.UserUsecase) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// Public endpoints
		api.POST("/register", userCtrl.Signup)
		api.POST("/login", userCtrl.Login)

		// Protected endpoints
		auth := api.Group("/")
		auth.Use(infrastructure.AuthMiddleware())
		{
			auth.GET("/tasks", tasksCtrl.GetAllTasks)

			// Admin-only endpoints
			admin := auth.Group("/")
			admin.Use(infrastructure.AdminOnly(userUsecase))
			{
				admin.POST("/tasks", tasksCtrl.AddTask)
				admin.PUT("/tasks/:id", tasksCtrl.UpdateTask)
				admin.DELETE("/tasks/:id", tasksCtrl.DeleteTask)
				admin.POST("/promote", userCtrl.PromoteUser)
			}
		}
	}
	return r
}
