package main

import "fmt"

// fields starting with capital letter will be accesible outside package (like public in JAVA or PHP)
// fields starting with small letter will be package private
type Image struct {
	URL         string
	ContentType string
	Size        int
}

func main() {
	images := []Image{
		Image{
			ContentType: "image/png",
			URL:         "http://static.ioki.pl/coś.png",
			Size:        2048,
		},
		// if we don't set value for field,
		// there will be assigned zero-ed value automatically (NOT NULL)
		// for strings "", int 0, float 0.0 etc.
		Image{},
		Image{
			URL: "http://static.ioki.pl/tylko-url.jpg",
		},
		// There is no need for naming struct fields during initialisation
		// But then you'll need to pass all struct fields.
		Image{
			"http://static.ioki.pl/nic.jpg",
			"image/jpg",
			1024,
		},
	}

	// let's iterate over our slice of structs
	for _, image := range images {
		fmt.Println(image.URL, "with content type:", image.ContentType) // we're accesing our Image struct fields.
	}

}
