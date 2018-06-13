package main

import "text/template"
import "strings"
import "log"
import "os"
import "fmt"

// let's count some spaces in string
func countSpaces(str string) string {
	count := strings.Count(str, " ")
	return fmt.Sprintf("Text \"%s\" has %d spaces inside", str, count)
}

func main() {
	// To use functions we need to declare function map
	funcMap := template.FuncMap{
		// names "title" "up" "countSpaces" can be used in template
		"title":       strings.Title,   // go function in string package
		"up":          strings.ToUpper, // go function in string package
		"countSpaces": countSpaces,     // our custom function
	}

	const templateText = `
Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
Output 3: {{printf "%q" . | countSpaces}}
Output 4: {{countSpaces .}}
Output 5 {{ . | up}} is obvious
`

	// Create new template and passing function map
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// execute template and set writer to stdout
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}
