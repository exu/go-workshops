package main

import "html/template"
import "fmt"
import "os"

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

	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")

	if err != nil {
		fmt.Errorf("Error %s", err)
	}
}
