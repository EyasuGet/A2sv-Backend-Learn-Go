package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/EyasuGet/task-manager-mongo/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	database   mongo.Database
	collection *mongo.Collection
}

func NewUserRepo(db *mongo.Database, collectionName string) *UserRepo {
	return &UserRepo{
		database:   *db,
		collection: db.Collection(collectionName),
	}
}

func initDB(username, password string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client.Database("taskmanager").Collection("tasks")

}

// func (ur *UserRepo) CheckUserExists(user domain.User) (bool, domain.User) {
// 	var result domain.User
// 	userCollection := ur.Collection(ur.collection)

// 	err := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&result)
// 	return err == nil, result
// }

func (ur *UserRepo) Create(user *domain.User) (string, error) {

	userCollection := ur.collection

	// Generate a new ID for the user
	objectID := primitive.NewObjectID()
	user.ID = objectID.Hex()
	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return "user successfully registered", nil

}

func (ur *UserRepo) Login(usernameOrEmail, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	filter := bson.M{
		"$or": []bson.M{
			{"username": usernameOrEmail},
			{"email": usernameOrEmail},
		},
	}
	err := ur.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func (ur *UserRepo) GetByID(id primitive.ObjectID) (*domain.User, error) {

	var user domain.User
	err := ur.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepo) DeleteByID(id primitive.ObjectID) error {
	result, err := ur.collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (ur *UserRepo) PromoteUser(userID primitive.ObjectID, newRole string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{"role": newRole}}
	result, err := ur.collection.UpdateByID(ctx, userID, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
