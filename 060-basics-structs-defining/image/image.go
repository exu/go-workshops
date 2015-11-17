package image

type image struct {
	URL         string
	contentType string // z małej litery prywatna zmienna w obrębie package
	size        int
}

// metoda z tzw receiverem "img"
func (img *image) GetSize() int {
	return img.size
}

func (img *image) SetSize(size int) {
	img.size = size
}

func NewImage(size int) image {
	return image{"aaa", "asdj", size}
}
