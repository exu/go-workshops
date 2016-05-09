package main

import (
	"fmt"
	"sort"
)

type Country struct {
	Name string
}

type Countries []Country

func (slice Countries) Len() int {
	return len(slice)
}

func (slice Countries) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

func (slice Countries) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func main() {
	countries := Countries{
		{Name: "United States"},
		{Name: "Bahamas"},
		{Name: "Japan"},
	}

	sort.Sort(countries)

	for i, c := range countries {
		fmt.Println(i, c.Name)
	}
}
