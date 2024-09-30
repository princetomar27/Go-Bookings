package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"bookings/pkg/render"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

var postNum = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// Update to true, when app is in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	// Set some parameters for our cookie
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateCompTemplateCache()
	if err != nil {
		log.Fatal("Got error while creating template cache : ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false
	// set things up for our handlers, by creating a repo variable which calls the function handlers.NewHandlers

	repo := handlers.NewRepo(&app)

	// pass repo back to the handlers
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Server is running on %s", postNum))
	//_ = http.ListenAndServe(postNum, nil) // Start listening to the server

	srv := &http.Server{
		Addr:    postNum,
		Handler: chiRoutes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
