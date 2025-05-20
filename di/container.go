package di

import (
	"go-gin-mongo/controllers"
	repImpl "go-gin-mongo/repositories/implementation"
	repoIface "go-gin-mongo/repositories/interface"
	"go-gin-mongo/services"
)

type Container struct {
	UserController *controllers.UserController
}

func InitContainer() *Container {
	// Repositories
	var userRepo repoIface.UserRepository = repImpl.NewUserRepository()

	// Services
	userService := services.NewUserService(userRepo)

	// Controllers
	userController := controllers.NewUserController(userService)

	return &Container{
		UserController: userController,
	}
}
