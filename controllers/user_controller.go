package controllers

import (
	"go-gin-mongo/models"
	"go-gin-mongo/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

// @Summary Login User
// @Schemes
// @Description Login User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User Info"
// @Success 200 {object} models.User
// @Failure 401 {object} models.User
// @Router /users/login [post]
func (uc *UserController) Login(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	user, err := uc.service.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "user": user.Email})
}

// @Summary Get All Users
// @Schemes
// @Description Get All Users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary Get User by ID
// @Schemes
// @Description Get User by ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (uc *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := uc.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Register User
// @Schemes
// @Description Register User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "User Info"
// @Success 201 {object} models.User
// @Failure 400 {object} models.User
// @Router /users/register [post]
func (uc *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := uc.service.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
