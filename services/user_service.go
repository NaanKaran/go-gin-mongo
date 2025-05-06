package services

import (
	"errors"
	"go-gin-mongo/models"
	"go-gin-mongo/repositories"
	"go-gin-mongo/utils"
)

func RegisterUser(user models.User) error {
	existing, _ := repositories.FindUserByEmail(user.Email)
	if existing.Email != "" {
		return errors.New("email already exists")
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return repositories.CreateUser(user)
}

func LoginUser(email, password string) (*models.User, error) {
	user, err := repositories.FindUserByEmail(email)
	if err != nil || !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

func GetAllUsers() ([]models.User, error) {
	users, err := repositories.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id string) (*models.User, error) {
	user, err := repositories.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
