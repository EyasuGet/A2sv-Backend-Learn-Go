package usecases

import (
	"errors"

	"github.com/EyasuGet/task-manager-mongo/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	userRepo domain.IUserRepo
}

func NewUserUsecase(userRepo domain.IUserRepo) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}
func (uc *UserUsecase) CreateUser(user *domain.User) (string, error) {
	return uc.userRepo.Create(user)
}

func (uc *UserUsecase) Login(email, password string) (*domain.User, error) {
	return uc.userRepo.Login(email, password)
}

func (uc *UserUsecase) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	return uc.userRepo.GetByID(id)
}

func (uc *UserUsecase) DeleteUserByID(id primitive.ObjectID) error{
	return uc.userRepo.DeleteByID(id)
}

func (uc *UserUsecase) PromoteUser(userID primitive.ObjectID, newRole string) error {
	if newRole == "" {
		return errors.New("new role must be provided")
	}
	return uc.userRepo.PromoteUser(userID, newRole)
}