// Go 1.6 needed
package main

import (
	"html/template"
	"os"
)

func main() {
	tmpl, err := template.New("test").Parse("{{ 23 -}} h jkh kjkh \n\n\n\n hjkhj hjkh jkh jkh  {{- 45 }}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
