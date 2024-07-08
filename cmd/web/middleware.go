package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// NoSurf middleware protects against CSRF attacks by checking the X-Surf-Token header.
func NoSurfe(next http.Handler) http.Handler {
	csrvHandler := nosurf.New(next)

	csrvHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Secure:   app.InProduction, // Set to true in production
		SameSite: http.SameSiteLaxMode,
	})

	return csrvHandler
}

// SessionLoad middleware loads and saves the session on every request.
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
