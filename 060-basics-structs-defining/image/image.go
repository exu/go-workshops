package image

type Image struct {
	URL         string
	contentType string // beginning with small letter - private for package`(like protected in JAVA or PHP)
	size        int
}

// GetSize - struct method with receiver `(img *Image)`
func (img *Image) GetSize() int {
	return img.size
}

func (img *Image) SetSize(size int) {
	img.size = size
}
