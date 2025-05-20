package repositories

import "go-gin-mongo/models"

type UserRepository interface {
	FindUserByEmail(email string) (*models.User, error)
	CreateUser(user models.User) error
	GetAllUsers() ([]models.User, error)
	GetUserById(id string) (*models.User, error)
}
