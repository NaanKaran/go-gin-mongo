package db

import (
	"context"
	"go-gin-mongo/logger"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	databaseURI := os.Getenv("DATABASE_URL")
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseURI == "" || databaseName == "" {
		logger.Logger.Fatal("Database URI or name not provided in environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(databaseURI))
	if err != nil {
		logger.Logger.Fatal("MongoDB connection error:", err)
	}

	// Check the connection
	if err := client.Ping(ctx, nil); err != nil {
		logger.Logger.Fatal("MongoDB ping error:", err)
	}

	DB = client.Database(databaseName)
	logger.Logger.Println("✅ Connected to MongoDB successfully")
}

func GetCollection(name string) *mongo.Collection {
	return DB.Collection(name)
}
