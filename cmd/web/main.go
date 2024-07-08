package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mistupustu/Bookings/pkg/config"
	"github.com/mistupustu/Bookings/pkg/handlers"
	"github.com/mistupustu/Bookings/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager
// Main is the main application function
func main() {
	app.InProduction = false // Set to true in production

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCashe, err := render.CreateTemplateCashe()
	if err != nil {
		log.Fatal("cannot create template cashe, err:", err)
	}
	app.TemplateCashe = templateCashe
	app.UseChash = false
	render.NewTemplate(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting application on port %s...\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("listenAndServe error: ", err)
	}
}
