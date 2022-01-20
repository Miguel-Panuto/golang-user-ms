package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type UserDb struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty"`
	Uuid      string             `bson:"uuid,omitempty"`
	Email     string             `bson:"email,omitempty"`
	Age       int                `bson:"age,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	Password  string             `bson:"password,omitempty"`
}

type UserReponse struct {
	Id   string `json:"_id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Uuid string `json:"uuid"`
}
