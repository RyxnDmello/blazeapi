package widgets

type Dimension struct {
	width  int
	height int
}

func (dimension Dimension) Width() (width int) {
	return width
}

func (dimension Dimension) SetWidth(width int) Dimension {
	dimension.width = width
	return dimension
}

func (dimension Dimension) Height() (height int) {
	return height
}

func (dimension Dimension) SetHeight(height int) Dimension {
	dimension.height = height
	return dimension
}
