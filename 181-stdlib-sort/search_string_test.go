package main

import (
	"log"
	"testing"
)

func inArray(array []string, element string) bool {
	for _, val := range array {
		if val == element {
			return true
		}
	}

	return false
}

func TestSearchString_ShouldFilterStringArray(t *testing.T) {

	collections := []string{
		"collection",
		"locked_courses",
		"collection",
		"workers",
		"collection",
	}

	exclude := []string{
		"workers",
		"locked_courses",
	}

	out := []string{}
	for _, collection := range collections {
		excluded := inArray(exclude, collection)

		log.Println(collection, excluded)

		if excluded {
			continue
		}
		out = append(out, collection)
	}

	if len(out) != 3 {
		t.Errorf("Want 3 but got %d", len(out))
	}
	log.Println(out, len(out))

}
