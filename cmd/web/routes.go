package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()
	router.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Print("Hi")
	})
	return router
}
