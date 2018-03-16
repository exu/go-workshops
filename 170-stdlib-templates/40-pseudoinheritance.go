package main

import (
	"html/template"
	"os"
)

type templateMap map[string]*template.Template

func main() {

	// let's define some templates which will be inherited
	// from base.html template
	// ParseFiles will receive those files
	templates := templateMap{
		"index.html": template.Must(template.ParseFiles("view/index.html", "view/base.html")),
		"other.html": template.Must(template.ParseFiles("view/other.html", "view/base.html")),
	}

	// let's pass some data to our template
	data := map[string]string{"name": "John"}

	// render templates
	for _, template := range templates {
		template.ExecuteTemplate(os.Stdout, "base", data)
	}
}
