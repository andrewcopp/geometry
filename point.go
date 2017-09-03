package geometry

import "math"

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

// To returns the distance to another point.
func (p *Point) To(other *Point) float64 {
	X2 := math.Pow(float64(p.X-other.X), 2)
	Y2 := math.Pow(float64(p.Y-other.Y), 2)
	return math.Sqrt(X2 + Y2)
}
