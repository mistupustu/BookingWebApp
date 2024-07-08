package render

import (
	"bytes"
	"github.com/mistupustu/Bookings/pkg/config"
	"github.com/mistupustu/Bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// NewTemplate sets the config for the template package
func NewTemplate(appConfig *config.AppConfig) {
	app = appConfig
}
//AddDefaultData adds default data to the templateData
func AddDefaultData(templateData *models.TemplateData) *models.TemplateData{
	//As of right now, no default data is added here. You can add more as needed.
	return templateData
}
// RenderTemplate renders a given template using html/template.
func RenderTemplate(w http.ResponseWriter, templateName string, templateData *models.TemplateData) {
	var templateCashe map[string]*template.Template
	var err error
	if app.UseChash {
		templateCashe = app.TemplateCashe
	} else {
		templateCashe, err = CreateTemplateCashe()
		if err != nil {
			log.Fatal(err)
		}
	}
	//Get the template cashe from app config

	// Get requested template from cashe
	template, ok := templateCashe[templateName]
	if !ok {
		log.Fatal("could not get template from cache")
	}

	buffer := new(bytes.Buffer)
	templateData = AddDefaultData(templateData)
	err = template.Execute(buffer, templateData)
	if err != nil {
		log.Fatal(err)
	}
	// Render the template
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCashe() (map[string]*template.Template, error) {
	//Create empty cashe
	myCashe := map[string]*template.Template{}

	// Get all pages from templates directory
	pages, err := filepath.Glob("./templates/*page.gohtml")
	if err != nil {
		return myCashe, err
	}
	//Range through pages
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCashe, err
		}
		matches, err := filepath.Glob("./templates/*layout.gohtml")
		if err != nil {
			return myCashe, err
		}
		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*layout.gohtml")
			if err != nil {
				return myCashe, err
			}
		}

		myCashe[name] = templateSet // Add parsed template to cashe
	}
	return myCashe, nil // Return the cashe with all parsed templates
}
