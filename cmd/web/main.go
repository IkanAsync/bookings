package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/IkanAsync/bookings/pkg/config"
	"github.com/IkanAsync/bookings/pkg/handlers"
	"github.com/IkanAsync/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplates(&app)

	fmt.Printf("startting application on port %s", portNumber)
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	log.Fatal(server.ListenAndServe())

}
