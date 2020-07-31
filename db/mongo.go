package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect -
func Connect(uri, name string, time time.Duration) (*mongo.Database, context.CancelFunc, error) {
	opts := options.Client().ApplyURI(os.Getenv("DB_URI"))
	ctx, cancel := context.WithTimeout(context.Background(), time)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, cancel, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, cancel, err
	}
	fmt.Println("connected to mongo db")
	return client.Database(name), cancel, nil
}
