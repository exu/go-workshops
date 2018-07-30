package main

import "fmt"

type Dimension struct {
	Width  int
	Height int
}

func (this Dimension) IsVertical() bool {
	if this.Width < this.Height {
		return true
	}

	return false

}

// Dimension is embedded
// we can composite struct with other ones
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

		fmt.Println(image.URL, "is vertical? : ", image.IsVertical())
		// we have direct access to fields and methods of embedded struct
	}

}
