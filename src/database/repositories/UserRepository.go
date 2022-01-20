package repositories

import (
	"github.com/miguel-panuto/user-ms/src/database/models"
	"github.com/miguel-panuto/user-ms/src/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/miguel-panuto/user-ms/src/database"
)

func NewUser(user types.UserDb) (*mongo.InsertOneResult, error) {
	collection := models.GetUserCollection()
	ctx := database.GetContext()
	return collection.InsertOne(ctx, user)
}

func FindUser(filter interface{}) ([]types.UserDb, error) {
	collection := models.GetUserCollection()
	ctx := database.GetContext()
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var users []types.UserDb
	for cursor.Next(ctx) {
		var userCursor bson.M
		var user types.UserDb
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}
		userBytes, _ := bson.Marshal(userCursor)
		bson.Unmarshal(userBytes, &user)
		users = append(users, user)
	}

	return users, nil
}
