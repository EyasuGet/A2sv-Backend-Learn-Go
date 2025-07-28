package infrastructure

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{

}

func NewPasswordService() *PasswordService{
	return &PasswordService{}
}

func (ps *PasswordService) HashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (ps *PasswordService) CheckPasswordHash(password string, hash string) error {
	return  bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	
}

