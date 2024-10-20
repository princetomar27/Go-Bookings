package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func chiRoutes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	//mux.Use(WriteToConsole)

	mux.Use(NoSurve)
	mux.Use(SessionLoad)

	// Register the routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/search-availability", handlers.Repo.Availability)

	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/contact", handlers.Repo.Contact)

	// Get static files to render
	staticFileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static/", staticFileServer))

	return mux
}
