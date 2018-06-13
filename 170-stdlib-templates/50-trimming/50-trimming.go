// Go 1.6 needed
package main

import (
	"html/template"
	"os"
)

// look into docs for more details
// https://golang.org/pkg/text/template/#hdr-Text_and_spaces
func main() {
	tmpl, err := template.New("test").Parse("{{ 23 -}}                h jkh kjkh \n\n\n\n hjkhj hjkh jkh jkh          {{- 45 }}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
