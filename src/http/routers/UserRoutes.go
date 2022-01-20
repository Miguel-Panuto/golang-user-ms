package routers

import (
	"github.com/go-chi/chi"
	"github.com/miguel-panuto/user-ms/src/http/controllers"
)

func NewUserRoutes(r chi.Router) {
	r.Get("/", controllers.IndexUsers)
	r.Post("/", controllers.CreateUser)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", controllers.ShowUser)
		r.Patch("/", controllers.UpdateUser)
	})
}
