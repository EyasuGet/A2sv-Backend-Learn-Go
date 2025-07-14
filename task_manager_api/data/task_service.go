package data

import (
	"errors"
	"time"
	"github.com/zaahidali/task_manager_api/models"
)


type TaskManager interface{
	GetAllTasks() []models.Task
	GetTaskById(taskid int) error
	CreateTask(task models.Task)
	UpdateTask(taskid int) error
	deleteTask(taskid int)
}


var tasks = []models.Task{
    {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

func GetAllTasks() []models.Task{
	return tasks
}

func GetTaskById(id string) (*models.Task, error){
	
	for i := range tasks{
		if tasks[i].ID == id{
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
	
}

func CreateTask(task models.Task){
	tasks = append(tasks, task)
}

func UpdateTask(id string, updatedTask models.Task) error{
	for i, task := range tasks{
		if task.ID == id{
			tasks[i] = updatedTask
			return nil
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id string) error {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

