package main

import "text/template"
import "os"
import "log"
import "fmt"

// templates in go working great with struct types
// you'll have great static typed templates
// errors checked during compilation

func main() {

	const letter = `
Dear {{.Name}},
{{if .Attended}}
    It was a pleasure to see you at the wedding.
{{else}}
    It is a shame you couldn't make it to the wedding.
{{end}}
{{with .Gift}}
    Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

	// template data
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin $100 pants", false},
		{"Cousin Rodney", "", false},
	}

	// Tworzymy nowy obiekt templatki i wrzucamy do niego sparsowany list
	t := template.Must(template.New("letter").Parse(letter))

	// Wykonujemy dla każdego odbiorcy
	for _, r := range recipients {
		fmt.Println("----------------------------------------")
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
