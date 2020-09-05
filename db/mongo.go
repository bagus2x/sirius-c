package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	DB *mongo.Database
}

// NewConnection -
var NewConnection = &Connection{}

// Init -
func (c *Connection) Init(uri, name string, time time.Duration) (context.CancelFunc, error) {
	opts := options.Client().ApplyURI(os.Getenv("DB_URI"))
	ctx, cancel := context.WithTimeout(context.Background(), time)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return cancel, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return cancel, err
	}
	fmt.Println("connected to mongo db")
	c.DB = client.Database(name)
	return cancel, nil
}
