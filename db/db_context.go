package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBContext struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Users      *mongo.Collection
	Products   *mongo.Collection
	Orders     *mongo.Collection
	Ctx        context.Context
	CancelFunc context.CancelFunc
}

func NewDBContext(uri string, dbName string) (*DBContext, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		cancel()
		return nil, err
	}

	// Check connection
	if err := client.Ping(ctx, nil); err != nil {
		cancel()
		return nil, err
	}

	db := client.Database(dbName)

	return &DBContext{
		Client:     client,
		Database:   db,
		Users:      db.Collection("users"),
		Products:   db.Collection("products"),
		Orders:     db.Collection("orders"),
		Ctx:        ctx,
		CancelFunc: cancel,
	}, nil
}

func (dbc *DBContext) Close() {
	dbc.CancelFunc()
	_ = dbc.Client.Disconnect(dbc.Ctx)
}
