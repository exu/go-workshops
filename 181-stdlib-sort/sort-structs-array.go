package main

import (
	"fmt"
	"sort"
)

// let's define country struct
type Country struct {
	Name string
}

type Countries []Country

// to enable sorting in struct you'll need to implement
// sort.Interface:
func (slice Countries) Len() int {
	return len(slice)
}

func (slice Countries) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

func (slice Countries) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// now we can create some slice of countries
func main() {
	countries := Countries{
		{Name: "United States"},
		{Name: "Bahamas"},
		{Name: "Japan"},
	}

	// and sort them with as implemented in our methods
	sort.Sort(countries)

	for i, c := range countries {
		fmt.Println(i, c.Name)
	}
}
