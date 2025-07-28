package main

import (
	"log"

	infrastructure "github.com/EyasuGet/task-manager-mongo/Infrastructure"
	repositories "github.com/EyasuGet/task-manager-mongo/Repositories"
	usecases "github.com/EyasuGet/task-manager-mongo/Usecases"
	"github.com/EyasuGet/task-manager-mongo/delivery/controllers"
	"github.com/EyasuGet/task-manager-mongo/delivery/routers"
)

func main() {
	client, err := infrastructure.ConnectMongo("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	db := client.Database("taskmanager")
	taskCollection := db.Collection("tasks")

	taskRepo := repositories.NewTaskRepository(taskCollection)
	userRepo := repositories.NewUserRepo(db, "users")

	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo)

	taskController := controllers.NewTaskController(taskUsecase)
	userController := controllers.NewUserController(userUsecase)

	router := routers.SetupRouter(userController, taskController, userUsecase)

	log.Println("Server running on :8080")
	router.Run(":8080")
}