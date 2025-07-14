package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/zaahidali/task_manager_api/data"
	"github.com/zaahidali/task_manager_api/models"
)



func GetTasks(c *gin.Context){
	tasks := data.GetAllTasks()
	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context){
	id := c.Param("id")
	task, error := data.GetTaskById(id)
	if error != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context){
	var newTask models.Task
	if error := c.ShouldBindJSON(&newTask); error != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}
	if newTask.ID == "" || newTask.Title == "" || newTask.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	data.CreateTask(newTask)
	c.JSON(http.StatusCreated, gin.H{"message": "task created"})
}

func UpdateTask(c *gin.Context){
	id := c.Param("id")

	var updatedTask  models.Task

	if error := c.ShouldBindJSON(&updatedTask); error != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return 
	}
	

	err := data.UpdateTask(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, updatedTask)

}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}