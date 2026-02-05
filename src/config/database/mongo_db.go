package database

import (
	"context"
	"golang_crud/src/plugin/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func MongoConnect(mongoURI string, dbName string) *mongo.Database {
	if mongoURI == "" || dbName == "" {
		logger.Log.Error("mongo config missing")
		panic("mongo config missing")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(mongoURI).
		SetServerSelectionTimeout(5 * time.Second)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Log.Error("Mongo connect failed", zap.Error(err))
		panic("PANIC - Mongo connect failed")
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Log.Error("Mongo ping failed", zap.Error(err))
		panic("PANIC - Mongo ping failed")
	}

	logger.Log.Info("Mongodb connected successfully", zap.String("database", dbName))

	return client.Database(dbName)
}
