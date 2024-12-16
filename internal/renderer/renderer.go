package renderer

import (
	"fmt"
	"html/template"
	"joshuamURD/testing/internal/config"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig
// NewTemplates sets the config for the template package
func NewTemplates (a *config.AppConfig) {
	app = a
}

// Render is a function that renders templates using html/template
func Render(w http.ResponseWriter, tmpl string) {
	tc, err := CreateTemplateCache()
	if err != nil {
		fmt.Println("Error creating template cache:", err)
		return
	}

	t, ok := tc[tmpl]
	if !ok {
		fmt.Println("Error getting template from cache:", err, tc)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
	
}

// createTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("web/template/*.page.html")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}
		matches, err := filepath.Glob("web/template/*.layout.html")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("web/template/*.layout.html")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = ts
	}

	return cache, nil
}
