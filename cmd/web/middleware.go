package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page...", r.URL)
		next.ServeHTTP(w, r)
	})
}

// NoSurve - Adds csrf protection to all POST requests
func NoSurve(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad - It saves and loads the session on every request made
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
