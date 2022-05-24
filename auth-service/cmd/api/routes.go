package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)
func (app *Config) routes() http.Handler {

	r := chi.NewRouter()

	// handle CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"}, // Use this to allow specific origin hosts
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 86400, // Maximum value not ignored by any of major browsers
	}))
    // A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
    r.Use(middleware.Timeout(60 * time.Second))

	r.Use(middleware.Heartbeat("/ping"))

	r.Post("/authenticate", app.Authenticate)

	return r
}
