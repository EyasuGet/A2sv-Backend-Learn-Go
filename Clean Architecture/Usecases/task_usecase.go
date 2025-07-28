package usecases

import (
	"errors"

	"github.com/EyasuGet/task-manager-mongo/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecase struct {
    repo domain.ITaskRepository
}

func NewTaskUsecase(repo domain.ITaskRepository) *TaskUsecase {
    return &TaskUsecase{repo: repo}
}
func (tu *TaskUsecase) CreateTask(task *domain.Task) (primitive.ObjectID, error) {
    createdTask, err := tu.repo.Create(task)
    if err != nil {
        return primitive.NilObjectID, err
    }
    objID, err := primitive.ObjectIDFromHex(createdTask.ID)
    if err != nil {
        return primitive.NilObjectID, err
    }
    return objID, nil
}

func (tu *TaskUsecase) GetAllTasks() ([]*domain.Task, error) {
    return tu.repo.GetAll()
}

func (tu *TaskUsecase) GetTaskByID(id string) (*domain.Task, error) {
    _, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, errors.New("invalid id")
    }
    return tu.repo.GetByID(id)
}

func (tu *TaskUsecase) DeleteTaskByID(id string) error {
    _, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return errors.New("invalid id")
    }
    return tu.repo.Delete(id)
}

func (tu *TaskUsecase) UpdateTask(id string, task *domain.Task) error {
    _, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return errors.New("invalid id")
    }
    _, err = tu.repo.Update(id, task)
    return err
}