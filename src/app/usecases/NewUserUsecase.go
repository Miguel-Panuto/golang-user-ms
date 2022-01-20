package usecases

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/miguel-panuto/user-ms/src/database/repositories"
	"github.com/miguel-panuto/user-ms/src/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(userBytes []byte) (string, error) {
	const fileName = "NewUserUsecase"
	callName := fmt.Sprintf("%v.Execute()", fileName)

	var userEntry types.User
	err := json.Unmarshal(userBytes, &userEntry)
	if err != nil {
		return "", err
	}

	fmt.Printf("%v - entered with user: %+v \n", callName, userEntry)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userEntry.Password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	id := primitive.NewObjectID()
	user := types.UserDb{
		Id:        id,
		Name:      userEntry.Name,
		Uuid:      uuid.NewString(),
		Email:     userEntry.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Age:       userEntry.Age,
		Password:  string(hashedPassword),
	}

	_, err = repositories.NewUser(user)

	return id.Hex(), err
}
