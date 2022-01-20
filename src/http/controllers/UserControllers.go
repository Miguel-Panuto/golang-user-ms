package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/miguel-panuto/user-ms/src/app/usecases"
)

const fileName = "UserController"

func IndexUsers(res http.ResponseWriter, req *http.Request) {
	callName := fmt.Sprintf("%v.IndexUsers()", fileName)
	queries := req.URL.Query()
	fmt.Printf("%v - entered with queries: %+v \n", callName, queries)
	users, err := usecases.FindAllUsers(queries)
	if err != nil {
		fmt.Errorf("Error ocored %v", err)
		fmt.Fprintf(res, "{\"error\":\"%v\"}", err)
		return
	}
	fmt.Fprintf(res, "%+v", string(users))
}

func ShowUser(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "showUser")

}

func CreateUser(res http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	repoRes, err := usecases.NewUser(reqBody)
	if err != nil {
		fmt.Printf("Error ocored %v", err)
		fmt.Fprint(res, "{\"error\":\"An error has been ocoured on insert a new user\"}")
		return
	}
	responseJson := fmt.Sprintf("{\"id\":\"%v\"}", repoRes)
	fmt.Fprintf(res, "%v", responseJson)
	fmt.Println("Endpoint Create User")

}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	reqBody, _ := ioutil.ReadAll(req.Body)
	fmt.Fprint(res, string(reqBody))
	fmt.Println("Endpoint Create User")
}
