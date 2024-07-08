package handlers

import (
	"github.com/mistupustu/Bookings/pkg/config"
	"github.com/mistupustu/Bookings/pkg/render"
	"github.com/mistupustu/Bookings/pkg/models"
	"net/http"
)


// Repo is the global repository for the handlers package.  It's a pointer to a Repository struct.
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository with the provided app config.
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// NewHandlers sets the repository for the handlers package.
func NewHandlers(repo *Repository) {
	Repo = repo
}

// Home is the Home page for the application
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{})
}

// About handles GET requests to /about and returns the about page.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"test": "Hello again!",
	}

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
