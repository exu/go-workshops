package main

import (
	"html/template"
	"os"
)

type templateMap map[string]*template.Template

func main() {

	// definiujemy templatki które mają dziedziczyć z base.html
	// parse files weżmie pod uwagę przekazane parametry
	templates := templateMap{
		"index.html": template.Must(template.ParseFiles("view/index.html", "view/base.html")),
		"other.html": template.Must(template.ParseFiles("view/other.html", "view/base.html")),
	}

	// przekazujemy trochę danych do template'a
	// cała mapa będzie dostępna jako "."
	data := map[string]string{"name": "John"}

	// renderujemy templatki
	for _, template := range templates {
		template.ExecuteTemplate(os.Stdout, "base", data)
	}
}
