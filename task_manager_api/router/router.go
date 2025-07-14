package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/controllers"
)

func CreateROuter() *gin.Engine{
	router := gin.Default()

	router.GET("/tasks/:id", controllers.GetTask)
	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.CreateTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	return router
}