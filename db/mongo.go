package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect -
func Connect(uri, name string, time time.Duration) (*mongo.Database, context.CancelFunc, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), time)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, cancel, err
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, cancel, err
	}
	fmt.Println("connected to mongo db")
	return client.Database(name), cancel, nil
}
