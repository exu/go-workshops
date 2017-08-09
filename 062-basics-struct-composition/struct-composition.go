package main

import "fmt"

// pola zaczynające się dużą literą będą dostępne na zewnątrz package
// nazwane małą literą są prywatne dla paczki

type Dimension struct {
	Width  int
	Height int
}

func (this Dimension) IsVertical() bool {
	if this.Width < this.Height {
		return true
	} else {
		return false
	}
}

// wymiar jest zagnieżdżony (embedded)
// możemy dowolnie składać strukturki
type Image struct {
	Dimension   // embedded struct
	URL         string
	ContentType string
	Size        int
}

func main() {
	images := []Image{
		Image{
			Dimension:   Dimension{100, 200},
			URL:         "http://static.ioki.pl/coś.png",
			ContentType: "image/png",
		},
		Image{
			Dimension{200, 100},
			"http://static.ioki.pl/nic.jpg",
			"image/jpg",
			1024,
		},
	}

	for _, image := range images {
		fmt.Println(image.URL, "is vertical? : ", image.Height) // dostęp do pól zagnieżdżonej strukturki
	}

}
