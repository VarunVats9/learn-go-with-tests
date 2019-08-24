package interfaces

import "math"

// Shape is a interface
type Shape interface {
	Area() float64
}

// Rectangle struct to store the width and height fields.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area on the rectangle receiver.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle struct to store the radius field.
type Circle struct {
	Radius float64
}

// Area calculates the area on the circle receiver.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle struct to store the base and height fields.
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area on the triangle receiver.
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter calculates the perimeter for the given width and height.
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area calculates the area for the given width and height.
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
