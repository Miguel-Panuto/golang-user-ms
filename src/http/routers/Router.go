package routers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/user", NewUserRoutes)

	return r
}
