package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func New(r chi.Router) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"check\": \"api up and running\"}")
	})
	log.Fatal(http.ListenAndServe(":3000", r))
}
