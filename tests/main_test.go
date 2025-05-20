package tests

import (
	"go-gin-mongo/db"
	"go-gin-mongo/di"
	"go-gin-mongo/logger"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {

	// Initialize logger
	logger.InitLogger()
	defer logger.CloseLogger()

	// Catch panics
	defer func() {
		if r := recover(); r != nil {
			logger.Logger.Println("Panic recovered:", r)
		}
	}()

	// Load .env
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Connect DB
	db.ConnectDB()
	defer db.DisconnectDB()

	// Initialize DI container
	di.InitContainer()

	// defer db.DisconnectDB()
	// Run tests
	os.Exit(m.Run())
}
