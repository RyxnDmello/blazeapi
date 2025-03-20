package widgets

type Dimension struct {
	width  int
	height int
}

// Width returns the width of the Dimension type.
//
// The width value is retrieved from the Dimension instance. This function
// provides access to the current width setting.
//
// # Returns
//
//	width int
//
// # Usage
//
//	width := dimension.Width()
func (dimension Dimension) Width() (width int) {
	return dimension.width
}

// SetWidth sets the width of the Dimension type.
//
// The width value is updated to the specified integer. Method chaining is
// supported by returning the modified Dimension.
//
// # Parameters
//
//	width int
//
// # Returns
//
//	Dimension
//
// # Usage
//
//	dimension = dimension.SetWidth(0)
func (dimension Dimension) SetWidth(width int) Dimension {
	dimension.width = width
	return dimension
}

// Height returns the height of the Dimension type.
//
// The height value is retrieved from the Dimension instance. This function
// provides access to the current height setting.
//
// # Returns
//
//	height int
//
// # Usage
//
//	height := dimension.Height()
func (dimension Dimension) Height() (height int) {
	return dimension.height
}

// SetHeight sets the height of the Dimension type.
//
// The height value is updated to the specified integer. Method chaining is
// supported by returning the modified Dimension.
//
// # Parameters
//
//	height int
//
// # Returns
//
//	Dimension
//
// # Usage
//
//	dimension = dimension.SetHeight(0)
func (dimension Dimension) SetHeight(height int) Dimension {
	dimension.height = height
	return dimension
}
