package main

import (
	"github.com/miguel-panuto/user-ms/src/database"
	"github.com/miguel-panuto/user-ms/src/database/models"
	"github.com/miguel-panuto/user-ms/src/http/routers"
	"github.com/miguel-panuto/user-ms/src/http/server"
)

func main() {
	database.Start()
	r := routers.NewRouter()
	ctx := database.GetContext()
	db := database.GetDatabase()
	models.ExecuteAllIndexes(ctx, db)
	server.New(r)
}
