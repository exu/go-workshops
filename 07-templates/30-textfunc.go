package main

import "text/template"
import "strings"
import "log"
import "os"
import "fmt"

// Policzmy spacje w stringu
func countSpaces(str string) string {
	count := strings.Count(str, " ")
	return fmt.Sprintf("Text \"%s\" has %d spaces inside", str, count)
}

func main() {
	// Mapa funkcji
	funcMap := template.FuncMap{
		// nazwy funkcji "title" "up" "countSpaces" będzie można wykorzystać
		// w templatce
		"title":       strings.Title,
		"up":          strings.ToUpper,
		"countSpaces": countSpaces,
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

	// Tworzenie templatki i dodanie naszej mapy z funkcjami
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Odpalamy templatke jako Writer podajemy Stdout
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}
