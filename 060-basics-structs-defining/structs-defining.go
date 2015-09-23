package main

import "fmt"

// pola zaczynające się dużą literą będą dostępne na zewnątrz package
// nazwane małą literą są prywatne dla paczki
type Image struct {
	URL         string
	ContentType string
	Size        int
}

func main() {
	images := []Image{
		Image{
			URL:         "http://static.ioki.pl/coś.png",
			ContentType: "image/png",
			Size:        2048,
		},
		// możemy też przekazać bez podawania nazw pól
		// jednak wtedy należy podać wszystkie wartości
		Image{
			"http://static.ioki.pl/nic.jpg",
			"image/jpg",
			1024,
		},
		// jeżeli nie podamy wartości któregokolwiek z pól,
		// zostanie przypisana wartość zerowa (NIE NULL)
		// dla stringów "", int 0, float 0.0 itd.
		Image{
			URL: "http://static.ioki.pl/tylko-url.jpg",
		},
	}

	for _, image := range images {
		fmt.Println(image.URL, "with content type:", image.ContentType) // dostęp do pól zagnieżdżonej strukturki
	}

}
