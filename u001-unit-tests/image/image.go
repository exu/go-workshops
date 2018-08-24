package image

import "fmt"

type Image struct {
	URL         string
	Width       int
	Height      int
	Type        string
	Description string
	Metadata    map[string]string
}

func (i *Image) IsVertical() bool {
	return i.Width < i.Height
}

func (i *Image) IsHorizontal() bool {
	return i.Width > i.Height
}

func (i *Image) String() string {
	return fmt.Sprintf("URL: %s\n(%dx%d)\n\n%s", i.URL, i.Width, i.Height, i.Description)
}
