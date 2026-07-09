package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// routes builds and returns the application's HTTP router.
// Keeping routing in its own file makes main.go responsible only for process
// startup, while this file owns URL paths and handler registration.
func routes() http.Handler {
	router := chi.NewRouter()

	router.Get("/", home)

	return router
}
