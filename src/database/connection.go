package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

var client *mongo.Client

func testConnection() {
	err := client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Start() {
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin@localhost:27017/")
	con, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	client = con

	testConnection()
}

func GetDatabase() *mongo.Database {
	testConnection()
	return client.Database("user-ms")
}

func GetContext() context.Context {
	return ctx
}
