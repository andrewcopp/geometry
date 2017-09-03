package geometry

// Point represents a one-dimensional space in a Cartesian plane.
type Point struct {
	X int
	Y int
}

// NewPoint creates a Point struct given an x and y coordinate.
func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}
