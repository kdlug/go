package mongodb

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var (
	database *mongo.Database
	once     sync.Once
)

func GetClient() (*mongo.Client, error) {
	cfg := GetConfig()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.OperationTimeout)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(cfg.Uri).
		SetConnectTimeout(cfg.ConnectionTimeout)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("connecting to mongo: %w", err)
	}

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("connecting to mongo timed out: %w", err)
		} else {
			return nil, fmt.Errorf("connecting to mongo: %w", err)
		}
	}

	//defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		client.Disconnect(ctx)
		return nil, fmt.Errorf("mongodb ping: %w", err)
	}

	return client, nil
}

func GetDatabase() (*mongo.Database, error) {
	var e error

	once.Do(func() {
		cfg := GetConfig()

		client, err := GetClient()
		if err != nil {
			e = err
			return
		}

		database = client.Database(cfg.Database)
	})

	return database, e
}
