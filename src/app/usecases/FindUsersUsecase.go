package usecases

import (
	"encoding/json"
	"fmt"

	"github.com/miguel-panuto/user-ms/src/database/repositories"
	"github.com/miguel-panuto/user-ms/src/types"
)

func FindAllUsers(queries interface{}) ([]byte, error) {
	const fileName = "FindUsersUsecase"
	callName := fmt.Sprintf("%v.Execute()", fileName)

	fmt.Printf("%v - preparing to find users", callName)

	users, err := repositories.FindUser(queries)

	var usersRes []types.UserReponse

	for _, user := range users {
		usersRes = append(usersRes, types.UserReponse{
			Id:   user.Id.Hex(),
			Name: user.Name,
			Age:  user.Age,
			Uuid: user.Uuid,
		})
	}
	resJson, err := json.Marshal(usersRes)
	return resJson, err
}
