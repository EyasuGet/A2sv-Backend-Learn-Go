package repository

import (
	"context"
	"errors"
	"time"

	"github.com/EyasuGet/task-manager-mongo/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct {
    taskCollection *mongo.Collection
}

func NewTaskRepository(col *mongo.Collection) *TaskRepository {
    
    return &TaskRepository{taskCollection: col}
}

func (tr *TaskRepository) GetAll() ([]*domain.Task, error) {
    cursor, err := tr.taskCollection.Find(context.TODO(), bson.D{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    var tasks []*domain.Task
    for cursor.Next(context.TODO()) {
        var taskDoc bson.M
        err := cursor.Decode(&taskDoc)
        if err != nil {
            return nil, err
        }

        id, ok := taskDoc["_id"].(primitive.ObjectID)
        if !ok {
            return nil, errors.New("invalid ID type")
        }

        dueDate, ok := taskDoc["due_date"].(primitive.DateTime)
        if !ok {
            dueDate = primitive.DateTime(time.Now().Unix() * 1000)
        }

        task := &domain.Task{
            ID:          id.Hex(),
            Title:       taskDoc["title"].(string),
            Description: taskDoc["description"].(string),
            DueDate:     dueDate.Time(),
            Status:      taskDoc["status"].(string),
        }
        tasks = append(tasks, task)
    }

    return tasks, nil
}

func (tr *TaskRepository) GetByID(id string) (*domain.Task, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    filter := bson.M{"_id": objectID}
    var taskDoc bson.M
    err = tr.taskCollection.FindOne(context.TODO(), filter).Decode(&taskDoc)
    if err != nil {
        return nil, err
    }

    oid, ok := taskDoc["_id"].(primitive.ObjectID)
    if !ok {
        return nil, errors.New("invalid ID type")
    }

    dueDate, ok := taskDoc["due_date"].(primitive.DateTime)
    if !ok {
        dueDate = primitive.DateTime(time.Now().Unix() * 1000) 
    }

    task := &domain.Task{
        ID:          oid.Hex(),
        Title:       taskDoc["title"].(string),
        Description: taskDoc["description"].(string),
        DueDate:     dueDate.Time(),
        Status:      taskDoc["status"].(string),
    }
    return task, nil
}

func (tr *TaskRepository) Create(task *domain.Task) (*domain.Task, error) {
    newID := primitive.NewObjectID()

    doc := bson.M{
        "_id":         newID,
        "title":       task.Title,
        "description": task.Description,
        "due_date":    task.DueDate,
        "status":      task.Status,
    }

    _, err := tr.taskCollection.InsertOne(context.TODO(), doc)
    if err != nil {
        return nil, err
    }

    task.ID = newID.Hex()
    return task, nil
}


func (tr *TaskRepository) Update(id string, task *domain.Task) (*domain.Task, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

    filter := bson.M{"_id": objectID}
    update := bson.M{
        "$set": bson.M{
            "title":       task.Title,
            "description": task.Description,
            "due_date":    task.DueDate,
            "status":      task.Status,
        },
    }

    _, err = tr.taskCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return nil, err
    }

    updatedTask, err := tr.GetByID(id)
    if err != nil {
        return nil, err
    }
    return updatedTask, nil
}

func (tr *TaskRepository) Delete(id string) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    _, err = tr.taskCollection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
    return err
}