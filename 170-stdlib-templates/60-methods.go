package main

import "text/template"
import "os"
import "log"
import "fmt"

type Recipient struct {
	Name, Gift string
	attended   bool
}

// It'll won't work for pointer receiver
// need to be value receiver.
// you'll find full available arguments list in documentation:
// https://golang.org/pkg/text/template/#hdr-Arguments
func (r Recipient) Attended() bool {
	return r.attended
}

func main() {

	const letter = `
Dear {{.Name}},
{{ if .Attended }}
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

	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin $100 pants", false},
		{"Cousin Rodney", "", false},
	}

	t := template.Must(template.New("letter").Parse(letter))

	for _, r := range recipients {
		fmt.Println("----------------------------------------")
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
