package main

import "text/template"
import "os"
import "log"
import "fmt"

type Recipient struct {
	Name, Gift string
	attended   bool
}

// Ten przykład nie zadziała nie może
// być wskaźnik ??? nie wiaodmo dlaczego
func (r *Recipient) Attended() bool {
	return r.attended
}

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
