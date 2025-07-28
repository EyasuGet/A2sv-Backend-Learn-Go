package controllers

import (
	"net/http"

	infrastructure "github.com/EyasuGet/task-manager-mongo/Infrastructure"
	usecases "github.com/EyasuGet/task-manager-mongo/Usecases"
	"github.com/EyasuGet/task-manager-mongo/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	UserUsecase *usecases.UserUsecase
}

func NewUserController(uc *usecases.UserUsecase) *UserController {
	return &UserController{UserUsecase: uc}
}

func (uc *UserController) Signup(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := uc.UserUsecase.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "user created", "id": id})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserUsecase.Login(input.Username, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	userObjID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user ID"})
		return
	}
	token, err := infrastructure.GenerateJWT(userObjID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) PromoteUser(ctx *gin.Context) {
	var input struct {
		UserID  string `json:"user_id"`
		NewRole string `json:"new_role"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	userID, err := primitive.ObjectIDFromHex(input.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	err = uc.UserUsecase.PromoteUser(userID, input.NewRole)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user role updated", "new_role": input.NewRole})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user ID required"})
		return
	}
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	err = uc.UserUsecase.DeleteUserByID(userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}