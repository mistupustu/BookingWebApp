package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the application config
type AppConfig struct {
	UseChash      bool                          //UseChache enables or disables template caching (Development or Production mode)
	TemplateCashe map[string]*template.Template //TemplateCashe is a map that holds parsed templates for the application
	InProduction  bool                          //InProduction Indicates whether the application is running in production mode
	InfoLog       *log.Logger
	Session       *scs.SessionManager //Session is the session manager for the application
}
