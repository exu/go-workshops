package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)

const DATA_PATH = "data/"

var templates = template.Must(template.ParseFiles(
	"tpl/edit.html", "tpl/view.html", "tpl/list.html"))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title   string
	RawBody []byte
	Body    []byte
}

func (p *Page) save() error {
	filename := DATA_PATH + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.RawBody, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := DATA_PATH + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile("\\[([a-zA-Z0-9]+)\\]")
	escapedBody := re.ReplaceAllFunc(body, func(str []byte) []byte {
		m := re.FindStringSubmatch(string(str))
		return []byte("<a href=\"/view/" + m[1] + "\">" + m[1] + "</a>")
	})

	return &Page{Title: title, RawBody: body, Body: escapedBody}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	files, _ := filepath.Glob(DATA_PATH + "*.txt")

	for i := range files {
		files[i] = strings.Replace(files[i], DATA_PATH, "", 1)
		files[i] = strings.Replace(files[i], ".txt", "", 1)
	}

	renderTemplate(w, "list", files)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")

	p := &Page{Title: title, RawBody: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func makeWikiFileHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/view/", makeWikiFileHandler(viewHandler))
	http.HandleFunc("/edit/", makeWikiFileHandler(editHandler))
	http.HandleFunc("/save/", makeWikiFileHandler(saveHandler))
	panic(http.ListenAndServe(":8080", nil))
}
