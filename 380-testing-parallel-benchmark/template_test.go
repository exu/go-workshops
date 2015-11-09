package main

import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	// RunParallel will create GOMAXPROCS goroutines
	// and distribute work among them.

	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
