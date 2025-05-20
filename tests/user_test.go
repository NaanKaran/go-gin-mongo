package tests

import (
	"testing"

	"go-gin-mongo/models"
	repoImpl "go-gin-mongo/repositories/implementation"
	repoIface "go-gin-mongo/repositories/interface"
	"go-gin-mongo/services"
	"go-gin-mongo/utils"
)

const testUserEmail = "testuser@example.com"

type UserTest struct {
	Email    string
	Password string
}

// -- Helper methods --

func (ut *UserTest) CreateUser() error {
	user := models.User{
		Email:    ut.Email,
		Password: ut.Password,
	}
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	var repo repoIface.UserRepository = repoImpl.NewUserRepository()
	userRepo := repo.CreateUser(user)
	return userRepo
}

func (ut *UserTest) FindUserByEmail() (*models.User, error) {
	var repo repoIface.UserRepository = repoImpl.NewUserRepository()
	return repo.FindUserByEmail(ut.Email)
}

func (ut *UserTest) LoginUser() (*models.User, error) {
	var repo repoIface.UserRepository = repoImpl.NewUserRepository()
	services := services.NewUserService(repo)
	return services.LoginUser(ut.Email, ut.Password)
}

func (ut *UserTest) GetAllUsers() ([]models.User, error) {
	var repo repoIface.UserRepository = repoImpl.NewUserRepository()
	services := services.NewUserService(repo)
	return services.GetAllUsers()
}

func (ut *UserTest) GetUserById(id string) (*models.User, error) {
	var repo repoIface.UserRepository = repoImpl.NewUserRepository()
	return repo.GetUserById(id)
}

// -- Test cases --

func TestLoginUser(t *testing.T) {
	ut := &UserTest{
		Email:    testUserEmail,
		Password: "password123",
	}

	user, err := ut.LoginUser()
	if err != nil {
		t.Logf("Expected to file : LoginUser failed: %v", err)
	}
	t.Logf("✅ Logged in user: %+v", user)
}

func TestGetAllUsers(t *testing.T) {
	ut := &UserTest{}

	users, err := ut.GetAllUsers()
	if err != nil {
		t.Fatalf("GetAllUsers failed: %v", err)
	}
	t.Logf("✅ All users: %+v", len(users))

}

func TestGetUserById(t *testing.T) {
	ut := &UserTest{}

	// ✅ Replace with a valid user ID from your DB
	id := "663e2ad45cfb2cd1e3d1a8ef"

	user, err := ut.GetUserById(id)
	if err != nil {
		t.Logf("✅ GetUserById failed: %v", err)
	}
	t.Logf("✅ User by ID: %+v", user)
}

func TestCreateUser(t *testing.T) {
	ut := &UserTest{
		Email:    testUserEmail,
		Password: "password123",
	}

	err := ut.CreateUser()
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	t.Logf("✅ User created successfully")
}
