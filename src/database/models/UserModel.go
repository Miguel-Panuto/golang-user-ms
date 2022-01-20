package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func newUserModel(db *mongo.Database) {
	collection = db.Collection("users")
}

func createIndexes(ctx context.Context) {
	collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
}

func GetUserCollection() *mongo.Collection {
	return collection
}
