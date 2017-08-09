package image

type Image struct {
	URL         string
	contentType string // z małej litery prywatna zmienna w obrębie package
	size        int
}

// metoda z tzw receiverem "img"
func (img *Image) GetSize() int {
	return img.size
}

func (img *Image) SetSize(size int) {
	img.size = size
}
