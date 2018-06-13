package main

import "html/template"
import "fmt"
import "os"

// the main power of HTML template engine is context awareness
// from now we can safely pass data to our templates.
func main() {

	//simple from string
	t, err := template.New("foo").Parse(`
{{define "T"}}
{{.}}
<a title='{{.}}'>
<a href="/{{.}}">
<a href="?q={{.}}">
<a on='f("{{.}}")'>
<a on='f({{.}})'>
<a on='pattern = /{{.}}/;'>
{{end}}
`)
	// now try to bypass our template with wimple injection:
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")

	// as we can see in each context data are escaped in different way.
	if err != nil {
		fmt.Errorf("Error %s", err)
	}
}
