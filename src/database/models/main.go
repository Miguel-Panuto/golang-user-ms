package models

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func ExecuteAllIndexes(ctx context.Context, db *mongo.Database) {
	newUserModel(db)
	createIndexes(ctx)
}
