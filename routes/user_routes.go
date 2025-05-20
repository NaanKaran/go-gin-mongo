package routes

import (
	"go-gin-mongo/di"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {

	container := di.InitContainer()

	userGroup := router.Group("/api/users")
	{
		userGroup.POST("/register", container.UserController.Register)
		userGroup.POST("/login", container.UserController.Login)
		userGroup.GET("/", container.UserController.GetAllUsers)
		userGroup.GET("/:id", container.UserController.GetUserById)
	}
}
