package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Completed   bool
	DueDate   time.Time
	Status    string
}

type User struct {
	ID       string
	Email    string
	Password string
	Role     string
}

type IUserRepo interface {
	Create(user *User) (string, error)
	Login(userEmail, password string) (*User, error)
	GetByID(primitive.ObjectID) (*User, error)
	DeleteByID(id primitive.ObjectID) error
	PromoteUser(userID primitive.ObjectID, newRole string) error
}


type ITaskRepository interface {
    GetAll() ([]*Task, error)
    GetByID(id string) (*Task, error)
    Create(task *Task) (*Task, error)
    Update(id string, task *Task) (*Task, error)
    Delete(id string) error
}