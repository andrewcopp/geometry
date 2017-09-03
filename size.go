package geometry

// Size represents the space of a rectangle in a two-dimensional space.
type Size struct {
	Width  int
	Height int
}

// NewSize creates a Size struct given a width and height value.
func NewSize(w, h int) *Size {
	return &Size{
		Width:  w,
		Height: h,
	}
}
