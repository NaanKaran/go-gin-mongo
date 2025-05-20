// services/user_service.go
package services

import (
	"errors"
	"go-gin-mongo/models"
	repoInterface "go-gin-mongo/repositories/interface"
	"go-gin-mongo/utils"
)

type UserService struct {
	repo repoInterface.UserRepository
}

func NewUserService(repo repoInterface.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) RegisterUser(user models.User) error {
	existing, _ := s.repo.FindUserByEmail(user.Email)
	if existing.Email != "" {
		return errors.New("email already exists")
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	return s.repo.CreateUser(user)
}

func (s *UserService) LoginUser(email, password string) (*models.User, error) {
	user, err := s.repo.FindUserByEmail(email)
	if err != nil || !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) GetUserById(id string) (*models.User, error) {
	return s.repo.GetUserById(id)
}
