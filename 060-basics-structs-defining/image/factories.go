package image

func NewImage(size int) Image {
	return Image{"aaa", "asdj", size}
}

// NewImageWith800Width returns image with 800px width
func NewImageWith800Width() Image {
	return Image{"", "", 800}
}
